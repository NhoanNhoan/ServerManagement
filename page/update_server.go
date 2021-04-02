package page

import (
	"CURD/database"
	"CURD/entity"
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"CURD/repo/server/hardware_repo"
	"database/sql"
	"errors"
)

type HardwareTemplate struct {
	CPUs        []hardware.CPU
	Clusters	[]hardware.ClusterServer
	Chassis		[]hardware.Chassis
	RAMs        []hardware.RAM
	Disks       []hardware.Disk
	NICs        []hardware.NIC
	Raids       []hardware.Raid
	PSUs		[]hardware.PSU
	Managements []hardware.Management
}

func (template *HardwareTemplate) New() (err error) {
	cpuRepo := hardware_repo.CPURepo{}
	if template.CPUs, err = cpuRepo.FetchAllCPUs(); nil != err {
		return err
	}

	clusterRepo := hardware_repo.ClusterRepo{}
	if template.Clusters, err = clusterRepo.FetchAllClusterServers(); nil != err {
		return err
	}

	ramRepo := hardware_repo.RAMRepo{}
	if template.RAMs, err = ramRepo.FetchAllRAMs(); nil != err {
		return err
	}

	diskRepo := hardware_repo.DiskRepo{}
	if template.Disks, err = diskRepo.FetchAllDisks(); nil != err {
		return err
	}

	nicRepo := hardware_repo.NICRepo{}
	if template.NICs, err = nicRepo.FetchAllNICs(); nil != err {
		return err
	}

	raidRepo := hardware_repo.RaidRepo{}
	if template.Raids, err = raidRepo.FetchAllRaids(); nil != err {
		return err
	}

	psuRepo := hardware_repo.PSURepo{}
	if template.PSUs, err = psuRepo.FetchAllPSUs(); nil != err {
		return err
	}

	mntRepo := hardware_repo.ManagementRepo{}
	if template.Managements, err = mntRepo.FetchAllManagements(); nil != err {
		return err
	}

	return
}

type SwitchInfo struct {
	entity.Switch
	entity.SwitchConnection
	entity.CableType
}

type UpdateServer struct {
	entity.Server
	DCs               []entity.DataCenter
	Racks             []entity.Rack
	RackUnits         []entity.RackUnit
	PortTypes         []entity.PortType
	ServerStates      []entity.ServerStatus
	SwitchArr         []entity.Switch
	Cables            []entity.CableType
	Tagged            []entity.Tag
	Untagged          []entity.Tag
	ConnectedSwitches []SwitchInfo
	HardwareTemplate
}

func (obj *UpdateServer) New(ServerId string) (err error) {
	if err = obj.makeServer(ServerId); nil != err {
		return
	}
	if err = obj.initDCs(); nil != err {
		return
	}
	if err = obj.initRacks(); nil != err {
		return
	}
	if err = obj.initRackUnits(); nil != err {
		return
	}
	if err = obj.initPortTypes(); nil != err {
		return
	}
	if err = obj.makeServerStates(); nil != err {
		return
	}
	if err = obj.initSwitchArr(); nil != err {
		return
	}
	if err = obj.initTaggedArr(); nil != err {
		return
	}
	if err = obj.initUntaggedArr(); nil != err {
		return
	}
	if err = obj.initCableTypes(); nil != err {
		return
	}
	if err = obj.initConnectedSwitches(obj.Server.Id); nil != err {
		return err
	}
	if err = obj.initSwitchArr(); nil != err {
		return err
	}

	return obj.HardwareTemplate.New()
}

func (obj *UpdateServer) NewByIpServer(ip string) (err error) {
	var ipAddress entity.IpAddress
	if !ipAddress.Parse(ip) {
		return errors.New(ip + " is not like format ip address")
	}

	serverId, err := server.ServerIpRepo{}.FetchServerIdByIp(ipAddress)
	if nil != err {
		return err
	}

	if "" == serverId {
		return errors.New("Not found the server that has ip: " + ip)
	}

	return obj.New(serverId)
}

func (obj *UpdateServer) makeServer(Id string) error {
	obj.Server.Id = Id
	comp := obj.makeFetchServerComp()
	servers, err := server.ServerRepo{}.Fetch(comp, obj.scanServer)

	if len(servers) > 0 {
		obj.Server = servers[0]
		obj.Server.Id = Id
		err = obj.setServerIpAddresses()
	}

	if nil == err {
		var redfishIpAddrs []entity.IpAddress
		redfishIpAddrs, err = server.ServerIpRepo{}.FetchRedfishIp(obj.Server.Id)
		if len(redfishIpAddrs) > 0 {
			obj.Server.RedfishIp = redfishIpAddrs[0].String()
		}
	}

	return err
}

func (obj UpdateServer) scanServer(content interface{}, row *sql.Rows) (interface{}, error) {
	server := content.(entity.Server)
	err := row.Scan(
		&server.DC.Id,
		&server.DC.Description,
		&server.Rack.Id,
		&server.Rack.Description,
		&server.UStart.Id,
		&server.UStart.Description,
		&server.UEnd.Id,
		&server.UEnd.Description,
		&server.SSD,
		&server.HDD,
		&server.Maker,
		&server.PortType.Id,
		&server.PortType.Description,
		&server.SerialNumber,
		&server.ServerStatus.Id,
		&server.ServerStatus.Description,
	)

	return server, err
}

func (obj *UpdateServer) setServerIpAddresses() (err error) {
	r := server.ServerIpRepo{}
	obj.Server.IpAddrs, err = r.FetchServerIpAddrs(obj.Server.Id)
	return err
}

func (obj *UpdateServer) makeFetchServerComp() database.QueryComponent {
	return database.QueryComponent{
		Tables: []string{
			"SERVER AS S",
			"DC AS D",
			"RACK AS R",
			"RACK_UNIT AS USTART",
			"RACK_UNIT AS UEND",
			"PORT_TYPE AS PT",
			"SERVER_STATUS AS SS",
			"STATUS_ROW AS SR",
		},

		Columns: []string{
			"D.ID",
			"D.DESCRIPTION",
			"R.ID",
			"R.DESCRIPTION",
			"USTART.ID",
			"USTART.DESCRIPTION",
			"UEND.ID",
			"UEND.DESCRIPTION",
			"S.SSD",
			"S.HDD",
			"S.MAKER",
			"PT.ID",
			"PT.DESCRIPTION",
			"S.SERIAL_NUMBER",
			"SS.ID",
			"SS.DESCRIPTION",
		},

		Selection: "S.ID = ? AND " +
			"S.ID_DC = D.ID AND " +
			"S.ID_RACK = R.ID AND " +
			"S.ID_U_START = USTART.ID AND " +
			"S.ID_U_END = UEND.ID AND " +
			"SR.DESCRIPTION = ? AND S.ID_STATUS_ROW = SR.ID AND " +
			"S.ID_PORT_TYPE = PT.ID AND " +
			"S.ID_SERVER_STATUS = SS.ID",

		SelectionArgs: []string{obj.Server.Id, "available"},
	}
}

func (obj *UpdateServer) initDCs() error {
	comp := database.QueryComponent{
		Tables:  []string{"DC"},
		Columns: []string{"ID", "DESCRIPTION"},
	}

	scanDC := func(obj interface{}, rows *sql.Rows) (interface{}, error) {
		dc := obj.(entity.DataCenter)
		err := rows.Scan(&dc.Id, &dc.Description)
		return dc, err
	}

	var err error
	obj.DCs, err = server.DCRepo{}.Fetch(comp, scanDC)

	return err
}

func (obj *UpdateServer) initRacks() error {
	comp := database.QueryComponent{
		Tables:  []string{"RACK"},
		Columns: []string{"ID", "DESCRIPTION"},
	}

	scanRack := func(obj interface{}, rows *sql.Rows) (interface{}, error) {
		rack := obj.(entity.Rack)
		err := rows.Scan(&rack.Id, &rack.Description)
		return rack, err
	}

	var err error
	obj.Racks, err = server.RackRepo{}.Fetch(comp, scanRack)

	return err
}

func (obj *UpdateServer) initRackUnits() error {
	comp := database.QueryComponent{
		Tables:  []string{"RACK_UNIT"},
		Columns: []string{"ID", "DESCRIPTION"},
	}

	scanRackUnit := func(obj interface{}, rows *sql.Rows) (interface{}, error) {
		unit := obj.(entity.RackUnit)
		err := rows.Scan(&unit.Id, &unit.Description)
		return unit, err
	}

	var err error
	obj.RackUnits, err = server.RackUnitRepo{}.Fetch(comp, scanRackUnit)

	return err
}

func (obj *UpdateServer) initPortTypes() error {
	comp := database.QueryComponent{
		Tables:  []string{"PORT_TYPE"},
		Columns: []string{"ID", "DESCRIPTION"},
	}

	scanPortType := func(obj interface{}, rows *sql.Rows) (interface{}, error) {
		portType := obj.(entity.PortType)
		err := rows.Scan(&portType.Id, &portType.Description)
		return portType, err
	}

	var err error
	obj.PortTypes, err = server.PortTypeRepo{}.Fetch(comp, scanPortType)

	return err
}

func (obj *UpdateServer) makeServerStates() error {
	comp := database.QueryComponent{
		Tables:  []string{"SERVER_STATUS"},
		Columns: []string{"ID", "DESCRIPTION"},
	}

	scanStates := func(obj interface{}, rows *sql.Rows) (interface{}, error) {
		state := obj.(entity.ServerStatus)
		err := rows.Scan(&state.Id, &state.Description)
		return state, err
	}

	var err error
	obj.ServerStates, err = server.ServerStatusRepo{}.Fetch(comp, scanStates)

	return err
}

func (obj *UpdateServer) initTaggedArr() (err error) {
	tagRepo := server.TagRepo{}
	obj.Tagged, err = tagRepo.FetchTaggedServer(obj.Server.Id)
	return err
}

func (obj *UpdateServer) initUntaggedArr() (err error) {
	tagRepo := server.TagRepo{}
	obj.Untagged, err = tagRepo.FetchUntaggedServer(obj.Server.Id)
	return err
}

func (obj *UpdateServer) initCableTypes() (err error) {
	comp := database.QueryComponent{
		Tables:  []string{"CABLE_TYPE"},
		Columns: []string{"ID", "NAME", "SIGN_PORT"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		cab := obj.(entity.CableType)
		err := row.Scan(&cab.Id, &cab.Name, &cab.SignPort)
		return cab, err
	}

	obj.Cables, err = server.CableTypeRepo{}.Fetch(comp, scan)
	return err
}

func (obj *UpdateServer) initConnectedSwitches(ServerId string) error {
	connArr, err := server.SwitchConnectionRepo{}.FetchByServerId(ServerId)
	if nil != err {
		return err
	}

	obj.ConnectedSwitches = make([]SwitchInfo, len(connArr))

	for i := range connArr {
		obj.ConnectedSwitches[i].Switch = server.SwitchRepo{}.FetchById(connArr[i].SwitchId)
		obj.ConnectedSwitches[i].SwitchConnection = connArr[i]
		obj.ConnectedSwitches[i].CableType = server.CableTypeRepo{}.FetchById(connArr[i].CableTypeId)
	}

	//comp := obj.queryConnectedSwitchesComp(ServerId)
	//rows, err := database.Query(comp)
	//defer rows.Close()
	//
	//var switchInfo SwitchInfo
	//for rows.Next() && nil == err {
	//	err = switchInfo.ParseFromRow(rows)
	//	switchInfo.Switch.FetchIpAddrs()
	//	obj.ConnectedSwitches = append(obj.ConnectedSwitches,
	//								switchInfo)
	//}
	//
	//return err
	return nil
}

func (obj *UpdateServer) queryConnectedSwitchesComp(ServerId string) database.QueryComponent {
	return database.QueryComponent{
		Tables: []string{
			"SWITCH AS SW",
			"SWITCH_CONNECTION AS SC",
			"DC AS D",
			"RACK AS R",
			"RACK_UNIT AS USTART",
			"RACK_UNIT AS UEND",
			"CABLE_TYPE AS CT",
		},

		Columns: []string{
			"SW.ID",
			"D.DESCRIPTION",
			"R.DESCRIPTION",
			"USTART.DESCRIPTION",
			"UEND.DESCRIPTION",
			"SW.MAXIMUM_PORT",
			"SC.PORT",
			"CT.NAME",
			"CT.SIGN_PORT",
		},

		Selection: "SC.ID_SERVER = ? AND " +
			"SC.ID_SWITCH = SW.ID AND " +
			"SW.ID_DC = D.ID AND " +
			"SW.ID_RACK = R.ID AND " +
			"SW.ID_U_START = USTART.ID AND " +
			"SW.ID_U_END = UEND.ID AND " +
			"SC.ID_CABLE_TYPE = CT.ID",

		SelectionArgs: []string{ServerId},
	}
}

func (obj *UpdateServer) initSwitchArr() (err error) {
	comp := obj.fetchSwitchComp()
	scanSwitch := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		sw := obj.(entity.Switch)
		err := row.Scan(&sw.Id,
			&sw.DC.Description,
			&sw.Rack.Description,
			&sw.UStart.Description,
			&sw.UEnd.Description)
		return sw, err
	}

	obj.SwitchArr, err = server.SwitchRepo{}.Fetch(comp, scanSwitch)
	return err
}

func (obj *UpdateServer) fetchSwitchComp() database.QueryComponent {
	return database.QueryComponent{
		Tables: []string{"SWITCH AS S",
			"DC AS D",
			"RACK AS R",
			"RACK_UNIT AS USTART",
			"RACK_UNIT AS UEND",
		},
		Columns: []string{
			"S.ID",
			"D.DESCRIPTION",
			"R.DESCRIPTION",
			"USTART.DESCRIPTION",
			"UEND.DESCRIPTION",
		},
		Selection: "S.ID_DC = D.ID AND " +
			"S.ID_RACK = R.ID AND " +
			"S.ID_U_START = USTART.ID AND " +
			" S.ID_U_END = UEND.ID",
	}
}

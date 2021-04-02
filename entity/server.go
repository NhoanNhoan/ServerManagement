package entity

import (
	"CURD/database"
)

type qcomp = database.QueryComponent
type ucomp = database.UpdateComponent
type icomp = database.InsertComponent

type Server struct {
	Id string
	DC DataCenter
	Rack
	UStart RackUnit
	UEnd RackUnit
	SSD	string
	HDD	string
	Maker string
	PortType
	SerialNumber string
	ServerStatus
	ServeCustomer
	RedfishIp string
	IpAddrs []IpAddress
	Services []string
	Events []ServerEvent
}

//func (s *Server) MakeByContext(c *gin.Context) {
//
//}
//
//func (obj *Server) New(Id string) (err error) {
//	obj.SetId(Id)
//	component := obj.makeQueryComponent(Id)
//	rows, err := database.Query(component)
//	defer rows.Close()
//
//	if nil != err {
//		return err
//	}
//
//	if rows.Next() {
//		err = rows.Scan(
//					&obj.DC.Id,
//					&obj.DC.Description,
//					&obj.Rack.Id,
//					&obj.Rack.Description,
//					&obj.UStart.Id,
//					&obj.UStart.Description,
//					&obj.UEnd.Id,
//					&obj.UEnd.Description,
//					&obj.SSD,
//					&obj.HDD,
//					&obj.Maker,
//					&obj.PortType.Id,
//					&obj.PortType.Description,
//					&obj.SerialNumber,
//					&obj.ServerStatus.Id,
//					&obj.ServerStatus.Description,
//			)
//	}
//
//	return err
//}
//
//func (obj *Server) SetId(Id string) {
//	obj.Id = Id
//}
//
//func (obj *Server) makeQueryComponent(IdServer string) qcomp {
//	return qcomp {
//		Tables: []string {
//				"SERVER AS S",
//				"DC AS D",
//				"RACK AS R",
//				"RACK_UNIT AS USTART",
//				"RACK_UNIT AS UEND",
//				"PORT_TYPE AS PT",
//				"SERVER_STATUS AS SS",
//				"STATUS_ROW AS SR",
//				"REDFISH_IP",
//			},
//
//		Columns: []string {
//				"D.ID",
//				"D.DESCRIPTION",
//				"R.ID",
//				"R.DESCRIPTION",
//				"USTART.ID",
//				"USTART.DESCRIPTION",
//				"UEND.ID",
//				"UEND.DESCRIPTION",
//				"S.SSD",
//				"S.HDD",
//				"S.MAKER",
//				"PT.ID",
//				"PT.DESCRIPTION",
//				"S.SERIAL_NUMBER",
//				"SS.ID",
//				"SS.DESCRIPTION",
//			},
//
//		Selection: "S.ID = ? AND " +
//					"S.ID_DC = D.ID AND " +
//					"S.ID_RACK = R.ID AND " +
//					"S.ID_U_START = USTART.ID AND " +
//					"S.ID_U_END = UEND.ID AND " +
//					"SR.DESCRIPTION = ? AND S.ID_STATUS_ROW = SR.ID AND " +
//					"S.ID_PORT_TYPE = PT.ID AND " +
//					"S.ID_SERVER_STATUS = SS.ID",
//
//		SelectionArgs: []string {IdServer, "available"},
//		GroupBy: "",
//		Having: "",
//		OrderBy: "",
//		Limit: "",
//	}
//}
//
//func (obj *Server) ParseRow(rows *sql.Rows) error {
//	return rows.Scan(
//					&obj.DC.Id,
//					&obj.DC.Description,
//					&obj.Rack.Id,
//					&obj.Rack.Description,
//					&obj.UStart.Id,
//					&obj.UStart.Description,
//					&obj.UEnd.Id,
//					&obj.UEnd.Description,
//					&obj.SSD,
//					&obj.HDD,
//					&obj.Maker,
//					&obj.PortType.Id,
//					&obj.PortType.Description,
//					&obj.SerialNumber,
//					&obj.ServerStatus.Id,
//					&obj.ServerStatus.Description,
//				)
//}
//
//func (obj *Server) FetchIpAddrs() {
//	component := obj.makeIpQueryComponent()
//	if rows, err := database.Query(component); err == nil {
//		defer rows.Close()
//		obj.initIpAddrs(rows)
//	}
//}
//
//func (obj Server) makeIpQueryComponent() qcomp {
//	return qcomp {
//		Tables: []string {"IP_SERVER", "IP_NET"},
//		Columns: []string {"IP_NET.ID", "IP_NET.VALUE", "IP_SERVER.IP_HOST"},
//		Selection: "IP_SERVER.ID_SERVER = ? AND IP_SERVER.ID_IP_NET = IP_NET.ID",
//		SelectionArgs: []string {obj.Id},
//	}
//}
//
//func (obj *Server) initIpAddrs(rows *sql.Rows) {
//	var ipNetId, ipNetHost, ipHost string
//	var ipAddr IpAddress
//
//	for rows.Next() && (nil == rows.Scan(&ipNetId, &ipNetHost, &ipHost)) {
//		ipAddr.New(ipNetHost, ipHost)
//		ipAddr.IpNet.Id = ipNetId
//		obj.IpAddrs = append(obj.IpAddrs, ipAddr)
//	}
//}
//
//func (obj *Server) FetchServices() {
//	comp := obj.makeServicesQueryComponent()
//	if rows, err := database.Query(comp); nil == err {
//		defer rows.Close()
//		obj.initServerServices(rows)
//	}
//}
//
//func (obj *Server) initServerServices(rows *sql.Rows) {
//	var service string
//	for rows.Next() && (nil == rows.Scan(&service)) {
//		obj.Services = append(obj.Services, service)
//	}
//}
//
//func (obj Server) makeServicesQueryComponent() database.QueryComponent {
//	return database.QueryComponent {
//		Tables: []string {"SERVICES"},
//		Columns: []string {"SERVICE"},
//		Selection: "id_SERVER = ?",
//		SelectionArgs: []string {obj.Id},
//	}
//}
//
//func (s *Server) Insert() error {
//	comp := s.makeInsertServerComponent()
//	return database.Insert(comp)
//}
//
//func (s *Server) makeInsertServerComponent() icomp {
//	return icomp {
//		Table: "SERVER",
//		Columns: []string {
//			"ID",
//			"ID_DC",
//			"ID_RACK",
//			"ID_U_START",
//			"ID_U_END",
//			"ID_PORT_TYPE",
//			"ID_SERVER_STATUS",
//			"SSD",
//			"HDD",
//			"MAKER",
//			"SERIAL_NUMBER",
//			"id_STATUS_ROW",
//		},
//		Values: [][]string {
//				[]string {s.Id,
//				s.DC.Id,
//				s.Rack.Id,
//				s.UStart.Id,
//				s.UEnd.Id,
//				s.PortType.Id,
//				s.ServerStatus.Id,
//				s.SSD,
//				s.HDD,
//				s.Maker,
//				s.SerialNumber,
//				"1",
//				},
//		},
//	}
//}
//
//func (s *Server) InsertIpAddresses() (err error) {
//	for i := 0; i < len(s.IpAddrs) && nil == err; i++{
//		err = s.IpAddrs[i].Insert(s.Id)
//		s.UpdateIpHostState(i, "used")
//	}
//	return err
//}
//
//func (s *Server) UpdateIpHostState(index int, state string) error {
//	ip := s.IpAddrs[index]
//
//	octets := strings.Split(ip.GetValue(), ".")
//	changedOctetIdx := ip.GetNetmask() / 8
//	octets = octets[0:changedOctetIdx]
//	host := strings.Join(octets, ".") + "." + ip.IpHost
//
//	ipHost := IpHost{
//		IpNet: ip.IpNet,
//		Host: host,
//		State: state,
//	}
//
//	return ipHost.Update()
//}
//
//func (s *Server) Delete() error {
//	comp := s.makeDeleteComp()
//	return database.Delete(comp)
//}
//
//func (s *Server) makeDeleteComp() database.DeleteComponent {
//	return database.DeleteComponent{
//		Table: "SERVER",
//		Selection: "ID = ?",
//		SelectionArgs: []string {s.Id},
//	}
//}
//
//func UpdateServer(server Server) (msg string) {
//	comp := makeUpdateServerComponent(server)
//	if err := database.Update(comp); nil != err {
//		msg = "Can't update the server information"
//		panic (err)
//	} else {
//		msg = "Success"
//	}
//
//	return
//}
//
//func makeUpdateServerComponent(server Server) ucomp {
//	return ucomp {
//		Table: "SERVER",
//		SetClause: "id_DC = ?, " +
//			"id_RACK = ?, " +
//			"id_U_start = ?, " +
//			"id_U_end  = ?, " +
//			"SSD = ?, " +
//			"HDD = ?, " +
//			"id_PORT_TYPE = ?, " +
//			"serial_number = ?, " +
//			"id_SERVER_STATUS = ?",
//		Values: []string {
//				server.DC.Id,
//				server.Rack.Id,
//				server.UStart.Id,
//				server.UEnd.Id,
//				server.SSD,
//				server.HDD,
//				server.PortType.Id,
//				server.SerialNumber,
//				server.ServerStatus.Id,
//			},
//		Selection: "id = ?",
//		SelectionArgs: []string {
//				server.Id,
//			},
//	}
//}
//
//// Fetch Event Array Area
//
//
//func (server *Server) FetchEvents() (err error) {
//	comp := server.makeFetchEventsComponent()
//	if rows, err := database.Query(comp); nil == err {
//		defer rows.Close()
//		return server.makeEventsByRows(rows)
//	}
//	return err
//}
//
//func (server *Server) makeFetchEventsComponent() qcomp {
//	return qcomp {
//		Tables: []string {"SERVER_EVENT"},
//		Columns: []string {"ID", "ID_SERVER", "DESCRIPTION", "OCCUR_AT"},
//		Selection: "ID_SERVER = ?",
//		SelectionArgs: []string {server.Id},
//	}
//}
//
//func (server *Server) makeEventsByRows(rows *sql.Rows) (err error) {
//	event := &ServerEvent{}
//
//	for rows.Next() && nil == err {
//		if err, event = MakeEventByRow(rows); nil != event {
//			server.Events = append(server.Events, *event)
//		}
//	}
//
//	return err
//}
//
//func scanEvent(row *sql.Rows) (event ServerEvent) {
//	err := row.Scan(&event.Id, &event.IdServer, &event.Description, &event.OccurAt)
//	if nil != err {
//		panic (err)
//	}
//	return event
//}
//
//func FetchServers(comp qcomp) (error, []Server) {
//	rows, err := database.Query(comp)
//	defer rows.Close()
//
//	var server Server
//	serverArr := make([]Server, 0)
//	for rows.Next() && nil == err {
//		err = server.ParseRow(rows)
//		serverArr = append(serverArr, server)
//	}
//
//	return err, serverArr
//}
//
//func FetchServer(comp qcomp) (Server, error) {
//	rows, err := database.Query(comp)
//	defer rows.Close()
//
//	var server Server
//	if rows.Next() {
//		err = rows.Scan(
//					&server.Id,
//					&server.DC.Id,
//					&server.DC.Description,
//					&server.Rack.Id,
//					&server.Rack.Description,
//					&server.UStart.Id,
//					&server.UStart.Description,
//					&server.UEnd.Id,
//					&server.UEnd.Description,
//					&server.SSD,
//					&server.HDD,
//					&server.Maker,
//					&server.PortType.Id,
//					&server.PortType.Description,
//					&server.SerialNumber,
//					&server.ServerStatus.Id,
//					&server.ServerStatus.Description,
//				)
//		return server, err
//	}
//
//	return server, err
//}
//
//func (obj *Server) DeleteAllIp() {
//	comp := obj.makeDeleteAllIpComponent()
//	err := database.Delete(comp)
//	if nil != err {
//		panic (err)
//	}
//}
//
//func (obj *Server) makeDeleteAllIpComponent() database.DeleteComponent {
//	return database.DeleteComponent {
//		Table: "IP_SERVER",
//		Selection: "ID_SERVER = ?",
//		SelectionArgs: []string {obj.Id},
//	}
//}
//
//// End of fetch event array area
//
//func FetchServicesByServerId(ServerId string) ([]string) {
//	comp := makeQueryServicesByServerIdComp(ServerId)
//	rows, err := database.Query(comp)
//	defer rows.Close()
//
//	var service string
//	serviceArr := make([]string, 0)
//	for rows.Next() {
//		err = rows.Scan(&service)
//		if nil != err {
//			panic(err)
//		}
//
//		serviceArr = append(serviceArr, service)
//	}
//
//	return serviceArr
//}
//
//func makeQueryServicesByServerIdComp(ServerId string) qcomp {
//	return qcomp {
//		Tables: []string {"SERVICES"},
//		Columns: []string {"SERVICE"},
//		Selection: "ID_SERVER = ?",
//		SelectionArgs: []string {ServerId},
//	}
//}
//
//func InsertServices(ServerId string, Services []string) error {
//	var err error
//	var id string
//	var comp icomp
//
//	for _, service := range Services {
//		id = generateServiceId()
//		comp = makeInsertServiceComponent(service, ServerId, id)
//		err = database.Insert(comp)
//		if nil != err {
//			return err
//		}
//	}
//
//	return nil
//}
//
//func generateServiceId() string {
//	id := database.GeneratePrimaryKey(true,
//									true, true, false, "SE", 6)
//	for IsExistsServiceId(id) {
//		id = database.GeneratePrimaryKey(true,
//									true, true, false, "SE", 6)
//	}
//
//	return id
//}
//
//func IsExistsServiceId(Id string) bool {
//	comp := makeCheckExistsServiceIdComp(Id)
//	rows, err := database.Query(comp)
//	defer rows.Close()
//	return (nil == err) && rows.Next()
//}
//
//func makeCheckExistsServiceIdComp(Id string) qcomp {
//	return qcomp {
//		Tables: []string {"SERVICES"},
//		Columns: []string {"ID"},
//		Selection: "ID = ?",
//		SelectionArgs: []string {Id},
//	}
//}
//
//func makeInsertServiceComponent(Service string, ServerId string, Id string) icomp {
//	return icomp {
//		Table: "SERVICES",
//		Columns: []string {"ID", "ID_SERVER", "SERVICE"},
//		Values: [][]string {[]string {Id, ServerId, Service}},
//	}
//}
//
//func DeleteServicesByServerId(ServerId string) error {
//	comp := makeDeleteServicesByServerIdComp(ServerId)
//	return database.Delete(comp)
//}
//
//func makeDeleteServicesByServerIdComp(ServerId string) database.DeleteComponent {
//	return database.DeleteComponent {
//		Table: "SERVICES",
//		Selection: "ID_SERVER = ?",
//		SelectionArgs: []string {ServerId},
//	}
//}
//
//func (s Server) MakeIpAddressStr(net, host string) string {
//	hosts := strings.Split(host, ".")
//	nets := strings.Split(net, ".")
//
//	for i := range nets {
//		if "" == nets[i] {
//			nets[i] = "0"
//		}
//	}
//
//	for len(nets) < 4 {
//		nets = append(nets, "0")
//	}
//
//	for i := range hosts {
//		nets[len(nets) - i - 1] = hosts[i]
//	}
//
//	return strings.Join(nets, ".")
//}
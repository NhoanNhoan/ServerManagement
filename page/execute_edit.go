package page

import (
	"CURD/database"
	"CURD/entity"
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"CURD/repo/server/hardware_repo"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// parser struct wil parse data that received from gin context
// and convert to entity object
// this class is helper of ExecuteModify struct
type parser struct {
	c *gin.Context
}

func (p parser) parse(content string) string {
	return p.c.PostForm(content)
}

func (p parser) server() entity.Server {
	return entity.Server{
		Id:           p.parse("txtIdServer"),
		SSD:          p.parse("txtSSD"),
		HDD:          p.parse("txtHDD"),
		SerialNumber: p.parse("txtSerialNumber"),
	}
}

func (p parser) dc() entity.DataCenter {
	return entity.DataCenter{
		Id: p.parse("cbDCId"),
	}
}

func (p parser) rack() entity.Rack {
	rack := entity.Rack{Description: p.c.PostForm("txtRack")}
	rackRepo := server.RackRepo{}

	if !rack.IsExistsRackDescription() {
		rack.Id = rackRepo.GenerateId()
		rackRepo.Insert(rack)
	} else {
		rack.Id = rackRepo.FetchId(rack.Description)
	}

	return rack
}

func (p parser) u_start() entity.RackUnit {
	rackUnit := entity.RackUnit{Description: p.c.PostForm("txtUStart")}
	rackUnitRepo := server.RackUnitRepo{}

	if !rackUnit.IsExistsRackUnitDescription() {
		rackUnit.Id = rackUnitRepo.GenerateId()
		rackUnitRepo.Insert(rackUnit)
	} else {
		rackUnit.Id = rackUnitRepo.FetchId(rackUnit.Description)
	}

	return rackUnit
}

func (p parser) uend() entity.RackUnit {
	rackUnit := entity.RackUnit{Description: p.c.PostForm("txtUEnd")}
	rackUnitRepo := server.RackUnitRepo{}

	if !rackUnit.IsExistsRackUnitDescription() {
		rackUnit.Id = rackUnitRepo.GenerateId()
		rackUnitRepo.Insert(rackUnit)
	} else {
		rackUnit.Id = rackUnitRepo.FetchId(rackUnit.Description)
	}

	return rackUnit
}

func (p parser) port_type() entity.PortType {
	return entity.PortType{
		Id: p.parse("cbPortTypeId"),
	}
}

func (p parser) server_state() entity.ServerStatus {
	return entity.ServerStatus{
		Id: p.parse("cbServerStatusId"),
	}
}

func (p parser) serve() entity.ServeCustomer {
	return entity.ServeCustomer{
		Id: p.parse("cbServeCus"),
	}
}

func (p parser) list_ip() ([]entity.IpAddress, error) {
	rawContent := p.parse("txtAllIp")
	if "" == rawContent {
		return nil, nil
	}

	rawIp := strings.Split(rawContent, ",")
	listIp := make([]entity.IpAddress, len(rawIp))

	for i := range listIp {
		values := strings.Split(rawIp[i], ".")
		if len(values) < 4 {
			return nil, errors.New(rawIp[i] + " is not like format")
		}

		numbers := strings.Split(values[3], "/")
		if len(numbers) < 2 {
			continue
		}

		var netmask int
		var err error
		if netmask, err = strconv.Atoi(numbers[1]); nil != err {
			return nil, err
		}

		listIp[i] = entity.IpAddress{
			Octet1: values[0],
			Octet2: values[1],
			Octet3: values[2],
			Octet4: numbers[0],
			Netmask: netmask,
		}
	}

	return listIp, nil
}

func (p parser) redfish_ip() (entity.IpAddress, error) {
	rawContent := p.parse("txtRedfishIp")
	if "" == rawContent {
		return entity.IpAddress{}, nil
	}

	rawIp := strings.Split(rawContent, ", ")
	listIp := make([]entity.IpAddress, len(rawIp))

	for i := range listIp {
		values := strings.Split(rawIp[i], ".")
		if len(values) < 4 {
			return entity.IpAddress{}, errors.New(rawIp[i] + " is not like format")
		}

		numbers := strings.Split(values[3], "/")
		if len(numbers) < 2 {
			continue
		}

		var netmask int
		var err error
		if netmask, err = strconv.Atoi(numbers[1]); nil != err {
			return entity.IpAddress{}, err
		}

		listIp[i] = entity.IpAddress{
			Octet1: values[0],
			Octet2: values[1],
			Octet3: values[2],
			Octet4: numbers[0],
			Netmask: netmask,
		}
	}

	return listIp[0], nil
}

func (p parser) tag_splice() []entity.Tag {
	raw := p.parse("txtAllTag")
	return p.tag_splice_init(strings.Split(raw, ","))
}

func (p parser) tag_splice_init(titles []string) (tags []entity.Tag) {
	tags = make([]entity.Tag, len(titles))
	for i := range titles {
		tags[i].Title = titles[i]
		tags[i].TagId, _ = server.TagRepo{}.IdOf(titles[i])
	}
	return
}

type ExecuteModify struct {
	Msg           string
	entity.Server
	SwitchConnArr []entity.SwitchConnection
	Tags          []entity.Tag
	hardware.HardwareConfig
	HardwareCpuItems  []hardware.HardwareCpu
	HardwareRamItems  []hardware.HardwareRam
	HardwareDiskItems []hardware.HardwareDisk
	HardwareNicItems  []hardware.HardwareNic
	HardwareRaidItems []hardware.HardwareRaid
	HardwarePsuItems  []hardware.HardwarePsu
	HardwareMntItems  []hardware.HardwareManagement
}

type execute = func() error

func (obj *ExecuteModify) New(c *gin.Context) (err error) {
	p := parser{c}

	obj.Server = p.server()
	obj.Server.DC = p.dc()
	obj.Server.Rack = p.rack()
	obj.Server.UStart = p.u_start()
	obj.Server.UEnd = p.uend()
	obj.Server.PortType = p.port_type()
	obj.Server.ServerStatus = p.server_state()
	obj.Server.ServeCustomer = p.serve()
	obj.Server.IpAddrs, err = p.list_ip()
	if nil != err {return}
	obj.Server.RedfishIp, err = p.redfish_ip()
	obj.Tags = p.tag_splice()

	hwParser := HardwareParser{c}
	obj.HardwareConfig = hwParser.HardwareConfig()
	obj.HardwareCpuItems = hwParser.HardwareCPUArray()
	obj.HardwareRamItems = hwParser.HardwareRAMArray()
	obj.HardwareDiskItems = hwParser.HardwareDiskArray()
	obj.HardwareNicItems = hwParser.HardwareNicArray()
	obj.HardwareRaidItems = hwParser.HardwareRaidArray()
	obj.HardwarePsuItems = hwParser.HardwarePsuArray()
	obj.HardwareMntItems = hwParser.HardwareMntArray()

	return
}

func (obj ExecuteModify) Execute() error {
	return obj.executes([]execute{
		obj.updateServer,
		obj.executeListIpServer,
		obj.executeTag,
		obj.executeSwitchConnection,
		obj.executeHardwareComponents,
		obj.executeRedfishIp,
	})
}

func (obj ExecuteModify) executes(methods []execute) (err error) {
	for i := 0; i < len(methods) && nil == err; i++ {
		err = methods[i]()
	}
	return err
}

func (obj ExecuteModify) updateServer() error {
	fmt.Println ("Server Id: ", obj.Server.Id)
	return server.ServerRepo{}.Update(obj.Server)
}

func (obj ExecuteModify) executeListIpServer() (err error) {
	if len(obj.IpAddrs) == 0 {
		return
	}

	r := server.ServerIpRepo{}
	err = r.Delete(obj.Server.Id, obj.Server.IpAddrs...)
	if nil == err {
		err = r.InsertNormalIpAddresses(obj.Server.Id, obj.Server.IpAddrs...)
	}
	return
}

func (obj ExecuteModify) executeRedfishIp() error {
	ipRepo := server.ServerIpRepo{}
	if err := ipRepo.Delete(obj.Server.Id, obj.Server.RedfishIp); nil != err {return err}
	if err := ipRepo.InsertRedfishIpAddresses(obj.Server.Id, obj.Server.RedfishIp); nil != err {
		return err
	}
	return nil
}

func (obj ExecuteModify) executeTag() (err error) {
	fmt.Println ("Tags: ", obj.Tags)
	fmt.Println ("Length: ", len(obj.Tags))
	r := server.ServerTagRepo{}
	err = r.Delete(r.MakeDeleteComp(obj.Server.Id))

	tagRepo := server.TagRepo{}
	for i := range obj.Tags {
		obj.Tags[i].TagId, err = tagRepo.IdOf(obj.Tags[i].Title)
		if nil != err {
			return err
		}
		fmt.Println ("Tags Id: ", obj.Tags[i].TagId)
	}

	if nil == err {
		err = r.Insert(obj.Server.Id, obj.Tags...)
	}
	return
}

func (obj ExecuteModify) executeSwitchConnection() error {
	r := server.SwitchConnectionRepo{}
	comp := database.DeleteComponent{
		Table:         "SWITCH_CONNECTION",
		Selection:     "ID_SERVER = ?",
		SelectionArgs: []string{obj.Server.Id},
	}
	err := r.Delete(comp)
	if nil != err {
		err = r.Insert(obj.SwitchConnArr...)
	}

	return err
}

func (obj *ExecuteModify) executeHardwareComponents() error {
	repo := hardware_repo.HardwareConfigRepo{}
	hwConfig, err := repo.FetchByServerId(obj.Server.Id)
	if nil != err {return err}

	serverRepo := server.ServerRepo{}

	if nil == hwConfig {
		obj.HardwareConfig.Id = repo.GenerateId()
		if err = serverRepo.UpdateHardwareConfig(obj.Server.Id, obj.HardwareConfig.Id); nil != err {
			return err
		}
		repo.Insert(obj.HardwareConfig)
	} else {
		obj.HardwareConfig.Id = hwConfig.Id
		if err = repo.Update(obj.HardwareConfig); nil != err {
			return err
		}
	}

	if err = obj.executeHardwareCPUs(obj.HardwareConfig.Id); nil != err {
		return err
	}

	if err = obj.executeHardwareRams(obj.HardwareConfig.Id); nil != err {
		return err
	}

	if err = obj.executeHardwareDisks(obj.HardwareConfig.Id); nil != err {
		return err
	}

	if err = obj.executeHardwareNics(obj.HardwareConfig.Id); nil != err {
		return err
	}

	if err = obj.executeHardwareRaids(obj.HardwareConfig.Id); nil != err {
		return err
	}

	if err = obj.executeHardwarePsus(obj.HardwareConfig.Id); nil != err {
		return err
	}

	return obj.executeHardwareMnts(obj.HardwareConfig.Id)
}

func (obj ExecuteModify) executeHardwareCPUs(HardwareId string) (err error) {
	cpuRepo := hardware_repo.HardwareCPURepo{}

	if err = cpuRepo.Delete(HardwareId); nil != err {
		return err
	}

	for i := range obj.HardwareCpuItems {
		obj.HardwareCpuItems[i].HardwareId = HardwareId
	}

	return cpuRepo.Insert(HardwareId, obj.HardwareCpuItems...)
}

func (obj ExecuteModify) executeHardwareRams(HardwareId string) (err error) {
	RamRepo := hardware_repo.HardwareRamRepo{}

	if err = RamRepo.Delete(HardwareId); nil != err {
		return err
	}

	for i := range obj.HardwareRamItems {
		obj.HardwareRamItems[i].HardwareId = HardwareId
	}


	return RamRepo.Insert(HardwareId, obj.HardwareRamItems...)
}

func (obj ExecuteModify) executeHardwareDisks(HardwareId string) (err error) {
	DiskRepo := hardware_repo.HardwareDiskRepo{}

	if err = DiskRepo.Delete(HardwareId); nil != err {
		return err
	}

	for i := range obj.HardwareDiskItems {
		obj.HardwareDiskItems[i].HardwareId = HardwareId
	}


	return DiskRepo.Insert(HardwareId, obj.HardwareDiskItems...)
}

func (obj ExecuteModify) executeHardwareNics(HardwareId string) (err error) {
	NicRepo := hardware_repo.HardwareNicRepo{}

	if err = NicRepo.Delete(HardwareId); nil != err {
		return err
	}

	for i := range obj.HardwareNicItems {
		obj.HardwareNicItems[i].HardwareId = HardwareId
	}

	return NicRepo.Insert(HardwareId, obj.HardwareNicItems...)
}

func (obj ExecuteModify) executeHardwareRaids(HardwareId string) (err error) {
	RaidRepo := hardware_repo.HardwareRaidRepo{}

	if err = RaidRepo.Delete(HardwareId); nil != err {
		return err
	}

	for i := range obj.HardwareRaidItems {
		obj.HardwareRaidItems[i].HardwareId = HardwareId
	}

	return RaidRepo.Insert(HardwareId, obj.HardwareRaidItems...)
}

func (obj ExecuteModify) executeHardwarePsus(HardwareId string) (err error) {
	PsuRepo := hardware_repo.HardwarePsuRepo{}

	if err = PsuRepo.Delete(HardwareId); nil != err {
		return err
	}

	for i := range obj.HardwarePsuItems {
		obj.HardwarePsuItems[i].HardwareId = HardwareId
	}

	return PsuRepo.Insert(HardwareId, obj.HardwarePsuItems...)
}

func (obj ExecuteModify) executeHardwareMnts(HardwareId string) (err error) {
	MntRepo := hardware_repo.HardwareManagementRepo{}

	if err = MntRepo.Delete(HardwareId); nil != err {
		return err
	}

	for i := range obj.HardwareMntItems {
		obj.HardwareMntItems[i].HardwareId = HardwareId
	}

	return MntRepo.Insert(HardwareId, obj.HardwareMntItems...)
}

//
//func (obj *ExecuteModify) New(c *gin.Context) {
//	obj.makeServer(c)
//	obj.makeTags(c)
//	obj.initServerIpsByPostForm(c)
//	obj.SwitchConnArr = parseSwitchConnection(c)
//	err := repo.ServerRepo{}.Update(obj.Server)
//	obj.ExecuteIpServer()
//	obj.Tags = parseTagsByPostForm(c)
//	obj.InsertTags()
//
//	if nil != obj.SwitchConnArr {
//		obj.InsertSwitchConnArr()
//	}
//}
//
//func (obj *ExecuteModify) makeServer(c *gin.Context) {
//	p := parser{c}
//	obj.Server = p.server()
//	obj.Server.DC = p.dc()
//	obj.Server.Rack = p.rack()
//	obj.Server.UStart = p.u_start()
//	obj.Server.UEnd = p.uend()
//	obj.PortType = p.port_type()
//	obj.ServerStatus = p.server_state()
//	obj.IpAddrs = p.toIps()
//}
//
//func (obj *ExecuteModify) makeTags(c *gin.Context) {
//	p := parser{c}
//	obj.Tags = p.tag_splice()
//}
//
//func (obj *ExecuteModify) initServerByPostForm(c *gin.Context) {
//	obj.Server.Id = c.PostForm("txtIdServer")
//	obj.Server.DC.Id = c.PostForm("cbDCId")
//	obj.Server.Rack = obj.ParseRackByPostForm(c)
//	obj.Server.UStart = obj.ParseRackUnitStartByPostForm(c)
//	obj.Server.UEnd = obj.ParseRackUnitEndByPostForm(c)
//	obj.Server.SSD = c.PostForm("txtSSD")
//	obj.Server.HDD = c.PostForm("txtHDD")
//	obj.Server.PortType.Id = c.PostForm("cbPortTypeId")
//	obj.Server.ServerStatus.Id = c.PostForm("cbServerStatusId")
//	obj.Server.SerialNumber = c.PostForm("txtSerialNumber")
//}
//
//func (obj *ExecuteModify) ParseRackByPostForm(c *gin.Context) entity.Rack {
//	rack := entity.Rack{Description: c.PostForm("txtRack")}
//
//	if !rack.IsExistsRackDescription() {
//		rack.GenerateId()
//		rack.Insert()
//	} else {
//		rack.Id = rack.GetIdRack()
//	}
//
//	return rack
//}
//
//func (obj *ExecuteModify) ParseRackUnitStartByPostForm(c *gin.Context) entity.RackUnit {
//	ustart := entity.RackUnit{Description: c.PostForm("txtUStart")}
//
//	if !ustart.IsExistsRackUnitDescription() {
//		ustart.GenerateId()
//		ustart.Insert()
//	} else {
//		ustart.Id = ustart.GetIdRackUnit()
//	}
//
//	return ustart
//}
//
//func (obj *ExecuteModify) ParseRackUnitEndByPostForm(c *gin.Context) entity.RackUnit {
//	uend := entity.RackUnit{Description: c.PostForm("txtUEnd")}
//
//	if !uend.IsExistsRackUnitDescription() {
//		uend.GenerateId()
//		uend.Insert()
//	} else {
//		uend.Id = uend.GetIdRackUnit()
//	}
//
//	return uend
//}
//
//func (obj *ExecuteModify) initServerIpsByPostForm(c *gin.Context) {
//	obj.Server.IpAddrs = parseIpAddressByPostForm(c)
//}
//
//func parseIpAddressByPostForm(c *gin.Context) []entity.IpAddress {
//	ipStr := c.PostForm("txtAllIp")
//	ipStrArr := strings.Split(ipStr, ",")
//	ipAddrs := make([]entity.IpAddress, len(ipStrArr))
//
//	for i := range ipStrArr {
//		values := strings.Split(ipStrArr[i], "-")
//		ipNet := entity.IpNet{Id: values[0]}
//		ipHost := values[1]
//		if "" != ipNet.Id && "" != ipHost {
//			ipAddr := entity.IpAddress{IpNet: ipNet, IpHost: ipHost}
//			ipAddrs[i] = ipAddr
//		}
//	}
//
//	return ipAddrs
//}
//
//func (obj *ExecuteModify) ExecuteIpServer() {
//	repo.ServerIpRepo{}.Delete(obj.Server.Id,
//		obj.Server.IpAddrs...)
//
//	for i := range obj.Server.IpAddrs {
//		err := obj.Server.IpAddrs[i].Insert(obj.Server.Id)
//		obj.Server.UpdateIpHostState(i, "used")
//		if nil != err {
//			panic(err)
//		}
//	}
//}
//
//func (obj *ExecuteModify) InsertTags() {
//	entity.DeleteServerTags(obj.Server.Id)
//	success := entity.InsertServerTags(obj.Server.Id, obj.Tags)
//	if nil != success {
//		obj.Msg = "Can't insert tags of server"
//		panic(success)
//	}
//}
//
//func parseTagsByPostForm(c *gin.Context) []entity.Tag {
//	titles := strings.Split(c.PostForm("txtAllTag"), ",")
//	tags := make([]entity.Tag, len(titles))
//
//	for i := range titles {
//		tags[i] = entity.Tag{Title: titles[i]}
//		tags[i].InitTagId()
//	}
//
//	return tags
//}
//
//func (obj *ExecuteModify) InsertSwitchConnArr() {
//	obj.DeleteSwitchConnectionByServerId()
//	var err error
//
//	for i := range obj.SwitchConnArr {
//		err = obj.SwitchConnArr[i].Insert()
//		if nil != err {
//			panic(err)
//		}
//	}
//}
//
//func parseSwitchConnection(c *gin.Context) []entity.SwitchConnection {
//	switchContent := c.PostForm("txtSwitch")
//	if "" == switchContent {
//		return nil
//	}
//
//	switchIdArr := strings.Split(switchContent, ",")
//
//	switchConnArr := make([]entity.SwitchConnection, len(switchIdArr))
//	serverId := c.PostForm("txtIdServer")
//
//	for i := range switchConnArr {
//		values := strings.Split(switchIdArr[i], "-")
//
//		switchConnArr[i] = entity.SwitchConnection{
//			ServerId:    serverId,
//			SwitchId:    values[0],
//			CableTypeId: values[1],
//			Port:        values[2],
//		}
//	}
//
//	return switchConnArr
//}
//
//func (obj *ExecuteModify) DeleteSwitchConnectionByServerId() {
//	comp := obj.DeleteSwitchConnectionByServerIdComponent()
//	err := database.Delete(comp)
//
//	if nil != err {
//		panic(err)
//	}
//}
//
//func (obj *ExecuteModify) DeleteSwitchConnectionByServerIdComponent() database.DeleteComponent {
//	return database.DeleteComponent{
//		Table:         "SWITCH_CONNECTION",
//		Selection:     "ID_SERVER = ?",
//		SelectionArgs: []string{obj.Server.Id},
//	}
//}

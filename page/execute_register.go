package page

import (
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

type ExecuteRegister struct {
	Msg string
	entity.Server
	Tags []entity.Tag

	hardware.HardwareConfig
	HardwareCpuItems  []hardware.HardwareCpu
	HardwareRamItems  []hardware.HardwareRam
	HardwareDiskItems []hardware.HardwareDisk
	HardwareNicItems  []hardware.HardwareNic
	HardwareRaidItems []hardware.HardwareRaid
	HardwarePsuItems  []hardware.HardwarePsu
	HardwareMntItems  []hardware.HardwareManagement
}

func (obj *ExecuteRegister) New(c *gin.Context) error {
	obj.Server = ParseServerByPostForm(c)
	obj.Server.DC = ParseDCByPostForm(c)
	obj.Server.Rack = ParseRackByPostForm(c)
	obj.Server.UStart = ParseRackUnitStartByPostForm(c)
	obj.Server.UEnd = ParseRackUnitEndByPostForm(c)
	obj.Server.PortType = ParsePortTypeByPostForm(c)
	obj.Server.ServerStatus = ParseServerStateByPostForm(c)
	obj.Server.ServeCustomer = ParseServeCustomerByPostForm(c)

	listIp, err := ParseServerIpAddrs(c)
	if nil != err {
		return err
	}
	obj.IpAddrs = listIp

	redfishIp, err := ParseServerRedfishIpAddr(c)
	if nil != err {
		return err
	}
	obj.RedfishIp = redfishIp


	obj.Tags, err = ParseServerTags(c)
	if nil != err {
		return err
	}

	hwParser := HardwareParser{c}
	obj.HardwareConfig = hwParser.HardwareConfig()
	obj.HardwareConfig.Id = hardware_repo.HardwareConfigRepo{}.GenerateId()
	obj.HardwareCpuItems = hwParser.HardwareCPUArray()
	obj.HardwareRamItems = hwParser.HardwareRAMArray()
	obj.HardwareDiskItems = hwParser.HardwareDiskArray()
	obj.HardwareNicItems = hwParser.HardwareNicArray()
	obj.HardwareRaidItems = hwParser.HardwareRaidArray()
	obj.HardwarePsuItems = hwParser.HardwarePsuArray()
	obj.HardwareMntItems = hwParser.HardwareMntArray()


	return nil



	//err := obj.InsertServer(c)
	//if nil != err {
	//	panic (err)
	//}
	//
	//obj.Tags = obj.parseTagsByPostForm(c)
	//obj.InsertServerTags()
}

func (obj *ExecuteRegister) Execute() (err error) {
	serverRepo := server.ServerRepo{}
	if err = serverRepo.Insert(obj.Server); nil != err {
		return err
	}

	if err = serverRepo.UpdateHardwareConfig(obj.Server.Id, obj.HardwareConfig.Id); nil != err {
		return err
	}

	ipRepo := server.ServerIpRepo{}
	if 0 == len(obj.Server.IpAddrs) {
		if err = ipRepo.InsertNormalIpAddresses(obj.Server.Id, obj.Server.IpAddrs...); nil != err {
			return err
		}
	}

	if err = ipRepo.InsertRedfishIpAddresses(obj.Server.Id, obj.Server.RedfishIp); nil != err {
		return err
	}

	serverTagRepo := server.ServerTagRepo{}
	if err = serverTagRepo.Insert(obj.Server.Id, obj.Tags...); nil != err {
		return err
	}

	return obj.executeHardwareComponents()
}

func ParseServerByPostForm(c *gin.Context) entity.Server {
	serialNumber := c.PostForm("txtSerialNumber")
	dc := ParseDCByPostForm(c)
	rack := ParseRackByPostForm(c)
	uStart := ParseRackUnitStartByPostForm(c)
	uEnd := ParseRackUnitEndByPostForm(c)
	portType := ParsePortTypeByPostForm(c)
	serverState := ParseServerStateByPostForm(c)

	return entity.Server  {
		Id: server.ServerRepo{}.GenerateId(),
		DC: dc,
		Rack: rack,
		UStart: uStart,
		UEnd: uEnd,
		PortType: portType,
		ServerStatus: serverState,
		SerialNumber: serialNumber,
	}
}

func ParseDCByPostForm(c *gin.Context) entity.DataCenter {
		return entity.DataCenter {Id: c.PostForm("cbDC")}
}

func ParseRackByPostForm(c *gin.Context) entity.Rack {
	rack := entity.Rack{Description: c.PostForm("txtRack")}
	rackRepo := server.RackRepo{}

	if !rack.IsExistsRackDescription() {
		rack.Id = rackRepo.GenerateId()
		rackRepo.Insert(rack)
	} else {
		rack.Id = rackRepo.FetchId(rack.Description)
	}

	return rack
}

func ParseRackUnitStartByPostForm(c *gin.Context) entity.RackUnit {
	rackUnit := entity.RackUnit{Description: c.PostForm("txtUStart")}
	rackUnitRepo := server.RackUnitRepo{}

	if !rackUnit.IsExistsRackUnitDescription() {
		rackUnit.Id = rackUnitRepo.GenerateId()
		rackUnitRepo.Insert(rackUnit)
	} else {
		rackUnit.Id = rackUnitRepo.FetchId(rackUnit.Description)
	}

	return rackUnit
}

func ParseRackUnitEndByPostForm(c *gin.Context) entity.RackUnit {
	rackUnit := entity.RackUnit{Description: c.PostForm("txtUEnd")}
	rackUnitRepo := server.RackUnitRepo{}

	if !rackUnit.IsExistsRackUnitDescription() {
		rackUnit.Id = rackUnitRepo.GenerateId()
		rackUnitRepo.Insert(rackUnit)
	} else {
		rackUnit.Id = rackUnitRepo.FetchId(rackUnit.Description)
	}

	return rackUnit
}

func ParsePortTypeByPostForm(c *gin.Context) entity.PortType {
	return entity.PortType {Id: c.PostForm("cbPortType")}
}

func ParseServerStateByPostForm(c *gin.Context) entity.ServerStatus {
	return entity.ServerStatus {Id: c.PostForm("cbServerState")}
}

func ParseServeCustomerByPostForm(c *gin.Context) entity.ServeCustomer {
	return entity.ServeCustomer{
		Id: c.PostForm("cbServeCustomer"),
	}
}

func ParseServerIpAddrs(c *gin.Context) ([]entity.IpAddress, error) {
	rawContent := c.PostForm("txtAllIp")
	if "" == rawContent {
		return nil, nil
	}

	rawIp := strings.Split(rawContent, ",")
	listIp := make([]entity.IpAddress, len(rawIp))

	for i := range listIp {
		values := strings.Split(rawIp[i], ".")
		if len(values) != 4 {
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

func ParseServerRedfishIpAddr(c *gin.Context) (entity.IpAddress, error) {
	rawContent := c.PostForm("txtRedfishIp")
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

func ParseServerTags(c *gin.Context) ([]entity.Tag, error) {
	tagRepo := server.TagRepo{}

	var err error
	titles := strings.Split(c.PostForm("txtAllTag"), ",")
	tags := make([]entity.Tag, len(titles))

	for i := range titles {
		tags[i] = entity.Tag{Title: titles[i]}
		tags[i].TagId, err = tagRepo.IdOf(tags[i].Title)

		if nil != err {
				return nil,err
			}
	}

		return tags, nil
}

func (obj ExecuteRegister) executeHardwareComponents() (err error) {
	if err = obj.executeHardwareConfig(); nil != err {
		return err
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

func (obj ExecuteRegister) executeHardwareConfig() (err error) {
	return hardware_repo.HardwareConfigRepo{}.Insert(obj.HardwareConfig)
}

func (obj ExecuteRegister) executeHardwareCPUs(HardwareId string) (err error) {
	for i := range obj.HardwareCpuItems {
		obj.HardwareCpuItems[i].HardwareId = obj.HardwareConfig.Id
	}
	
	cpuRepo := hardware_repo.HardwareCPURepo{}
	return cpuRepo.Insert(HardwareId, obj.HardwareCpuItems...)
}

func (obj ExecuteRegister) executeHardwareRams(HardwareId string) (err error) {
	for i := range obj.HardwareRamItems {
		obj.HardwareRamItems[i].HardwareId = obj.HardwareConfig.Id
	}
	
	RamRepo := hardware_repo.HardwareRamRepo{}
	return RamRepo.Insert(HardwareId, obj.HardwareRamItems...)
}

func (obj ExecuteRegister) executeHardwareDisks(HardwareId string) (err error) {
	fmt.Println ("Number of disk: ", len(obj.HardwareDiskItems))
	fmt.Println ("Disks: ", obj.HardwareDiskItems)
	for i := range obj.HardwareDiskItems {
		obj.HardwareDiskItems[i].HardwareId = obj.HardwareConfig.Id
	}
	
	DiskRepo := hardware_repo.HardwareDiskRepo{}
	return DiskRepo.Insert(HardwareId, obj.HardwareDiskItems...)
}

func (obj ExecuteRegister) executeHardwareNics(HardwareId string) (err error) {
	for i := range obj.HardwareNicItems {
		obj.HardwareNicItems[i].HardwareId = obj.HardwareConfig.Id
	}
	
	NicRepo := hardware_repo.HardwareNicRepo{}
	return NicRepo.Insert(HardwareId, obj.HardwareNicItems...)
}

func (obj ExecuteRegister) executeHardwareRaids(HardwareId string) (err error) {
	for i := range obj.HardwareRaidItems {
		obj.HardwareRaidItems[i].HardwareId = obj.HardwareConfig.Id
	}
	
	RaidRepo := hardware_repo.HardwareRaidRepo{}
	return RaidRepo.Insert(HardwareId, obj.HardwareRaidItems...)
}

func (obj ExecuteRegister) executeHardwarePsus(HardwareId string) (err error) {
	for i := range obj.HardwarePsuItems {
		obj.HardwarePsuItems[i].HardwareId = obj.HardwareConfig.Id
	}
	
	PsuRepo := hardware_repo.HardwarePsuRepo{}
	return PsuRepo.Insert(HardwareId, obj.HardwarePsuItems...)
}

func (obj ExecuteRegister) executeHardwareMnts(HardwareId string) (err error) {
	for i := range obj.HardwareMntItems {
		obj.HardwareMntItems[i].HardwareId = obj.HardwareConfig.Id
	}
	
	MntRepo := hardware_repo.HardwareManagementRepo{}
	return MntRepo.Insert(HardwareId, obj.HardwareMntItems...)
}
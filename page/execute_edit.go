package page

import (
	"CURD/database"
	"CURD/entity"
	"CURD/repo/server"
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
	return entity.Rack{
		Description: p.parse("txtRack"),
	}
}

func (p parser) u_start() entity.RackUnit {
	return entity.RackUnit{
		Description: p.parse("txtUStart"),
	}
}

func (p parser) uend() entity.RackUnit {
	return entity.RackUnit{
		Description: p.parse("txtUEnd"),
	}
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

func (p parser) list_ip() []entity.IpAddress {
	raw := p.parse("txtAllIp")
	values := strings.Split(raw, ",")
	return p.list_ip_by(values)
}

func (p parser) list_ip_by(values []string) (list []entity.IpAddress) {
	for i := range values {
		ip := entity.IpAddress{}
		if ip.Parse(values[i]) {
			list = append(list, ip)
		}
	}
	return
}

func (p parser) tag_splice() []entity.Tag {
	raw := p.parse("txtAllTag")
	return p.tag_splice_init(strings.Split(raw, ", "))
}

func (p parser) tag_splice_init(titles []string) (tags []entity.Tag) {
	for i := range titles {
		tags[i].Title = titles[i]
		tags[i].TagId, _ = server.TagRepo{}.IdOf(titles[i])
	}
	return
}

type ExecuteModify struct {
	entity.Server
	Msg           string
	SwitchConnArr []entity.SwitchConnection
	Tags          []entity.Tag
}

type execute = func() error

func (obj *ExecuteModify) New(c *gin.Context) {
	p := parser{c}

	obj.Server = p.server()
	obj.Server.DC = p.dc()
	obj.Server.Rack = p.rack()
	obj.Server.UStart = p.u_start()
	obj.Server.UEnd = p.uend()
	obj.Server.PortType = p.port_type()
	obj.Server.ServerStatus = p.server_state()
	obj.Server.IpAddrs = p.list_ip()
	obj.Tags = p.tag_splice()
}

func (obj ExecuteModify) Execute() error {
	return obj.executes([]execute{
		obj.updateServer,
		obj.executeListIpServer,
		obj.executeTag,
		obj.executeSwitchConnection,
	})
}

func (obj ExecuteModify) executes(methods []execute) (err error) {
	for i := 0; i < len(methods) && nil != err; i++ {
		err = methods[i]()
	}
	return err
}

func (obj ExecuteModify) updateServer() error {
	return server.ServerRepo{}.Update(obj.Server)
}

func (obj ExecuteModify) executeListIpServer() (err error) {
	r := server.ServerIpRepo{}
	err = r.Delete(obj.Server.Id, obj.Server.IpAddrs...)
	if nil == err {
		err = r.Insert(r.MakeIcomp(obj.Server.Id, obj.Server.IpAddrs...))
	}
	return
}

func (obj ExecuteModify) executeTag() (err error) {
	r := server.ServerTagRepo{}
	err = r.Delete(r.MakeDelteComp(obj.Server.Id))
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

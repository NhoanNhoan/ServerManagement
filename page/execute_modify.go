package page

import (
	"strings"

	"CURD/entity"
	"CURD/database"

	"github.com/gin-gonic/gin"
)

type ExecuteModify struct {
	entity.Server
	SwitchConnArr []entity.SwitchConnection
	Msg string
	Tags	[]entity.Tag
}

func (obj *ExecuteModify) New(c *gin.Context) {
	obj.initServerByPostForm(c)
	obj.initServerIpsByPostForm(c)
	obj.SwitchConnArr = parseSwitchConnection(c)
	obj.Msg = entity.UpdateServer(obj.Server)
	obj.ExecuteIpServer()
	obj.Tags = parseTagsByPostForm(c)
	obj.InsertTags()

	if nil != obj.SwitchConnArr {
		obj.InsertSwitchConnArr()
	}
}

func (obj *ExecuteModify) initServerByPostForm(c *gin.Context) {
	obj.Server.Id = c.PostForm("txtIdServer")
	obj.Server.DC.Id = c.PostForm("cbDCId")
	obj.Server.Rack.Id = c.PostForm("cbRackId")
	obj.Server.UStart.Id = c.PostForm("cbUStartId")
	obj.Server.UEnd.Id = c.PostForm("cbUEndId")
	obj.Server.SSD = c.PostForm("txtSSD")
	obj.Server.HDD = c.PostForm("txtHDD")
	obj.Server.PortType.Id = c.PostForm("cbPortTypeId")
	obj.Server.ServerStatus.Id = c.PostForm("cbServerStatusId")
	obj.Server.SerialNumber = c.PostForm("txtSerialNumber")
}

func (obj *ExecuteModify) initServerIpsByPostForm(c *gin.Context) {
	obj.Server.IpAddrs = parseIpAddressByPostForm(c)
}

func parseIpAddressByPostForm(c *gin.Context) []entity.IpAddress {
	ipStr := c.PostForm("txtAllIp")
	ipStrArr := strings.Split(ipStr, ",")
	ipAddrs := make([]entity.IpAddress, len(ipStrArr))

	for i := range ipStrArr {
		values := strings.Split(ipStrArr[i], "-")
		ipNet := entity.IpNet{Id: values[0]}
		ipHost := values[1]
		if "" != ipNet.Id && "" != ipHost {
			ipAddr := entity.IpAddress {IpNet: ipNet, IpHost: ipHost}
			ipAddrs[i] = ipAddr
		}
	}

	return ipAddrs
}

func (obj *ExecuteModify) ExecuteIpServer() {
	obj.Server.DeleteAllIp()

	for i := range obj.Server.IpAddrs {
		err := obj.Server.IpAddrs[i].Insert(obj.Server.Id)
		obj.Server.UpdateIpHostState(i)
		if nil != err {
			panic(err)
		}	
	}
}

func (obj *ExecuteModify) InsertTags() {
	entity.DeleteServerTags(obj.Server.Id)
	success := entity.InsertServerTags(obj.Server.Id, obj.Tags)
	if nil != success {
		obj.Msg = "Can't insert tags of server"
		panic (success)
	}
}

func parseTagsByPostForm(c *gin.Context) []entity.Tag {
	titles := strings.Split(c.PostForm("txtAllTag"), ",")
	tags := make([]entity.Tag, len(titles))

	for i := range titles {
		tags[i] = entity.Tag{Title: titles[i]}
		tags[i].InitTagId()
	}

	return tags
}

func (obj *ExecuteModify) InsertSwitchConnArr() {
	obj.DeleteSwitchConnectionByServerId()
	var err error

	for i := range obj.SwitchConnArr {
		err = obj.SwitchConnArr[i].Insert()
		if nil != err {
			panic (err)
		}
	}
}

func parseSwitchConnection(c *gin.Context) []entity.SwitchConnection {
	switchContent := c.PostForm("txtSwitch")
	if "" == switchContent {
		return nil
	}

	switchIdArr := strings.Split(switchContent, ",")

	switchConnArr := make([]entity.SwitchConnection, len(switchIdArr))
	serverId := c.PostForm("txtIdServer")

	for i := range switchConnArr {
		values := strings.Split(switchIdArr[i], "-")

		switchConnArr[i] = entity.SwitchConnection {
			ServerId: serverId,
			SwitchId: values[0],
			CableTypeId: values[1],
			Port: values[2],
		}
	}

	return switchConnArr
}

func (obj *ExecuteModify) DeleteSwitchConnectionByServerId() {
	comp := obj.DeleteSwitchConnectionByServerIdComponent()
	err := database.Delete(comp)

	if nil != err {
		panic (err)
	}
}

func (obj *ExecuteModify) DeleteSwitchConnectionByServerIdComponent() database.DeleteComponent {
	return database.DeleteComponent {
		Table: "SWITCH_CONNECTION",
		Selection: "ID_SERVER = ?",
		SelectionArgs: []string {obj.Server.Id},
	}
}
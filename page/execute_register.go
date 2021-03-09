package page

import (
	"strings"

	"CURD/database"
	"CURD/entity"

	"github.com/gin-gonic/gin"
)

type ExecuteRegister struct {
	Msg string
	entity.Server
	Tags []entity.Tag
}

func (obj *ExecuteRegister) New(c *gin.Context) {
	obj.Server = ParseServerByPostForm(c)
	err := obj.InsertServer(c)
	if nil != err {
		panic (err)
	}

	obj.Tags = obj.parseTagsByPostForm(c)
	obj.InsertServerTags()
}

func (obj *ExecuteRegister) InsertServer(c *gin.Context) error {
	err := obj.Server.Insert()

	if nil != err {
		obj.Msg = "Couldn't insert a new value"
		return err
	}

	ipAddr := ParseIpAddressByPostForm(c)
	InsertIpServer(obj.Server.Id, ipAddr)

	if nil != err {
		obj.Msg = "Couldn't insert a new value"
		return err
	}

	obj.Msg = "Success"
	return nil
}

func ParseServerByPostForm(c *gin.Context) entity.Server {
	serialNumber := c.PostForm("txtSerialNumber")
	ssd := c.PostForm("txtSSD")
	hdd := c.PostForm("txtHDD")
	maker := c.PostForm("txtMaker")
	dc := ParseDCByPostForm(c)
	rack := ParseRackByPostForm(c)
	uStart := ParseRackUnitStartByPostForm(c)
	uEnd := ParseRackUnitEndByPostForm(c)
	portType := ParsePortTypeByPostForm(c)
	serverState := ParseServerStateByPostForm(c)

	return entity.Server  {
		Id: generateServerId(),
		DC: dc,
		Rack: rack,
		UStart: uStart,
		UEnd: uEnd,
		SSD: ssd,
		HDD: hdd,
		PortType: portType,
		ServerStatus: serverState,
		Maker: maker,
		SerialNumber: serialNumber,
	}
}

func ParseDCByPostForm(c *gin.Context) entity.DataCenter {
		return entity.DataCenter {Id: c.PostForm("cbDC")}
}

func ParseRackByPostForm(c *gin.Context) entity.Rack {
	return entity.Rack {Id: c.PostForm("cbRack")}
}

func ParseRackUnitStartByPostForm(c *gin.Context) entity.RackUnit {
	return entity.RackUnit {Id: c.PostForm("cbUStart")}
}

func ParseRackUnitEndByPostForm(c *gin.Context) entity.RackUnit {
	return entity.RackUnit {Id: c.PostForm("cbUEnd")}
}

func ParsePortTypeByPostForm(c *gin.Context) entity.PortType {
	return entity.PortType {Id: c.PostForm("cbPortType")}
}

func ParseServerStateByPostForm(c *gin.Context) entity.ServerStatus {
	return entity.ServerStatus {Id: c.PostForm("cbServerState")}
}

func ParseIpNetByPostForm(c *gin.Context) entity.IpNet {
	return entity.IpNet {Id: c.PostForm("cbIpNet")}
}

func generateServerId() string {
	id := database.GeneratePrimaryKey(true,
								true, true, false, "SV", 12)

	for isExistsServerId(id) {
		id = database.GeneratePrimaryKey(true,
								true, true, false, "SV", 12)
	}

	return id
}

func isExistsServerId(ServerId string) bool {
	comp := database.QueryComponent {
		Tables: []string {"SERVER"},
		Columns: []string {"ID"},
		Selection: "ID = ?",
		SelectionArgs: []string {ServerId},
	}
	row, err := database.Query(comp)
	defer row.Close()
	return (nil == err) && row.Next()
}

func ParseIpAddressByPostForm(c *gin.Context) []entity.IpAddress {
	ipStr := c.PostForm("txtAllIp")
	ipStrs := strings.Split(ipStr, ",")
	ipAddrs := make([]entity.IpAddress, len(ipStrs))

	for i := range ipStrs {
		values := strings.Split(ipStrs[i], "-")
		ipNet := entity.IpNet{Id: values[0]}
		ipHost := values[1]
		ipAddr := entity.IpAddress {IpNet: ipNet, IpHost: ipHost}
		ipAddrs[i] = ipAddr
	}

	return ipAddrs
}

func InsertIpServer(ServerId string, IpAddrs []entity.IpAddress) {
	for i := range IpAddrs {
		err := IpAddrs[i].Insert(ServerId)
		if nil != err {
			panic(err)
		}
	}
}

func (obj *ExecuteRegister) InsertServerTags() {
	success := entity.InsertServerTags(obj.Server.Id, obj.Tags)
	if nil != success {
		obj.Msg = "Can't insert tags of server"
	}
}

func (obj *ExecuteRegister) parseTagsByPostForm(c *gin.Context) []entity.Tag {
	titles := strings.Split(c.PostForm("txtAllTag"), ",")
	tags := make([]entity.Tag, len(titles))

	for i := range titles {
		tags[i] = entity.Tag{Title: titles[i]}
		tags[i].InitTagId()
	}

	return tags
}
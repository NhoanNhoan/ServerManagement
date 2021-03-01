package page

import (
	"CURD/entity"
	"CURD/database"

	"github.com/gin-gonic/gin"
)

type UpdateServer struct {
	DCs []entity.DataCenter
	Racks []entity.Rack
	RackUnits []entity.RackUnit
	PortTypes []entity.PortType
	ServerStates []entity.ServerStatus
	Tagged []entity.Tag
	Untagged []entity.Tag
	IpNets []entity.IpNet
	entity.Server
}

func (obj *UpdateServer) New(c *gin.Context) {
	obj.initServerByPostForm(c)
	obj.Server.FetchIpAddrs()

	obj.DCs = entity.GetDCs()
	obj.Racks = entity.GetRacks()
	obj.RackUnits = entity.GetRackUnits()
	obj.PortTypes = entity.GetPortTypes()
	obj.ServerStates = entity.GetServerStates()
	obj.IpNets = entity.GetIpNets()

	obj.FetchTagged()
	obj.FetchUntagged()
}

func (obj *UpdateServer) initServerByPostForm(c *gin.Context) {
	obj.Server.Id = c.PostForm("txtIdServer")
	obj.Server.DC.Id = c.PostForm("txtDCId")
	obj.Server.Rack.Id = c.PostForm("txtRackId")
	obj.Server.UStart.Id = c.PostForm("txtUStartId")
	obj.Server.UEnd.Id = c.PostForm("txtUEndId")
	obj.Server.PortType.Id = c.PostForm("txtPortTypeId")
	obj.Server.ServerStatus.Id = c.PostForm("txtServerStatusId")
	obj.Server.Maker = c.PostForm("txtMaker")
	obj.Server.NumDisks = c.PostForm("txtNumDisks")
	obj.Server.SerialNumber = c.PostForm("txtSerialNumber")
}

func (obj *UpdateServer) FetchTagged() {
	comp := obj.makeTaggedQueryComponent()
	obj.Tagged = entity.FetchTags(comp)
}

func (obj *UpdateServer) FetchUntagged() {
	comp := obj.makeUntaggedQueryComponent()
	obj.Untagged = entity.FetchTags(comp)
}

func (obj *UpdateServer) makeTaggedQueryComponent() database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"TAG", "SERVER_TAG"},
		Columns: []string {"TAG.TAGID", "TAG.TITLE"},
		Selection: "SERVER_TAG.SERVERID = ? AND " +
				"SERVER_TAG.TAGID = TAG.TAGID",
		SelectionArgs: []string {obj.Server.Id},
	}
}

func (obj *UpdateServer) makeUntaggedQueryComponent() database.QueryComponent {
	selection := "TAGID NOT IN (SELECT S.TAGID " +
				"FROM SERVER_TAG AS S " +
				"WHERE S.SERVERID = ?)"
	return database.QueryComponent {
		Tables: []string {"TAG"},
		Columns: []string {"TAGID", "TITLE"},
		Selection: selection,
		SelectionArgs: []string {obj.Server.Id},
	}
}
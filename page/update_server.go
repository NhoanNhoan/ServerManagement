package page

import (
	"CURD/database"
	"CURD/entity"
	"github.com/gin-gonic/gin"
)

type UpdateServer struct {
	DCs []entity.DataCenter
	Racks []entity.Rack
	RackUnits []entity.RackUnit
	PortTypes []entity.PortType
	ServerStates []entity.ServerStatus
	SwitchArr	[]entity.Switch
	Cables []entity.CableType
	Tagged []entity.Tag
	Untagged []entity.Tag
	IpNets []entity.IpNet
	entity.Server
	ConnectedSwitches []SwitchInfo
}

func (obj *UpdateServer) New(serverId string) {
	obj.Server.New(serverId)
	obj.Server.FetchIpAddrs()


	obj.DCs = entity.GetDCs()
	obj.Racks = entity.GetRacks()
	obj.RackUnits = entity.GetRackUnits()
	obj.PortTypes = entity.GetPortTypes()
	obj.ServerStates = entity.GetServerStates()
	obj.FetchSwitchArr()
	obj.Cables = entity.QueryAllCabs()
	obj.IpNets = entity.GetIpNets()

	obj.FetchTagged()
	obj.FetchUntagged()
	err := obj.initConnectedSwitches(obj.Server.Id)
	if nil != err {
		panic (err)
	}
}

func (obj *UpdateServer) initConnectedSwitches(ServerId string) error {
	comp := obj.queryConnectedSwitchesComp(ServerId)
	rows, err := database.Query(comp)
	defer rows.Close()

	var switchInfo SwitchInfo
	for rows.Next() && nil == err {
		err = switchInfo.ParseFromRow(rows)
		switchInfo.Switch.FetchIpAddrs()
		obj.ConnectedSwitches = append(obj.ConnectedSwitches, 
									switchInfo)
	}

	return err
}

func (obj *UpdateServer) queryConnectedSwitchesComp(ServerId string) database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {
			"SWITCH AS SW",
			"SWITCH_CONNECTION AS SC",
			"DC AS D",
			"RACK AS R",
			"RACK_UNIT AS USTART",
			"RACK_UNIT AS UEND",
			"CABLE_TYPE AS CT",
		},

		Columns: []string {
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

		SelectionArgs: []string {ServerId},
	}
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
	obj.Server.SSD = c.PostForm("txtSSD")
	obj.Server.HDD = c.PostForm("txtHDD")
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

func (obj *UpdateServer) FetchSwitchArr() {
	comp := obj.fetchSwitchComp()
	rows, err := database.Query(comp)
	defer rows.Close()

	var sw entity.Switch
	for rows.Next() && nil == err {
		err = rows.Scan(&sw.Id,
					&sw.DC.Description,
					&sw.Rack.Description,
					&sw.UStart.Description,
					&sw.UEnd.Description)
		obj.SwitchArr = append(obj.SwitchArr, sw)
	}

	if nil != err {
		panic (err)
	}
}

func (obj *UpdateServer) fetchSwitchComp() database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"SWITCH AS S",
						"DC AS D",
						"RACK AS R",
						"RACK_UNIT AS USTART",
						"RACK_UNIT AS UEND",
					},
		Columns: []string {
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
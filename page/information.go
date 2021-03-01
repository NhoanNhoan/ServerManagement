package page

import (
	"CURD/database"
	"CURD/entity"
)

type Information struct {
	entity.Server
	entity.Switch
	Tagged []entity.Tag
	Untagged []entity.Tag
}

func (s *Information) FetchServerByIp(IpAddr string) {
	comp := s.makeQueryServerByIpComp(IpAddr)
	var err error
	s.Server, err = entity.FetchServer(comp)
	s.Server.FetchServices()
	if nil != err {
		panic (err)
	}
}

func (s *Information) makeQueryServerByIpComp(IpAddr string) database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"SERVER AS S",
						"IP_NET",
						"IP_SERVER",
						"DC AS D",
						"RACK AS R",
						"RACK_UNIT AS USTART",
						"RACK_UNIT AS UEND",
						"PORT_TYPE AS PT",
						"SERVER_STATUS AS SS",
						"STATUS_ROW AS SR",
		},
		Columns: []string {
			"S.ID",
			"D.ID",
			"D.DESCRIPTION",
			"R.ID",
			"R.DESCRIPTION",
			"USTART.ID", 
			"USTART.DESCRIPTION",
			"UEND.ID",
			"UEND.DESCRIPTION",
			"S.NUM_DISKS",
			"S.MAKER",
			"PT.ID",
			"PT.DESCRIPTION",
			"S.SERIAL_NUMBER",
			"SS.ID",
			"SS.DESCRIPTION",
		},
		Selection: "S.ID = IP_SERVER.ID_SERVER AND " +
					"IP_SERVER.ID_IP_NET = IP_NET.ID AND " +
					"? = IP_NET.VALUE || IP_SERVER.IP_HOST AND " +
					"S.ID_DC = D.ID AND " +
					"S.ID_RACK = R.ID AND " + 
					"S.ID_U_START = USTART.ID AND " + 
					"S.ID_U_END = UEND.ID AND " + 
					"SR.DESCRIPTION = ? AND S.ID_STATUS_ROW = SR.ID AND " +
					"S.ID_PORT_TYPE = PT.ID AND " +
					"S.ID_SERVER_STATUS = SS.ID",
		SelectionArgs: []string {IpAddr, "available"},
	}
}

func (obj *Information) Prepare(IpAddr string) {
	obj.FetchServerByIp(IpAddr)
	IdSwitch := obj.findSwitchId(obj.Server.Id)

	if "" != IdSwitch {
		obj.initSwitch(IdSwitch)
	}

	obj.FetchTagged()
	obj.FetchUntagged()
	obj.Server.FetchIpAddrs()
	obj.Server.FetchEvents()
	obj.Server.FetchServices()
}

func (obj *Information) New(IdServer string) {
	obj.initServer(IdServer)
	IdSwitch := obj.findSwitchId(IdServer)

	if "" != IdSwitch {
		obj.initSwitch(IdSwitch)
	}

	obj.FetchTagged()
	obj.FetchUntagged()
}

func (obj *Information) initServer(IdServer string) {
	err := obj.Server.New(IdServer)
	obj.Server.FetchIpAddrs()
	obj.Server.FetchServices()
	obj.Server.FetchEvents()

	if nil != err {
		panic (err)
	}
}

func (obj *Information) initSwitch(IdSwitch string) {
	obj.Switch.New(IdSwitch)
	obj.Switch.FetchIpAddrs()
}

func (obj *Information) findSwitchId(IdServer string) string {
	component := obj.makeFindSwitchIdQueryComponent(IdServer)
	rows, err := database.Query(component)
	defer rows.Close()

	var IdSwitch string

	if nil != err {
		return IdSwitch
	}

	if rows.Next() && nil == rows.Scan(&IdSwitch) {
		return IdSwitch
	}

	return IdSwitch
}

func (obj *Information) makeFindSwitchIdQueryComponent(IdServer string) database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"SWITCH_CONNECTION"},
		Columns: []string {"ID_SWITCH"},
		Selection: "ID_SERVER = ?",
		SelectionArgs: []string {IdServer},
		GroupBy: "",
		Having: "",
		OrderBy: "",
		Limit: "",
	}
}

func (obj *Information) FetchTagged() {
	comp := obj.makeTaggedQueryComponent()
	obj.Tagged = entity.FetchTags(comp)
}

func (obj *Information) FetchUntagged() {
	comp := obj.makeUntaggedQueryComponent()
	obj.Untagged = entity.FetchTags(comp)
}

func (obj *Information) makeTaggedQueryComponent() database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"TAG", "SERVER_TAG"},
		Columns: []string {"TAG.TAGID", "TAG.TITLE"},
		Selection: "SERVER_TAG.SERVERID = ? AND " +
				"SERVER_TAG.TAGID = TAG.TAGID",
		SelectionArgs: []string {obj.Server.Id},
	}
}

func (obj *Information) makeUntaggedQueryComponent() database.QueryComponent {
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
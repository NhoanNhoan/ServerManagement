package page

import (
	"CURD/database"
	"CURD/entity"
)

type Information struct {
	entity.Server
	entity.Switch
}

func (obj *Information) New(IdServer string) {
	obj.initServer(IdServer)

	IdSwitch := obj.findSwitchId(IdServer)

	if "" != IdSwitch {
		obj.initSwitch(IdSwitch)
	}
}

func (obj *Information) initServer(IdServer string) {
	err := obj.Server.New(IdServer)
	obj.Server.FetchIpAddrs()
	obj.Server.FetchServices()

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
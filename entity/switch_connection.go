package entity

import (
	"CURD/database"
)

type SwitchConnection struct {
	Id string
	SwitchId string
	ServerId string
	CableTypeId string
	Port string
}

func (obj *SwitchConnection) Insert() error {
	obj.GenerateId()
	comp := obj.insertComponent()
	return database.Insert(comp)
}

func (obj *SwitchConnection) GenerateId() {
	obj.Id = database.GeneratePrimaryKey(true, 
						true, true, false, "SC", 6)

	for obj.Exists() {
		obj.Id = database.GeneratePrimaryKey(true, 
						true, true, false, "SC", 6)
	}
}

func (obj *SwitchConnection) insertComponent() icomp {
	return icomp {
		Table: "SWITCH_CONNECTION",
		Columns: []string {"ID", "ID_SWITCH", "ID_SERVER", "ID_CABLE_TYPE", "PORT"},
		Values: [][]string { []string {obj.Id, obj.SwitchId, obj.ServerId, obj.CableTypeId, obj.Port}},
	}
}

func (obj *SwitchConnection) Exists() bool {
	comp := obj.existsComp()
	rows, err := database.Query(comp)
	defer rows.Close()
	return (nil == err) && rows.Next()
}

func (obj *SwitchConnection) existsComp() qcomp {
	return qcomp {
		Tables: []string {"SWITCH_CONNECTION"},
		Columns: []string {"ID"},
		Selection: "ID = ?",
		SelectionArgs: []string {obj.Id},
	}
}

func (obj *SwitchConnection) insertComp() icomp {
	return icomp {
		Table: "SWITCH_CONNECTION",
		Columns: []string {"ID", "ID_SWITCH", "ID_SERVER", "ID_CABLE_TYPE", "PORT"},
		Values: [][]string {
			[]string {obj.Id, obj.SwitchId, obj.ServerId, obj.CableTypeId, obj.Port},
		},
	}
}

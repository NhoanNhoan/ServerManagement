package entity

import (
	"CURD/database"

	"strconv"
)

type RackUnit struct {
	Id string
	Description string
}

func (unit *RackUnit) GenerateId() {
	unit.Id = database.GeneratePrimaryKey(true,
		true, true,
		false, "U", 6)

	for unit.IsExistsRackUnitDescription() {
		unit.Id = database.GeneratePrimaryKey(true,
			true, true,
			false, "U", 6)
	}
}

func (unit *RackUnit) IsExistsRackUnitDescription() bool {
	comp := unit.getIdRackUnitComp()
	rows, err := database.Query(comp)
	defer rows.Close()
	return (nil == err) && rows.Next()
}

func (unit *RackUnit) getIdRackUnitComp() qcomp {
	return qcomp {
		Tables: []string {"RACK_UNIT"},
		Columns: []string {"ID"},
		Selection: "DESCRIPTION = ?",
		SelectionArgs: []string {unit.Description},
	}
}

func (unit *RackUnit) GetIdRackUnit() string {
	comp := unit.getIdRackUnitComp()
	rows, err := database.Query(comp)
	defer rows.Close()

	if nil != err {
		panic (err)
	}

	var id string
	if rows.Next() {
		rows.Scan(&id)
	}

	return id
}

func (obj *RackUnit) GenerateNewId() {
	obj.Id = database.GeneratePrimaryKey(true, true, true, false, "RU", 6)

	for obj.IsExistsId() {
		obj.Id = database.GeneratePrimaryKey(true, true, true, false, "RU", 6)
	}
}

func (obj *RackUnit) IsExistsId() bool {
	comp := obj.makeQueryExistsComponent()
	rows, err := database.Query(comp)
	return nil == err && rows.Next()
}

func (obj *RackUnit) makeQueryExistsComponent() database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"Rack_Unit"},
		Columns: []string {"Id"},
		Selection: "Id = ?",
		SelectionArgs: []string {obj.Id},
	}
}

func (obj *RackUnit) Insert() {
	comp := obj.makeInsertRackUnitComp()
	database.Insert(comp)
}

func (obj *RackUnit) makeInsertRackUnitComp() database.InsertComponent {
	return database.InsertComponent {
		Table: "RACK_UNIT",
		Columns: []string {"ID", "DESCRIPTION"},
		Values: [][]string {[]string {obj.Id, obj.Description}},
	}
}

func InsertRackUnits() {
	for i := 0; i < 50; i++ {
		unit := RackUnit{}
		unit.GenerateNewId()
		unit.Description = "U" + strconv.Itoa(i + 1)
		unit.Insert()
	}
}
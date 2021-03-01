package entity

import (
	"database/sql"
	"CURD/database"

	"strconv"
)

type RackUnit struct {
	Id string
	Description string
}

func (unit RackUnit) ToInstance(args ...string) Entity {
	Id, Des := args[0], args[1]
	return RackUnit{Id, Des}
}

func GetRackUnits() []RackUnit {
	comp := database.MakeQueryAll([]string {"RACK_UNIT"},
							[]string {"Id", "description"})
	rackUnits := getEntities(comp, func (rows *sql.Rows) Entity {
		var id, des string
		err := rows.Scan(&id, &des)

		if nil != err {
			panic (err)
		}

		return RackUnit {id, des}
	})

	return toRackUnitSplice(rackUnits)
}

func toRackUnitSplice(entities []Entity) []RackUnit {
	rackUnits := make([]RackUnit, len(entities))

	for i := range entities {
		rackUnits[i] = entities[i].(RackUnit)
	}

	return rackUnits
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
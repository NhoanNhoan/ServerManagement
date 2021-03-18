package entity

import (
	"database/sql"
	"CURD/database"
)

type Rack struct {
	Id string
	Description string
}

func (rack *Rack) GenerateId() {
	rack.Id = database.GeneratePrimaryKey(true,
		true, true,
		false, "R", 6)

	for rack.IsExistsRackDescription() {
		rack.Id = database.GeneratePrimaryKey(true,
			true, true,
			false, "R", 6)
	}
}

func (rack *Rack) IsExistsRackDescription() bool {
	comp := rack.getIdRackComp()
	rows, err := database.Query(comp)
	defer rows.Close()
	return (nil == err) && rows.Next()
}

func (rack *Rack) getIdRackComp() qcomp {
	return qcomp {
		Tables: []string {"RACK"},
		Columns: []string {"ID"},
		Selection: "DESCRIPTION = ?",
		SelectionArgs: []string {rack.Description},
	}
}

func (rack *Rack) GetIdRack() string {
	comp := rack.getIdRackComp()
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

func (rack *Rack) Insert() {
	comp := rack.insertComp()
	err := database.Insert(comp)
	if nil != err {
		panic (err)
	}
}

func (rack *Rack) insertComp() icomp {
	return icomp {
		Table: "RACK",
		Columns: []string {"ID", "DESCRIPTION"},
		Values:[][]string {[]string{rack.Id, rack.Description}},
	}
}

func (rack Rack) ToInstance(args ...string) Entity {
	Id, Des := args[0], args[1]
	return Rack {Id: Id, Description: Des}
}

func GetRacks() []Rack {
	comp := database.MakeQueryAll([]string {"RACK"},
							[]string {"Id", "description"})
	racks := getEntities(comp, func (rows *sql.Rows) Entity {
		var id, des string
		err := rows.Scan(&id, &des)

		if nil != err {
			panic (err)
		}

		return Rack {id, des}
	})

	return toRackSplice(racks)
}

func toRackSplice(entities []Entity) []Rack {
	racks := make([]Rack, len(entities))

	for i := range entities {
		racks[i] = entities[i].(Rack)
	}

	return racks
}
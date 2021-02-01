package entity

import (
	"database/sql"
	"CURD/database"
)

type Rack struct {
	Id string
	Description string
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
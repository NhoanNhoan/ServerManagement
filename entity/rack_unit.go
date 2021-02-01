package entity

import (
	"database/sql"
	"CURD/database"
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
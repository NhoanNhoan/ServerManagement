package entity

import (
	"database/sql"
	"CURD/database"
)

type DataCenter struct {
	Id string
	Description string
}

func (dc DataCenter) ToInstance(args ...string) Entity {
	Id, Des := args[0], args[1]
	return DataCenter {Id, Des}
}

func GetDCs() []DataCenter {
	comp := database.MakeQueryAll([]string {"DC"},
							[]string {"Id", "description"})
	DCs := getEntities(comp, func (rows *sql.Rows) Entity {
		var id, des string
		err := rows.Scan(&id, &des)

		if nil != err {
			panic (err)
		}

		return DataCenter {id, des}
	})

	return toDCSplice(DCs)
}

func toDCSplice(entities []Entity) []DataCenter {
	DCs := make([]DataCenter, len(entities))

	for i := range entities {
		DCs[i] = entities[i].(DataCenter)
	}

	return DCs
}

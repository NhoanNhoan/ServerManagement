package entity

import (
	"database/sql"
	"CURD/database"
)

type PortType struct {
	Id string
	Description string
}

func (obj PortType) ToInstance(args ...string) Entity {
	Id, Des := args[0], args[1]
	return PortType{Id, Des}
}

func GetPortTypes() []PortType {
	comp := database.MakeQueryAll([]string {"PORT_TYPE"},
							[]string {"Id", "description"})
	portTypes := getEntities(comp, func (rows *sql.Rows) Entity {
		var id, des string
		err := rows.Scan(&id, &des)

		if nil != err {
			panic (err)
		}

		return PortType {id, des}
	})

	return toPortTypeSlice(portTypes)
}

func toPortTypeSlice(entities []Entity) []PortType {
	portTypes := make([]PortType, len(entities))

	for i := range entities {
		portTypes[i] = entities[i].(PortType)
	}

	return portTypes
}
package entity

import (
	"database/sql"
	"CURD/database"
)

type ServerStatus struct {
	Id string
	Description string
}

func (status ServerStatus) ToInstance(args ...string) Entity {
	Id, Des := args[0], args[1]
	return ServerStatus{Id, Des}
}

func GetServerStates() []ServerStatus {
	comp := database.MakeQueryAll([]string {"SERVER_STATUS"},
							[]string {"Id", "description"})
	states := getEntities(comp, func (rows *sql.Rows) Entity {
		var id, des string
		err := rows.Scan(&id, &des)

		if nil != err {
			panic (err)
		}

		return ServerStatus {id, des}
	})

	return toServerStatusSplice(states)
}

func toServerStatusSplice(entities []Entity) []ServerStatus {
	states := make([]ServerStatus, len(entities))

	for i := range entities {
		states[i] = entities[i].(ServerStatus)
	}

	return states
}
package entity

import (
	"database/sql"
)

type ServerEvent struct {
	Id string
	IdServer string
	Description string
	OccurAt string
}

func MakeEventByRow(row *sql.Rows) (error, *ServerEvent) {
	var event ServerEvent
	return row.Scan(&event.Id, 
								&event.IdServer, 
								&event.Description, 
								&event.OccurAt), 
		&event
}

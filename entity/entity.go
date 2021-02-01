package entity

import (
	"database/sql"
	"CURD/database"
)

type Entity interface {
	ToInstance(args ...string) Entity
}

func getEntities(comp database.QueryComponent, 
			makeEntity func (*sql.Rows) Entity) ([]Entity) {
	rows, err := database.Query(comp)

	if nil != err {
		panic (err)
	}

	return entitiesByRows(rows, makeEntity)
}

func entitiesByRows(rows *sql.Rows, 
				makeEntity func (*sql.Rows) Entity, 
				args ...*string) []Entity {
	entities := make([]Entity, 0)

	for rows.Next() {
		newEntity := makeEntity(rows)
		entities = append(entities, newEntity)
	}

	return entities
}

package server

import (
	"CURD/database"
	"database/sql"
)

type SqliteRepo struct {

}

func (sqlite SqliteRepo) Insert(comp icomp) error {
	return database.Insert(comp)
}

func (sqlite SqliteRepo) Update(comp ucomp) error {
	return database.Update(comp)
}

func (sqlite SqliteRepo) Delete(comp dcomp) error {
	return database.Delete(comp)
}

func (sqlite SqliteRepo) Query(comp qcomp,
	makeEntity func() interface{},
	scan func(e interface{}, row *sql.Rows) (interface{}, error)) ([]interface{}, error) {

	rows, err := database.Query(comp)
	if nil != err {
		return nil, err
	}

	defer rows.Close()

	entities := make([]interface{}, 0)
	for rows.Next() {
		e := makeEntity()

		if e, err = scan(e, rows); nil != err {
			entities = nil
			break
		}

		entities = append(entities, e)
	}

	return entities, err
}
package server

import (
	"CURD/entity"
	"database/sql"
)

type ServerStatusRepo struct {
	SqliteRepo
}

func (repo ServerStatusRepo) Fetch(comp qcomp,
	scan func(obj interface{}, rows *sql.Rows) (interface{}, error)) ([]entity.ServerStatus, error) {
	makeServerStatus := func() interface{} { return entity.ServerStatus{} }
	entities, err := repo.SqliteRepo.Query(comp, makeServerStatus, scan)
	if nil != err {
		return nil, err
	}

	states := make([]entity.ServerStatus, len(entities))
	for i := range states {
		states[i] = entities[i].(entity.ServerStatus)
	}

	return states, nil
}

func (ServerState ServerStatusRepo) FetchAll() ([]entity.ServerStatus, error) {
	comp := qcomp{
		Tables: []string {"SERVER_STATUS"},
		Columns: []string {"ID", "DESCRIPTION"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		center := obj.(entity.ServerStatus)
		err := row.Scan(&center.Id, &center.Description)
		return center, err
	}

	return ServerState.Fetch(comp, scan)
}


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

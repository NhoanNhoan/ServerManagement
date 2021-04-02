package server

import (
	"CURD/entity"
	"database/sql"
)

type RackRepo struct {
	SqliteRepo
}

func (r RackRepo) Fetch(comp qcomp,
	scan func(obj interface{}, rows *sql.Rows) (interface{}, error)) ([]entity.Rack, error) {
	makeRack := func() interface{} { return entity.Rack{} }
	entities, err := r.SqliteRepo.Query(comp, makeRack, scan)
	if nil != err {
		return nil, err
	}

	racks := make([]entity.Rack, len(entities))
	for i := range racks {
		racks[i] = entities[i].(entity.Rack)
	}

	return racks, nil
}

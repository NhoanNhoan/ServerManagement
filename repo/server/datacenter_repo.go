package server

import (
	"CURD/entity"
	"database/sql"
)

type DCRepo struct {
	SqliteRepo
}

func (dc DCRepo) Fetch(comp qcomp,
	scan func(obj interface{}, rows *sql.Rows) (interface{}, error)) ([]entity.DataCenter, error) {
	makeDC := func() interface{} { return entity.DataCenter{} }
	entities, err := dc.SqliteRepo.Query(comp, makeDC, scan)
	if nil != err {
		return nil, err
	}

	DCs := make([]entity.DataCenter, len(entities))
	for i := range DCs {
		DCs[i] = entities[i].(entity.DataCenter)
	}

	return DCs, nil
}

func (dc DCRepo) FetchAll() ([]entity.DataCenter, error) {
	comp := qcomp{
		Tables: []string {"DC"},
		Columns: []string {"ID", "DESCRIPTION"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		center := obj.(entity.DataCenter)
		err := row.Scan(&center.Id, &center.Description)
		return center, err
	}

	return dc.Fetch(comp, scan)
}

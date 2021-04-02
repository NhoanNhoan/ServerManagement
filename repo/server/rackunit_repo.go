package server

import (
	"CURD/entity"
	"database/sql"
)

type RackUnitRepo struct {
	SqliteRepo
}

func (unit RackUnitRepo) Fetch(comp qcomp,
	scan func(obj interface{}, rows *sql.Rows) (interface{}, error)) ([]entity.RackUnit, error) {
	makeRackUnit := func() interface{} { return entity.RackUnit{} }
	entities, err := unit.SqliteRepo.Query(comp, makeRackUnit, scan)
	if nil != err {
		return nil, err
	}

	units := make([]entity.RackUnit, len(entities))
	for i := range units {
		units[i] = entities[i].(entity.RackUnit)
	}

	return units, nil
}

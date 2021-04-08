package server

import (
	"CURD/entity"
	"database/sql"
)

type PortTypeRepo struct {
	SqliteRepo
}

func (repo PortTypeRepo) Fetch(comp qcomp,
	scan func(obj interface{}, rows *sql.Rows) (interface{}, error)) ([]entity.PortType, error) {
	makePortType := func() interface{} { return entity.PortType{} }
	entities, err := repo.SqliteRepo.Query(comp, makePortType, scan)
	if nil != err {
		return nil, err
	}

	portTypes := make([]entity.PortType, len(entities))
	for i := range portTypes {
		portTypes[i] = entities[i].(entity.PortType)
	}

	return portTypes, nil
}
func (portType PortTypeRepo) FetchAll() ([]entity.PortType, error) {
	comp := qcomp{
		Tables: []string {"PORT_TYPE"},
		Columns: []string {"ID", "DESCRIPTION"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		center := obj.(entity.PortType)
		err := row.Scan(&center.Id, &center.Description)
		return center, err
	}

	return portType.Fetch(comp, scan)
}

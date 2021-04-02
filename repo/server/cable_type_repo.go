package server

import (
	"CURD/entity"
	"database/sql"
)

type CableTypeRepo struct {
	SqliteRepo
}

func (repo CableTypeRepo) Fetch(comp qcomp,
	scan func(obj interface{}, row *sql.Rows) (interface{}, error)) ([]entity.CableType, error) {
	makeCable := func() interface{} { return entity.CableType{} }
	entities, err := repo.SqliteRepo.Query(comp, makeCable, scan)
	if nil != err {
		return nil, err
	}

	cables := make([]entity.CableType, len(entities))
	for i := range cables {
		cables[i] = entities[i].(entity.CableType)
	}

	return cables, nil
}

func (repo CableTypeRepo) FetchById(CabId string) entity.CableType {
	comp := qcomp{
		Tables: []string {"CABLE"},
		Columns: []string {"NAME", "SIGN_PORT"},
		Selection: "ID = ?",
		SelectionArgs: []string {CabId},
	}

	makeCableType := func() interface{} {return entity.CableType{}}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		cab := obj.(entity.CableType)
		err := row.Scan(&cab.Name, &cab.SignPort)
		return cab, err
	}

	entities, err := repo.SqliteRepo.Query(comp, makeCableType, scan)
	if nil != err {
		return entity.CableType{}
	}

	if len(entities) > 0 {
		cab := entities[0].(entity.CableType)
		cab.Id = CabId
		return cab
	}

	return entity.CableType{}
}

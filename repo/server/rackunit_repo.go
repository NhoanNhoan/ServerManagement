package server

import (
	"CURD/database"
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

func (RackUnit RackUnitRepo) FetchAll() ([]entity.RackUnit, error) {
	comp := qcomp{
		Tables: []string {"Rack_Unit"},
		Columns: []string {"ID", "DESCRIPTION"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		center := obj.(entity.RackUnit)
		err := row.Scan(&center.Id, &center.Description)
		return center, err
	}

	return RackUnit.Fetch(comp, scan)
}

func (RackUnit RackUnitRepo) FetchId(Description string) string {
	comp := qcomp {
		Tables: []string {"Rack_Unit"},
		Columns: []string {"ID"},
		Selection: "DESCRIPTION  =?",
		SelectionArgs: []string {Description},
	}

	scan := func (obj interface{}, row *sql.Rows) (interface{}, error) {
		RackUnit := obj.(entity.RackUnit)
		err := row.Scan(&RackUnit.Id)
		return RackUnit, err
	}

	entities, err := RackUnit.Fetch(comp, scan)
	if nil != err {
		return ""
	}

	if len(entities) == 0 {
		return ""
	}

	return entities[0].Id
}

func (RackUnit RackUnitRepo) IsExistsDescription(Description string) bool {
	comp := qcomp {
		Tables: []string {"Rack_Unit"},
		Columns: []string {"ID"},
		Selection: "DESCRIPTION = ?",
		SelectionArgs: []string {Description},
	}

	row, err := database.Query(comp)
	if nil != err {
		return false
	}

	defer row.Close()

	return row.Next()
}

func (repo RackUnitRepo) IsExists(RackUnitId string) bool {
	comp := qcomp {
		Tables: []string {"Rack_Unit"},
		Columns: []string {"ID"},
		Selection: "ID = ?",
		SelectionArgs: []string {RackUnitId},
	}

	row, err := database.Query(comp)
	if nil != err {
		return false
	}

	defer row.Close()

	return row.Next()
}

func (repo RackUnitRepo) GenerateId() string {
	Id := database.GeneratePrimaryKey(true,
		true, true,
		false, "R", 6)

	for repo.IsExists(Id) {
		Id = database.GeneratePrimaryKey(true,
			true, true,
			false, "R", 6)
	}

	return Id
}

func (repo RackUnitRepo) Insert(RackUnits ...entity.RackUnit) error {
	values := make([][]string, len(RackUnits))
	for i := range values {
		values[i] = []string {RackUnits[i].Id, RackUnits[i].Description}
	}

	comp := icomp {
		Table: "Rack_Unit",
		Columns: []string {"ID", "DESCRIPTION"},
		Values: values,
	}

	return repo.SqliteRepo.Insert(comp)
}

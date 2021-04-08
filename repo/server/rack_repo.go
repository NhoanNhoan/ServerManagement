package server

import (
	"CURD/database"
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

func (Rack RackRepo) FetchAll() ([]entity.Rack, error) {
	comp := qcomp{
		Tables: []string {"Rack"},
		Columns: []string {"ID", "DESCRIPTION"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		center := obj.(entity.Rack)
		err := row.Scan(&center.Id, &center.Description)
		return center, err
	}

	return Rack.Fetch(comp, scan)
}

func (Rack RackRepo) FetchId(Description string) string {
	comp := qcomp {
		Tables: []string {"RACK"},
		Columns: []string {"ID"},
		Selection: "DESCRIPTION  =?",
		SelectionArgs: []string {Description},
	}

	scan := func (obj interface{}, row *sql.Rows) (interface{}, error) {
		rack := obj.(entity.Rack)
		err := row.Scan(&rack.Id)
		return rack, err
	}

	entities, err := Rack.Fetch(comp, scan)
	if nil != err {
		return ""
	}

	if len(entities) == 0 {
		return ""
	}

	return entities[0].Id
}

func (Rack RackRepo) IsExistsDescription(Description string) bool {
	comp := qcomp {
		Tables: []string {"RACK"},
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

func (repo RackRepo) IsExists(RackId string) bool {
	comp := qcomp {
		Tables: []string {"RACK"},
		Columns: []string {"ID"},
		Selection: "ID = ?",
		SelectionArgs: []string {RackId},
	}

	row, err := database.Query(comp)
	if nil != err {
		return false
	}

	defer row.Close()

	return row.Next()
}

func (repo RackRepo) GenerateId() string {
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

func (repo RackRepo) Insert(racks ...entity.Rack) error {
	values := make([][]string, len(racks))
	for i := range values {
		values[i] = []string {racks[i].Id, racks[i].Description}
	}

	comp := icomp {
		Table: "RACK",
		Columns: []string {"ID", "DESCRIPTION"},
		Values: values,
	}

	return repo.SqliteRepo.Insert(comp)
}
package server

import (
	"CURD/database"
	"CURD/entity"
	"database/sql"
)

type ServeCustomerRepo struct {
	SqliteRepo
}

func (r ServeCustomerRepo) Fetch(comp qcomp,
	scan func(obj interface{}, rows *sql.Rows) (interface{}, error)) ([]entity.ServeCustomer, error) {
	makeServe := func() interface{} { return entity.ServeCustomer{} }
	entities, err := r.SqliteRepo.Query(comp, makeServe, scan)
	if nil != err {
		return nil, err
	}

	Serves := make([]entity.ServeCustomer, len(entities))
	for i := range Serves {
		Serves[i] = entities[i].(entity.ServeCustomer)
	}

	return Serves, nil
}

func (Serve ServeCustomerRepo) FetchAll() ([]entity.ServeCustomer, error) {
	comp := qcomp{
		Tables: []string {"Serve"},
		Columns: []string {"ID", "DESCRIPTION"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		center := obj.(entity.ServeCustomer)
		err := row.Scan(&center.Id, &center.Description)
		return center, err
	}

	return Serve.Fetch(comp, scan)
}

func (Serve ServeCustomerRepo) FetchId(Description string) string {
	comp := qcomp {
		Tables: []string {"Serve"},
		Columns: []string {"ID"},
		Selection: "DESCRIPTION  =?",
		SelectionArgs: []string {Description},
	}

	scan := func (obj interface{}, row *sql.Rows) (interface{}, error) {
		Serve := obj.(entity.ServeCustomer)
		err := row.Scan(&Serve.Id)
		return Serve, err
	}

	entities, err := Serve.Fetch(comp, scan)
	if nil != err {
		return ""
	}

	if len(entities) == 0 {
		return ""
	}

	return entities[0].Id
}

func (Serve ServeCustomerRepo) IsExistsDescription(Description string) bool {
	comp := qcomp {
		Tables: []string {"Serve"},
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

func (repo ServeCustomerRepo) IsExists(ServeId string) bool {
	comp := qcomp {
		Tables: []string {"Serve"},
		Columns: []string {"ID"},
		Selection: "ID = ?",
		SelectionArgs: []string {ServeId},
	}

	row, err := database.Query(comp)
	if nil != err {
		return false
	}

	defer row.Close()

	return row.Next()
}

func (repo ServeCustomerRepo) GenerateId() string {
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

func (repo ServeCustomerRepo) Insert(Serves ...entity.ServeCustomer) error {
	values := make([][]string, len(Serves))
	for i := range values {
		values[i] = []string {Serves[i].Id, Serves[i].Description}
	}

	comp := icomp {
		Table: "Serve",
		Columns: []string {"ID", "DESCRIPTION"},
		Values: values,
	}

	return repo.SqliteRepo.Insert(comp)
}
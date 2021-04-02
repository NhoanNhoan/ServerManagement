package hardware

import (
	"CURD/database"
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
)

type ManagementRepo struct {
	server.SqliteRepo
}

func (repo ManagementRepo) Fetch(comp database.QueryComponent,
	scan func(obj interface{}, row *sql.Rows) (interface{}, error)) ([]hardware.Management, error) {
	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.Management{}},
		scan)

	if nil != err {
		return nil, err
	}

	listManagement := make([]hardware.Management, len(entities))
	for i := range listManagement {
		listManagement[i] = entities[i].(hardware.Management)
	}

	return listManagement, err
}

func (repo ManagementRepo) FetchAllManagements() ([]hardware.Management, error) {
	comp := qcomp {
		Tables: []string {"Management"},
		Columns: []string {"ID", "INFORMATION"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		Management := obj.(hardware.Management)
		err := row.Scan(&Management.Id, &Management.Information)
		return Management, err
	}

	return repo.Fetch(comp, scan)
}
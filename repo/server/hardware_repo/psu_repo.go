package hardware

import (
	"CURD/database"
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
)

type PSURepo struct {
	server.SqliteRepo
}

func (repo PSURepo) Fetch(comp database.QueryComponent,
	scan func(obj interface{}, row *sql.Rows) (interface{}, error)) ([]hardware.PSU, error) {
	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.PSU{}},
		scan)

	if nil != err {
		return nil, err
	}

	listPSU := make([]hardware.PSU, len(entities))
	for i := range listPSU {
		listPSU[i] = entities[i].(hardware.PSU)
	}

	return listPSU, err
}

func (repo PSURepo) FetchAllPSUs() ([]hardware.PSU, error) {
	comp := qcomp {
		Tables: []string {"PSU"},
		Columns: []string {"ID", "INFORMATION"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		PSU := obj.(hardware.PSU)
		err := row.Scan(&PSU.Id, &PSU.Information)
		return PSU, err
	}

	return repo.Fetch(comp, scan)
}
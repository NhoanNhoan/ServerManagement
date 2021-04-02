package hardware_repo

import (
	"CURD/database"
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
)

type RAMRepo struct {
	server.SqliteRepo
}

func (repo RAMRepo) Fetch(comp database.QueryComponent,
	scan func(obj interface{}, row *sql.Rows) (interface{}, error)) ([]hardware.RAM, error) {
	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.RAM{}},
		scan)

	if nil != err {
		return nil, err
	}

	listRAM := make([]hardware.RAM, len(entities))
	for i := range listRAM {
		listRAM[i] = entities[i].(hardware.RAM)
	}

	return listRAM, err
}

func (repo RAMRepo) FetchAllRAMs() ([]hardware.RAM, error) {
	comp := qcomp {
		Tables: []string {"RAM"},
		Columns: []string {"ID", "INFORMATION"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		RAM := obj.(hardware.RAM)
		err := row.Scan(&RAM.Id, &RAM.Information)
		return RAM, err
	}

	return repo.Fetch(comp, scan)
}
package hardware

import (
	"CURD/database"
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
)

type RaidRepo struct {
	server.SqliteRepo
}

func (repo RaidRepo) Fetch(comp database.QueryComponent,
	scan func(obj interface{}, row *sql.Rows) (interface{}, error)) ([]hardware.Raid, error) {
	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.Raid{}},
		scan)

	if nil != err {
		return nil, err
	}

	listRaid := make([]hardware.Raid, len(entities))
	for i := range listRaid {
		listRaid[i] = entities[i].(hardware.Raid)
	}

	return listRaid, err
}

func (repo RaidRepo) FetchAllRaids() ([]hardware.Raid, error) {
	comp := qcomp {
		Tables: []string {"Raid"},
		Columns: []string {"ID", "INFORMATION"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		Raid := obj.(hardware.Raid)
		err := row.Scan(&Raid.Id, &Raid.Information)
		return Raid, err
	}

	return repo.Fetch(comp, scan)
}
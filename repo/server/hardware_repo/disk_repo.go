package hardware

import (
	"CURD/database"
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
)

type DiskRepo struct {
	server.SqliteRepo
}

func (repo DiskRepo) Fetch(comp database.QueryComponent,
	scan func(obj interface{}, row *sql.Rows) (interface{}, error)) ([]hardware.Disk, error) {
	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.Disk{}},
		scan)

	if nil != err {
		return nil, err
	}

	listDisk := make([]hardware.Disk, len(entities))
	for i := range listDisk {
		listDisk[i] = entities[i].(hardware.Disk)
	}

	return listDisk, err
}

func (repo DiskRepo) FetchAllDisks() ([]hardware.Disk, error) {
	comp := qcomp {
		Tables: []string {"Disk"},
		Columns: []string {"ID", "INFORMATION"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		Disk := obj.(hardware.Disk)
		err := row.Scan(&Disk.Id, &Disk.Information)
		return Disk, err
	}

	return repo.Fetch(comp, scan)
}
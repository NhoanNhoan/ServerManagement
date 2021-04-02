package hardware_repo

import (
	"CURD/database"
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
)

type NICRepo struct {
	server.SqliteRepo
}

func (repo NICRepo) Fetch(comp database.QueryComponent,
	scan func(obj interface{}, row *sql.Rows) (interface{}, error)) ([]hardware.NIC, error) {
	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.NIC{}},
		scan)

	if nil != err {
		return nil, err
	}

	listNIC := make([]hardware.NIC, len(entities))
	for i := range listNIC {
		listNIC[i] = entities[i].(hardware.NIC)
	}

	return listNIC, err
}

func (repo NICRepo) FetchAllNICs() ([]hardware.NIC, error) {
	comp := qcomp {
		Tables: []string {"NIC"},
		Columns: []string {"ID", "INFORMATION"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		NIC := obj.(hardware.NIC)
		err := row.Scan(&NIC.Id, &NIC.Information)
		return NIC, err
	}

	return repo.Fetch(comp, scan)
}
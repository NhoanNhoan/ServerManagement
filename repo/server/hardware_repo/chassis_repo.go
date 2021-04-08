package hardware_repo

import (
	"CURD/database"
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
)

type ChassisRepo struct {
	server.SqliteRepo
}

func (repo ChassisRepo) FetchById(ChassisId string) (hardware.Chassis, error) {
	comp := qcomp {
		Tables: []string {"Chassis"},
		Columns: []string {"ID", "INFORMATION"},
		Selection: "ID = ?",
		SelectionArgs: []string {ChassisId},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		Chassis := obj.(hardware.Chassis)
		err := row.Scan(&Chassis.Id, &Chassis.Information)
		return Chassis, err
	}

	Chassiss, err := repo.Fetch(comp, scan)
	if nil != err {
		return hardware.Chassis{}, err
	}

	if len(Chassiss) > 0 {
		return Chassiss[0], err
	}

	return hardware.Chassis{}, err
}

func (repo ChassisRepo) Fetch(comp database.QueryComponent,
	scan func(obj interface{}, row *sql.Rows) (interface{}, error)) ([]hardware.Chassis, error) {
	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.Chassis{}},
		scan)

	if nil != err {
		return nil, err
	}

	listChassis := make([]hardware.Chassis, len(entities))
	for i := range listChassis {
		listChassis[i] = entities[i].(hardware.Chassis)
	}

	return listChassis, err
}

func (repo ChassisRepo) FetchAllChassiss() ([]hardware.Chassis, error) {
	comp := qcomp {
		Tables: []string {"Chassis"},
		Columns: []string {"ID", "INFORMATION"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		Chassis := obj.(hardware.Chassis)
		err := row.Scan(&Chassis.Id, &Chassis.Information)
		return Chassis, err
	}

	return repo.Fetch(comp, scan)
}
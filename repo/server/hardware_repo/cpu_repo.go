package hardware

import (
	"CURD/database"
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
)

type CPURepo struct {
	server.SqliteRepo
}

func (repo CPURepo) FetchById(CpuId string) (hardware.CPU, error) {
	comp := qcomp {
		Tables: []string {"CPU"},
		Columns: []string {"ID", "INFORMATION"},
		Selection: "ID = ?",
		SelectionArgs: []string {CpuId},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		cpu := obj.(hardware.CPU)
		err := row.Scan(&cpu.Id, &cpu.Information)
		return cpu, err
	}

	cpus, err := repo.Fetch(comp, scan)
	if nil != err {
		return hardware.CPU{}, err
	}

	if len(cpus) > 0 {
		return cpus[0], err
	}

	return hardware.CPU{}, err
}

func (repo CPURepo) Fetch(comp database.QueryComponent,
	scan func(obj interface{}, row *sql.Rows) (interface{}, error)) ([]hardware.CPU, error) {
	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.CPU{}},
		scan)

	if nil != err {
		return nil, err
	}

	listCpu := make([]hardware.CPU, len(entities))
	for i := range listCpu {
		listCpu[i] = entities[i].(hardware.CPU)
	}

	return listCpu, err
}

func (repo CPURepo) FetchAllCPUs() ([]hardware.CPU, error) {
	comp := qcomp {
		Tables: []string {"CPU"},
		Columns: []string {"ID", "INFORMATION"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		cpu := obj.(hardware.CPU)
		err := row.Scan(&cpu.Id, &cpu.Information)
		return cpu, err
	}

	return repo.Fetch(comp, scan)
}
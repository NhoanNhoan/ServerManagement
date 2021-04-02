package hardware_repo

import (
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
)

type HardwarePsuRepo struct {
	server.SqliteRepo
}

func (repo HardwarePsuRepo) FetchByHardwareId(HardwareId string) ([]hardware.HardwarePsu, error) {
	comp := qcomp {
		Tables: []string {"HARDWARE_CPU"},
		Columns: []string {"CPU_ID", "NUMBER_CPU"},
		Selection: "HARDWARE_ID = ?",
		SelectionArgs: []string {HardwareId},
	}

	scan := func (obj interface{}, row *sql.Rows) (interface{}, error) {
		hw := obj.(hardware.HardwarePsu)
		err := row.Scan(&hw.PsuId, &hw.NumberPsu)
		return hw, err
	}

	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.HardwarePsu{}},
		scan)

	if nil != err {
		return nil, err
	}

	HardwarePsus := make([]hardware.HardwarePsu, len(entities))
	for i := range HardwarePsus {
		HardwarePsus[i] = entities[i].(hardware.HardwarePsu)
	}

	return HardwarePsus, err
}
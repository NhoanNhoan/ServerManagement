package hardware_repo

import (
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
)

type HardwareRamRepo struct {
	server.SqliteRepo
}

func (repo HardwareRamRepo) FetchByHardwareId(HardwareId string) ([]hardware.HardwareRam, error) {
	comp := qcomp {
		Tables: []string {"HARDWARE_CPU"},
		Columns: []string {"CPU_ID", "NUMBER_CPU"},
		Selection: "HARDWARE_ID = ?",
		SelectionArgs: []string {HardwareId},
	}

	scan := func (obj interface{}, row *sql.Rows) (interface{}, error) {
		hw := obj.(hardware.HardwareRam)
		err := row.Scan(&hw.RamId, &hw.NumberRam)
		return hw, err
	}

	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.HardwareRam{}},
		scan)

	if nil != err {
		return nil, err
	}

	HardwareRams := make([]hardware.HardwareRam, len(entities))
	for i := range HardwareRams {
		HardwareRams[i] = entities[i].(hardware.HardwareRam)
	}

	return HardwareRams, err
}
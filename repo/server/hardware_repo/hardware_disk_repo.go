package hardware_repo

import (
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
)

type HardwareDiskRepo struct {
	server.SqliteRepo
}

func (repo HardwareDiskRepo) FetchByHardwareId(HardwareId string) ([]hardware.HardwareDisk, error) {
	comp := qcomp {
		Tables: []string {"HARDWARE_CPU"},
		Columns: []string {"CPU_ID", "NUMBER_CPU"},
		Selection: "HARDWARE_ID = ?",
		SelectionArgs: []string {HardwareId},
	}

	scan := func (obj interface{}, row *sql.Rows) (interface{}, error) {
		hw := obj.(hardware.HardwareDisk)
		err := row.Scan(&hw.DiskId, &hw.NumberDisk)
		return hw, err
	}

	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.HardwareDisk{}},
		scan)

	if nil != err {
		return nil, err
	}

	HardwareDisks := make([]hardware.HardwareDisk, len(entities))
	for i := range HardwareDisks {
		HardwareDisks[i] = entities[i].(hardware.HardwareDisk)
	}

	return HardwareDisks, err
}
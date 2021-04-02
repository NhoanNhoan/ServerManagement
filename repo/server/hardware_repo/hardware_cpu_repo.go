package hardware

import (
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
)

type HardwareCPURepo struct {
	server.SqliteRepo
}

func (repo HardwareCPURepo) FetchByHardwareId(HardwareId string) ([]hardware.HardwareCpu, error) {
	comp := qcomp {
		Tables: []string {"HARDWARE_CPU"},
		Columns: []string {"CPU_ID", "NUMBER_CPU"},
		Selection: "HARDWARE_ID = ?",
		SelectionArgs: []string {HardwareId},
	}

	scan := func (obj interface{}, row *sql.Rows) (interface{}, error) {
		hw := obj.(hardware.HardwareCpu)
		err := row.Scan(&hw.CpuId, &hw.NumberCpu)
		return hw, err
	}

	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.HardwareCpu{}},
		scan)

	if nil != err {
		return nil, err
	}

	hardwareCpus := make([]hardware.HardwareCpu, len(entities))
	for i := range hardwareCpus {
		hardwareCpus[i] = entities[i].(hardware.HardwareCpu)
	}

	return hardwareCpus, err
}
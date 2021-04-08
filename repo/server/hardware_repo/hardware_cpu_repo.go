package hardware_repo

import (
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
	"strconv"
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

func (repo HardwareCPURepo) Insert(HardwareId string, HardwareCpuArr ...hardware.HardwareCpu) error {
	comp := icomp {
		Table: "HARDWARE_CPU",
		Columns: []string {"HARDWARE_ID", "CPU_ID", "NUMBER_CPU"},
		Values: repo.makeInsertValues(HardwareCpuArr...),
	}

	return repo.SqliteRepo.Insert(comp)
}

func (repo HardwareCPURepo) makeInsertValues(HardwareCpuArr ...hardware.HardwareCpu) [][]string {
	values := make([][]string, len(HardwareCpuArr))
	for i := range values {
		values[i] = []string {HardwareCpuArr[i].HardwareId,
			HardwareCpuArr[i].CpuId,
			strconv.Itoa(HardwareCpuArr[i].NumberCpu)}
	}
	return values
}

func (repo HardwareCPURepo) Delete(HardwareConfigId string) error {
	comp := dcomp{
		Table:         "HARDWARE_CPU",
		Selection:     "HARDWARE_ID = ?",
		SelectionArgs: []string{HardwareConfigId},
	}

	return repo.SqliteRepo.Delete(comp)
}
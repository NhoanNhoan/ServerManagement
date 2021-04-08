package hardware_repo

import (
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
	"strconv"
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

func (repo HardwareRamRepo) Insert(HardwareId string, HardwareRamArr ...hardware.HardwareRam) error {
	comp := icomp {
		Table: "HARDWARE_Ram",
		Columns: []string {"HARDWARE_ID", "Ram_ID", "NUMBER_Ram"},
		Values: repo.makeInsertValues(HardwareRamArr...),
	}

	return repo.SqliteRepo.Insert(comp)
}

func (repo HardwareRamRepo) makeInsertValues(HardwareRamArr ...hardware.HardwareRam) [][]string {
	values := make([][]string, len(HardwareRamArr))
	for i := range values {
		values[i] = []string {HardwareRamArr[i].HardwareId,
			HardwareRamArr[i].RamId,
			strconv.Itoa(HardwareRamArr[i].NumberRam)}
	}
	return values
}

func (repo HardwareRamRepo) Delete(HardwareConfigId string) error {
	comp := dcomp{
		Table:         "HARDWARE_Ram",
		Selection:     "HARDWARE_ID = ?",
		SelectionArgs: []string{HardwareConfigId},
	}

	return repo.SqliteRepo.Delete(comp)
}
package hardware_repo

import (
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
	"strconv"
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

func (repo HardwarePsuRepo) Insert(HardwareId string, HardwarePsuArr ...hardware.HardwarePsu) error {
	comp := icomp {
		Table: "HARDWARE_Psu",
		Columns: []string {"HARDWARE_ID", "Psu_ID", "NUMBER_Psu"},
		Values: repo.makeInsertValues(HardwarePsuArr...),
	}

	return repo.SqliteRepo.Insert(comp)
}

func (repo HardwarePsuRepo) makeInsertValues(HardwarePsuArr ...hardware.HardwarePsu) [][]string {
	values := make([][]string, len(HardwarePsuArr))
	for i := range values {
		values[i] = []string {HardwarePsuArr[i].HardwareId,
			HardwarePsuArr[i].PsuId,
			strconv.Itoa(HardwarePsuArr[i].NumberPsu)}
	}
	return values
}

func (repo HardwarePsuRepo) Delete(HardwareConfigId string) error {
	comp := dcomp{
		Table:         "HARDWARE_Psu",
		Selection:     "HARDWARE_ID = ?",
		SelectionArgs: []string{HardwareConfigId},
	}

	return repo.SqliteRepo.Delete(comp)
}
package hardware_repo

import (
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
	"strconv"
)

type HardwareNicRepo struct {
	server.SqliteRepo
}

func (repo HardwareNicRepo) FetchByHardwareId(HardwareId string) ([]hardware.HardwareNic, error) {
	comp := qcomp {
		Tables: []string {"HARDWARE_CPU"},
		Columns: []string {"CPU_ID", "NUMBER_CPU"},
		Selection: "HARDWARE_ID = ?",
		SelectionArgs: []string {HardwareId},
	}

	scan := func (obj interface{}, row *sql.Rows) (interface{}, error) {
		hw := obj.(hardware.HardwareNic)
		err := row.Scan(&hw.NicId, &hw.NumberNic)
		return hw, err
	}

	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.HardwareNic{}},
		scan)

	if nil != err {
		return nil, err
	}

	HardwareNics := make([]hardware.HardwareNic, len(entities))
	for i := range HardwareNics {
		HardwareNics[i] = entities[i].(hardware.HardwareNic)
	}

	return HardwareNics, err
}

func (repo HardwareNicRepo) Insert(HardwareId string, HardwareNicArr ...hardware.HardwareNic) error {
	comp := icomp {
		Table: "HARDWARE_Nic",
		Columns: []string {"HARDWARE_ID", "Nic_ID", "NUMBER_Nic"},
		Values: repo.makeInsertValues(HardwareNicArr...),
	}

	return repo.SqliteRepo.Insert(comp)
}

func (repo HardwareNicRepo) makeInsertValues(HardwareNicArr ...hardware.HardwareNic) [][]string {
	values := make([][]string, len(HardwareNicArr))
	for i := range values {
		values[i] = []string {HardwareNicArr[i].HardwareId,
			HardwareNicArr[i].NicId,
			strconv.Itoa(HardwareNicArr[i].NumberNic)}
	}
	return values
}

func (repo HardwareNicRepo) Delete(HardwareConfigId string) error {
	comp := dcomp{
		Table:         "HARDWARE_Nic",
		Selection:     "HARDWARE_ID = ?",
		SelectionArgs: []string{HardwareConfigId},
	}

	return repo.SqliteRepo.Delete(comp)
}
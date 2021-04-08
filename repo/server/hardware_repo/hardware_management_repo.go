package hardware_repo

import (
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
	"strconv"
)

type HardwareManagementRepo struct {
	server.SqliteRepo
}

func (repo HardwareManagementRepo) FetchByHardwareId(HardwareId string) ([]hardware.HardwareManagement, error) {
	comp := qcomp {
		Tables: []string {"HARDWARE_CPU"},
		Columns: []string {"CPU_ID", "NUMBER_CPU"},
		Selection: "HARDWARE_ID = ?",
		SelectionArgs: []string {HardwareId},
	}

	scan := func (obj interface{}, row *sql.Rows) (interface{}, error) {
		hw := obj.(hardware.HardwareManagement)
		err := row.Scan(&hw.ManagementId, &hw.NumberManagement)
		return hw, err
	}

	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.HardwareManagement{}},
		scan)

	if nil != err {
		return nil, err
	}

	HardwareManagements := make([]hardware.HardwareManagement, len(entities))
	for i := range HardwareManagements {
		HardwareManagements[i] = entities[i].(hardware.HardwareManagement)
	}

	return HardwareManagements, err
}

func (repo HardwareManagementRepo) Insert(HardwareId string, HardwareManagementArr ...hardware.HardwareManagement) error {
	comp := icomp {
		Table: "HARDWARE_MNT",
		Columns: []string {"HARDWARE_ID", "MNT_ID", "NUMBER_MNT"},
		Values: repo.makeInsertValues(HardwareManagementArr...),
	}

	return repo.SqliteRepo.Insert(comp)
}

func (repo HardwareManagementRepo) makeInsertValues(HardwareManagementArr ...hardware.HardwareManagement) [][]string {
	values := make([][]string, len(HardwareManagementArr))
	for i := range values {
		values[i] = []string {HardwareManagementArr[i].HardwareId,
			HardwareManagementArr[i].ManagementId,
			strconv.Itoa(HardwareManagementArr[i].NumberManagement)}
	}
	return values
}

func (repo HardwareManagementRepo) Delete(HardwareConfigId string) error {
	comp := dcomp{
		Table:         "HARDWARE_MNT",
		Selection:     "HARDWARE_ID = ?",
		SelectionArgs: []string{HardwareConfigId},
	}

	return repo.SqliteRepo.Delete(comp)
}
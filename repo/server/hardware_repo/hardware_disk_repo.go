package hardware_repo

import (
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
	"strconv"
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

func (repo HardwareDiskRepo) Insert(HardwareId string, HardwareDiskArr ...hardware.HardwareDisk) error {
	comp := icomp {
		Table: "HARDWARE_Disk",
		Columns: []string {"HARDWARE_ID", "Disk_ID", "NUMBER_Disk"},
		Values: repo.makeInsertValues(HardwareDiskArr...),
	}

	return repo.SqliteRepo.Insert(comp)
}

func (repo HardwareDiskRepo) makeInsertValues(HardwareDiskArr ...hardware.HardwareDisk) [][]string {
	values := make([][]string, len(HardwareDiskArr))
	for i := range values {
		values[i] = []string {HardwareDiskArr[i].HardwareId,
			HardwareDiskArr[i].DiskId,
			strconv.Itoa(HardwareDiskArr[i].NumberDisk)}
	}
	return values
}

func (repo HardwareDiskRepo) Delete(HardwareConfigId string) error {
	comp := dcomp{
		Table:         "HARDWARE_Disk",
		Selection:     "HARDWARE_ID = ?",
		SelectionArgs: []string{HardwareConfigId},
	}

	return repo.SqliteRepo.Delete(comp)
}
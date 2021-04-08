package hardware_repo

import (
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
	"strconv"
)

type HardwareRaidRepo struct {
	server.SqliteRepo
}

func (repo HardwareRaidRepo) FetchByHardwareId(HardwareId string) ([]hardware.HardwareRaid, error) {
	comp := qcomp {
		Tables: []string {"HARDWARE_CPU"},
		Columns: []string {"CPU_ID", "NUMBER_CPU"},
		Selection: "HARDWARE_ID = ?",
		SelectionArgs: []string {HardwareId},
	}

	scan := func (obj interface{}, row *sql.Rows) (interface{}, error) {
		hw := obj.(hardware.HardwareRaid)
		err := row.Scan(&hw.RaidId, &hw.NumberRaid)
		return hw, err
	}

	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.HardwareRaid{}},
		scan)

	if nil != err {
		return nil, err
	}

	HardwareRaids := make([]hardware.HardwareRaid, len(entities))
	for i := range HardwareRaids {
		HardwareRaids[i] = entities[i].(hardware.HardwareRaid)
	}

	return HardwareRaids, err
}

func (repo HardwareRaidRepo) Insert(HardwareId string, HardwareRaidArr ...hardware.HardwareRaid) error {
	comp := icomp {
		Table: "HARDWARE_Raid",
		Columns: []string {"HARDWARE_ID", "Raid_ID", "NUMBER_Raid"},
		Values: repo.makeInsertValues(HardwareRaidArr...),
	}

	return repo.SqliteRepo.Insert(comp)
}

func (repo HardwareRaidRepo) makeInsertValues(HardwareRaidArr ...hardware.HardwareRaid) [][]string {
	values := make([][]string, len(HardwareRaidArr))
	for i := range values {
		values[i] = []string {HardwareRaidArr[i].HardwareId,
			HardwareRaidArr[i].RaidId,
			strconv.Itoa(HardwareRaidArr[i].NumberRaid)}
	}
	return values
}

func (repo HardwareRaidRepo) Delete(HardwareConfigId string) error {
	comp := dcomp{
		Table:         "HARDWARE_Raid",
		Selection:     "HARDWARE_ID = ?",
		SelectionArgs: []string{HardwareConfigId},
	}

	return repo.SqliteRepo.Delete(comp)
}
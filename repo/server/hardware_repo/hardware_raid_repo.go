package hardware_repo

import (
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
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
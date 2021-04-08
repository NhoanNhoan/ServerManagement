package hardware_repo

import (
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
)

type HardwareTemplateRepo struct {
	server.SqliteRepo
}

func (repo HardwareTemplateRepo) FetchAll() ([]hardware.HardwareTemplate, error) {
	comp := qcomp {
		Tables: []string {"HARDWARE_TEMPLATE"},
		Columns: []string {"CPU_ID", "RAM_ID", "DISK_ID", "RAID_ID", "NIC_ID", "PSU_ID", "MANAGEMENT_ID"},
	}

	scan := func (obj interface{}, row *sql.Rows) (interface{}, error) {
		hw := obj.(hardware.HardwareTemplate)
		err := row.Scan(&hw.CpuId, &hw.RamId, &hw.DiskId, &hw.RaidId, &hw.NicId, &hw.PsuId, &hw.ManagementId)
		return hw, err
	}

	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.HardwareTemplate{}},
		scan)

	if nil != err {
		return nil, err
	}

	templates := make([]hardware.HardwareTemplate, len(entities))
	for i := range templates {
		templates[i] = entities[i].(hardware.HardwareTemplate)
	}

	return templates, err
}
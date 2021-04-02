package hardware

import (
	"CURD/database"
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
)

type qcomp = database.QueryComponent
type icomp = database.InsertComponent
type ucomp = database.UpdateComponent
type dcomp = database.DeleteComponent

type HardwareConfigRepo struct {
	server.SqliteRepo
}

func (repo HardwareConfigRepo) Fetch(comp database.QueryComponent,
	scan func(obj interface{}, row *sql.Rows) (interface{}, error)) ([]hardware.HardwareConfig, error) {
	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.HardwareConfig{}},
		scan)

	if nil != err {
		return nil, err
	}

	listHardwareConfig := make([]hardware.HardwareConfig, len(entities))
	for i := range listHardwareConfig {
		listHardwareConfig[i] = entities[i].(hardware.HardwareConfig)
	}

	return listHardwareConfig, err
}

func (repo HardwareConfigRepo) FetchAllHardwareConfigs() ([]hardware.HardwareConfig, error) {
	comp := qcomp {
		Tables: []string {"HardwareConfig"},
		Columns: []string {"ID", "CHASSIS_ID", "CLUSTER_ID"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		HardwareConfig := obj.(hardware.HardwareConfig)
		err := row.Scan(&HardwareConfig.Id, &HardwareConfig.ChassisId, &HardwareConfig.ClusterId)
		return HardwareConfig, err
	}

	return repo.Fetch(comp, scan)
}
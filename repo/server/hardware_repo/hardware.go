package hardware_repo

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

func (repo HardwareConfigRepo) FetchByServerId(ServerId string) (*hardware.HardwareConfig, error) {
	comp := qcomp {
		Tables: []string {"SERVER", "HARDWARE_CONFIG"},
		Columns: []string {"HARDWARE_CONFIG.ID", "CHASSIS_ID", "CLUSTER_SERVER_ID"},
		Selection: "SERVER.ID = ? AND SERVER.HARDWARE_CONFIG_ID = HARDWARE_CONFIG.ID",
		SelectionArgs: []string {ServerId},
	}

	scan := func (obj interface{}, row *sql.Rows) (interface{}, error) {
		hw := obj.(hardware.HardwareConfig)
		err := row.Scan(&hw.Id, &hw.ChassisId, &hw.ClusterId)
		return hw, err
	}

	list_hw, err := repo.Fetch(comp, scan)
	if nil != err {
		return nil, err
	}

	if len(list_hw) == 0 {
		return nil, nil
	}

	return &list_hw[0], nil
}

func (repo HardwareConfigRepo) Insert(listHw ...hardware.HardwareConfig) error {
	values := make([][]string, len(listHw))
	for i := range listHw {
		values[i] = []string {listHw[i].Id, listHw[i].ChassisId, listHw[i].ClusterId}
	}

	comp := icomp {
		Table: "HARDWARE_CONFIG",
		Columns: []string {"ID", "CHASSIS_ID", "CLUSTER_SERVER_ID"},
		Values: values,
	}

	return repo.SqliteRepo.Insert(comp)
}

func (repo HardwareConfigRepo) GenerateId() string {
	Id := database.GeneratePrimaryKey(true,
		true, true,
		false, "HW", 6)

	for repo.IsExists(Id) {
		Id = database.GeneratePrimaryKey(true,
			true, true,
			false, "HW", 6)
	}

	return Id
}

func (repo HardwareConfigRepo) IsExists(Id string) bool {
	comp := qcomp {
		Tables: []string {"HARDWARE_CONFIG"},
		Columns: []string {"ID"},
		Selection: "ID = ?",
		SelectionArgs: []string {Id},
	}

	row, err := database.Query(comp)
	if nil != err {
		return false
	}

	defer row.Close()

	return row.Next()
}

func (repo HardwareConfigRepo) FetchConfigId(ServerId string) (string, error) {
	comp := qcomp {
		Tables: []string {"SERVER AS S", "HARDWARE_CONFIG AS H"},
		Columns: []string {"H.ID"},
		Selection: "S.HARDWARE_CONFIG_ID = ?",
		SelectionArgs: []string {ServerId},
	}

	scan := func (obj interface{}, row *sql.Rows) (interface{}, error) {
		hw := obj.(hardware.HardwareConfig)
		err := row.Scan(&hw.Id)
		return hw, err
	}

	hws, err := repo.Fetch(comp, scan)
	if nil != err || len(hws) == 0 {return "", err}
	return hws[0].Id, err
}

func (repo HardwareConfigRepo) FetchClusterServer(ServerId string) (string, error) {
	comp := qcomp {
		Tables: []string {"SERVER AS S", "HARDWARE_CONFIG AS H", "CLUSTER_SERVER AS C"},
		Columns: []string {"C.NAME"},
		Selection: "S.ID = ? AND S.HARDWARE_CONFIG_ID = H.ID AND H.CLUSTER_SERVER_ID = C.ID",
		SelectionArgs: []string {ServerId},
	}

	row, err := database.Query(comp)
	if nil != err {return "", err}
	defer row.Close()
	clusterName := ""
	if row.Next(){err = row.Scan(&clusterName)}
	return clusterName, err
}

func (repo HardwareConfigRepo) Delete(HwId string) error {
	comp := dcomp {
		Table: "HARDWARE_CONFIG",
		Selection: "ID = ?",
		SelectionArgs: []string {HwId},
	}

	return repo.SqliteRepo.Delete(comp)
}

func (repo HardwareConfigRepo) Update(HwConfig hardware.HardwareConfig) error {
	comp := ucomp {
		Table: "HARDWARE_CONFIG",
		SetClause: "CHASSIS_ID = ?, CLUSTER_SERVER_ID = ?",
		Values: []string {HwConfig.ChassisId, HwConfig.ClusterId},
		Selection: "ID = ?",
		SelectionArgs: []string {HwConfig.Id},
	}

	return repo.SqliteRepo.Update(comp)
}
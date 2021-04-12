package hardware_repo

import "CURD/repo/server"

type HardwareComponentRepo struct {
	TableName string
	Columns []string
	server.SqliteRepo
}

type hr = HardwareComponentRepo

func (h hr) Delete(configId string) error {
	service := hs{h.TableName, h.Columns}
	comp := service.MakeDeleteByHardwareConfig(configId)
	return h.SqliteRepo.Delete(comp)
}

type HardwareComponentRepoService struct {
	TableName string
	Columns []string
}

type hs = HardwareComponentRepoService

func (service hs) MakeDeleteByHardwareConfig(configId string) dcomp {
	return dcomp {
		Table: service.TableName,
		Selection: service.Columns[0] + " = ?",
		SelectionArgs: []string {configId},
	}
}

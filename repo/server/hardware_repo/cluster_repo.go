package hardware_repo

import (
	"CURD/database"
	"CURD/entity/server/hardware"
	"CURD/repo/server"
	"database/sql"
)

type ClusterRepo struct {
	server.SqliteRepo
}

func (repo ClusterRepo) Fetch(comp database.QueryComponent,
	scan func(obj interface{}, row *sql.Rows) (interface{}, error)) ([]hardware.ClusterServer, error) {
	entities, err := repo.SqliteRepo.Query(comp,
		func() interface{} {return hardware.ClusterServer{}},
		scan)

	if nil != err {
		return nil, err
	}

	listClusterServer := make([]hardware.ClusterServer, len(entities))
	for i := range listClusterServer {
		listClusterServer[i] = entities[i].(hardware.ClusterServer)
	}

	return listClusterServer, err
}

func (repo ClusterRepo) FetchAllClusterServers() ([]hardware.ClusterServer, error) {
	comp := qcomp {
		Tables: []string {"CLUSTER_SERVER"},
		Columns: []string {"ID", "NAME"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		ClusterServer := obj.(hardware.ClusterServer)
		err := row.Scan(&ClusterServer.Id, &ClusterServer.Name)
		return ClusterServer, err
	}

	return repo.Fetch(comp, scan)
}
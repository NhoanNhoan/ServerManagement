package server

import (
	"CURD/database"
	"CURD/entity"
	"database/sql"
)

type SwitchConnectionRepo struct {
	SqliteRepo
}

type sc_repo = SwitchConnectionRepo

func (repo sc_repo) IsExists(Id string) bool {
	row, err := database.Query(qcomp{
		Tables: []string {"SWITCH_CONNECTION"},
		Columns: []string {"ID"},
		Selection: "ID = ?",
		SelectionArgs: []string {Id},
	})

	if nil != err {
		return false
	}

	defer row.Close()
	return row.Next()
}

func (repo sc_repo) generateId() string {
	id := database.GeneratePrimaryKey(
		true, true,
		true, false,
		"SW", 6)

	for repo.IsExists(id) {
		id = database.GeneratePrimaryKey(
			true, true,
			true, false,
			"SW", 6)
	}

	return id
}

func (repo SwitchConnectionRepo) FetchByServerId(ServerId string) ([]entity.SwitchConnection, error) {
	comp := qcomp{
		Tables: []string {"SWITCH_CONNECTION"},
		Columns: []string {"ID", "ID_SERVER", "ID_SWITCH", "ID_CABLE_TYPE", "PORT"},
		Selection: "ID_SERVER = ?",
		SelectionArgs: []string {ServerId},
	}

	makeSwitchConn := func() interface{} {return entity.SwitchConnection{}}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		sw := obj.(entity.SwitchConnection)
		err := row.Scan(&sw.Id, &sw.ServerId, &sw.SwitchId, &sw.CableTypeId, &sw.CableTypeId)
		return sw, err
	}

	entities, err := repo.Query(comp, makeSwitchConn, scan)
	if nil != err {
		return nil, err
	}

	connArr := make([]entity.SwitchConnection, len(entities))

	for i := range connArr {
		connArr[i] = entities[i].(entity.SwitchConnection)
	}

	return connArr, nil
}

func (repo SwitchConnectionRepo) Insert(connections ...entity.SwitchConnection) error {
	comp := icomp{
		Table: "SWITCH_CONNECTION",
		Columns: []string {"ID_SERVER", "ID_SWITCH", "ID_CABLE_TYPE", "PORT"},
		Values: repo.makeValues(connections...),
	}

	return repo.SqliteRepo.Insert(comp)
}

func (repo SwitchConnectionRepo) makeValues(connections ...entity.SwitchConnection) [][]string {
	values := make([][]string, len(connections))
	for i := range connections {
		values[i] = []string {repo.generateId(),
			connections[i].ServerId,
			connections[i].SwitchId,
			connections[i].CableTypeId,
			connections[i].Port,
		}
	}
	return values
}
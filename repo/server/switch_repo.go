package server

import (
	"CURD/entity"
	"database/sql"
)

type SwitchRepo struct {
	SqliteRepo
}

func (repo SwitchRepo) Fetch(comp qcomp,
	scan func(obj interface{}, row *sql.Rows) (interface{}, error)) ([]entity.Switch, error) {
	makeSwitch := func() interface{} { return entity.Switch{} }
	entities, err := repo.SqliteRepo.Query(comp, makeSwitch, scan)
	if nil != err {
		return nil, err
	}

	switches := make([]entity.Switch, len(entities))
	for i := range switches {
		switches[i] = entities[i].(entity.Switch)
	}

	return switches, nil
}

func (repo SwitchRepo) Insert(comp icomp) error {
	return repo.SqliteRepo.Insert(comp)
}

func (repo SwitchRepo) FetchById(SwitchId string) entity.Switch {
	comp := qcomp{
		Tables: []string {"SWITCH", "DC", "RACK", "RACK_UNIT AS USTART", "RACK_UNIT AS UEND"},
		Columns: []string {"DC.DESCRIPTION", "RACK.DESCRIPTION", "USTART.DESCRIPTION", "UEND.DESCRIPTION", "MAXIMUM_PORT"},
		Selection: "SWITCH.ID = ? AND SWITCH.ID_DC = DC.ID AND SWITCH.ID_RACK = RACK.ID AND SWITCH.ID_U_START = USTART.ID AND SWITCH.ID_U_END = UEND.ID",
		SelectionArgs: []string {SwitchId},
	}

	makeSwitch := func() interface{} {return entity.Switch{}}

	scan := func (obj interface{}, row *sql.Rows) (interface{}, error) {
		sw := obj.(entity.Switch)
		err := row.Scan(&sw.Id, &sw.DC.Description, &sw.Rack.Description, &sw.UStart.Description, &sw.UEnd.Description, &sw.MaximumPort)
		return sw, err
	}

	entities, err := repo.SqliteRepo.Query(comp, makeSwitch, scan)
	if nil != err && len(entities) > 0 {
		return entities[0].(entity.Switch)
	}

	return entity.Switch{}
}
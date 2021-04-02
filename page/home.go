package page

import (
	"CURD/database"
	"CURD/entity"
	"CURD/repo/server"
)

type Home struct {
	DCs []entity.DataCenter
	Tags []entity.Tag
	AllIpNet []entity.IpNet
	NumUnresolvedErrors int
}

func (h *Home) New() error {
	var err error
	h.DCs, err = server.DCRepo{}.FetchAll()
	if nil != err {
		return err
	}

	h.Tags, err = server.TagRepo{}.FetchAll()
	if nil != err {
		return err
	}

	h.initNumUnresolvedErrors()
	//h.AllIpNet = entity.GetIpNets()

	return nil
}

func (h *Home) initNumUnresolvedErrors() {
	comp := h.makeCountNumUnresolvedErrors()
	rows, err := database.Query(comp)
	defer rows.Close()

	if nil == err && rows.Next() {
		rows.Scan(&h.NumUnresolvedErrors)
	}
}

func (h *Home) makeCountNumUnresolvedErrors() database.QueryComponent {
	return h.makeCountErrorsByStateQueryComponent("UNRESOLVED")
}

func (h *Home) makeCountErrorsByStateQueryComponent(state string) database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"ERROR", "ERROR_STATE"},
		Columns: []string {"count(ERROR.ID)"},
		Selection: "ERROR_STATE.DESCRIPTION = ? AND ERROR.ID_ERROR_STATE = ERROR_STATE.ID",
		SelectionArgs: []string {state},
	}
}
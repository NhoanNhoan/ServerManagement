package page

import (
	"CURD/database"
	"database/sql"
	"CURD/entity"
	"CURD/model"
)

type Home struct {
	DCs []entity.DataCenter
	Tags []entity.Tag
	AllIpNet []entity.IpNet
	NumUnresolvedErrors int
}

func (h *Home) New() {
	component := makeDCQueryComponent()
	rows, err := database.Query(component)
	defer rows.Close()

	if nil != err {
		panic (err)
	}

	err = h.fetch(rows)

	if nil != err {
		panic (err)
	}

	h.Tags = entity.FetchAllTags()
	h.initNumUnresolvedErrors()
	h.AllIpNet = entity.GetIpNets()
}

func makeDCQueryComponent() database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"DC"},
		Columns: []string {model.DC_ID, 
						model.DC_DESCRIPTION,
					},
		Selection: "",
		SelectionArgs: nil,
		GroupBy: "",
		Having: "",
		OrderBy: "",
		Limit: "",
	}
}

func (h *Home)fetch(rows *sql.Rows) (err error) {
	var id, name string

	for rows.Next() && rows.Scan(&id,&name) == nil {
		h.DCs = append(h.DCs, entity.DataCenter {id, name})
	}

	return
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
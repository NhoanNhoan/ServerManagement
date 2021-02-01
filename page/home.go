package page

import (
	"CURD/database"
	"database/sql"
	"CURD/entity"
	"CURD/model"
)

type Home struct {
	DCs []entity.DataCenter
}

func (h *Home) New() {
	component := makeDCQueryComponent()
	rows, err := database.Query(component)

	if nil != err {
		panic (err)
	}

	err = h.fetch(rows)

	if nil != err {
		panic (err)
	}
}

func makeDCQueryComponent() database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {model.DC_TB_NAME},
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
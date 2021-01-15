package page

import (
	"CURD/database"
	"CURD/model"
)

type Home struct {
	DCs []model.DataCenter
}

func (h *Home) Init() {
	sql := "SELECT ID, NAME FROM DC"
	db := database.DBConn()
	rows, err := db.Query(sql)

	if nil != err {
		panic (err)
	}

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)

		if nil != err {
			panic(err)
		}

		h.DCs = append(h.DCs, model.DataCenter {id, name})
	}
}
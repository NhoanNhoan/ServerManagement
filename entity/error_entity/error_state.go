package error_entity

import (
	"CURD/database"

	"database/sql"
)

type ErrorState struct {
	Id string
	Description string
}

func (obj *ErrorState) New(IdErrorState string) error {
	obj.Id = IdErrorState
	row := obj.query()
	return obj.Fetch(row)
}

func (obj *ErrorState) query() *sql.Rows {
	comp := obj.queryComp()
	row, err := database.Query(comp)
	row.Close()

	if nil != err {
		panic (err)
	}

	return row
}

func (obj *ErrorState) queryComp() database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"ERROR_STATE"},
		Columns: []string {"ID", "DESCRIPTION"},
		Selection: "ID = ?",
		SelectionArgs: []string {obj.Id},
		GroupBy: "",
		Having: "",
		OrderBy: "",
		Limit: "",
	}
}

func (obj *ErrorState) Fetch(row *sql.Rows) error {
	var err error

	if row.Next() {
		err = row.Scan(&obj.Id, &obj.Description)
	}

	return err
}

func FetchErrorStates() []ErrorState {
	states := make([]ErrorState, 0)
	rows := queryErrorStates()
	var err error
	var state ErrorState

	for rows.Next() {
		err = rows.Scan(&state.Id, 
						&state.Description)
		if nil != err {
			panic (err)
		}
		
		states = append (states, state)
	}

	return states
}

func queryErrorStates() *sql.Rows {
	comp := errorStatesComp()
	rows, err := database.Query(comp)
	defer rows.Close()

	if nil != err {
		panic (err)
	}

	return rows
}

func errorStatesComp() database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"ERROR_STATE"},
		Columns: []string {"ID", "DESCRIPTION"},
		Selection: "",
		SelectionArgs: nil,
		GroupBy: "",
		Having: "",
		OrderBy: "",
		Limit: "",
	}
}

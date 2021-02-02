package error_entity

import (
	"CURD/database"
	"CURD/entity"
	"database/sql"
	"errors"
	"strings"
)

type qcomp = database.QueryComponent
type ucomp = database.UpdateComponent
type icomp = database.InsertComponent

type Error struct {
	Id string
	Summary string
	Description string
	Solution string
	Occurs string
	entity.Server
	ErrorState
}

func (obj *Error) New(IdError string) (err error) {
	obj.Id = IdError
	comp := obj.makeQueryComponent(IdError)
	row, err := database.Query(comp)
	row.Next()
	return obj.Fetch(row)
}

func (obj *Error) makeQueryComponent(IdError string) qcomp {
	return database.QueryComponent {
		Tables: []string {
				"ERROR AS E",
				"ERROR_STATE AS ES",
				"STATUS_ROW AS ST",
			},
			
		Columns: []string {
				"E.ID",
				"E.SUMMARY",
				"E.DESCRIPTION",
				"E.SOLUTION",
				"E.OCCURS",
				"E.ID_SERVER",
				"ES.ID",
				"ES.DESCRIPTION",
			},
			
		Selection: "E.ID = ? AND " + 
				"ST.DESCRIPTION = ? AND " + 
				"E.ID_STATUS_ROW = ST.ID",
		
		SelectionArgs: []string {
				IdError,
				"available",
			},
			
		GroupBy: "",
		Having: "",
		OrderBy: "",
		Limit: "",
	}
}

func (obj *Error) Fetch(row *sql.Rows) (err error) {
		err = row.Scan(
			&obj.Id,
			&obj.Summary,
			&obj.Description,
			&obj.Solution,
			&obj.Occurs,
			&obj.Server.Id,
			&obj.ErrorState.Id,
			&obj.ErrorState.Description,
		)
	
	return
}

func (obj *Error) FetchServer() (err error) {
	IdServer := obj.Server.Id
	return obj.Server.New(IdServer)
}

func (obj *Error) Update() error {
	if !obj.isExists() {
		msg := "This error instance don't isExists"
		return errors.New(msg)
	}

	comp := obj.updateComp()
	return database.Update(comp)
}

func (obj *Error) isExists() bool {
	comp := obj.isExistsComp()
	rows, err := database.Query(comp)
	return (nil == err) && rows.Next()
}

func (obj *Error) isExistsComp() qcomp {
	tables := []string {"ERROR"}
	columns := []string {"ID"}
	selection := "ID = ?"
	selectionArgs := []string {obj.Id}

	return qcomp {
		Tables: tables,
		Columns: columns,
		Selection: selection,
		SelectionArgs: selectionArgs,
	}
}

func (obj *Error) updateComp() ucomp {
	tables := "ERROR"
	setClause, values := obj.makeSetClause()
	selection := "ID = ?"
	args := []string {obj.Id}

	return ucomp {
			Table: tables, 
			SetClause: setClause, 
			Values: values, 
			Selection: selection, 
			SelectionArgs: args,
		}
}

func (obj *Error) makeSetClause() (string, []string) {
	const EMPTY string = ""

	setFields := make([]string, 0)
	args := make([]string, 0)

	if EMPTY != obj.Summary {
		setFields = append(setFields, "SUMMARY = ?")
		args = append(args, obj.Summary)
	}

	if EMPTY != obj.Description {
		setFields = append(setFields, "DESCRIPTION = ?")
		args = append(args, obj.Description)
	}

	if EMPTY != obj.Solution {
		 setFields = append(setFields,"SOLUTION = ?")
		 args = append(args, obj.Solution)
	}

	if EMPTY != obj.Occurs {
		setFields = append(setFields, "OCCURS = ?")
		args = append(args, obj.Occurs)
	}

	if EMPTY != obj.Server.Id {
		setFields = append(setFields, "ID_SERVER = ?")
		args = append(args, obj.Server.Id)
	}

	if EMPTY != obj.ErrorState.Id {
		setFields = append(setFields, "ID_ERROR_STATE = ?")
		args = append(args, obj.ErrorState.Id)
	}

	return strings.Join(setFields, ", "), args
}
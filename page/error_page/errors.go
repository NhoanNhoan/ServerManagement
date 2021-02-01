package error_page

import (
	"CURD/database"
	"CURD/entity/error_entity"

	"database/sql"
)

type Errors struct {
	ErrArr []error_entity.Error
}

func (errors *Errors) New() error {
	comp := errors.makeQueryComp()
	rows, err := database.Query(comp)
	if nil == err {
		errors.fetch(rows)
	}
	
	return err
}

func (errors *Errors) makeQueryComp() database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"ERROR", "ERROR_STATE"},
		
		Columns: []string {
			"ERROR.ID",
			"ERROR.SUMMARY",
			"ERROR.DESCRIPTION",
			"ERROR.SOLUTION",
			"ERROR.OCCURS",
			"ERROR.ID_SERVER",
			"ERROR.ID_ERROR_STATE",
			"ERROR_STATE.DESCRIPTION",
		},
		
		Selection: "ERROR.ID_ERROR_STATE = ERROR_STATE.ID",		
		SelectionArgs: nil,	
		GroupBy: "",
		Having: "",
		OrderBy: "",
		Limit: "",
	}
}

func (instance *Errors) fetch(rows *sql.Rows) (err error) {
	var obj error_entity.Error

	for rows.Next() {
		err = obj.Fetch(rows)
		
		if nil != err {
			return err
		}
		
		instance.ErrArr = append(instance.ErrArr, obj)
	}
	
	return nil
}
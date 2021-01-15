package model

import (
	"database/sql"
)

type ServerManagementHandler struct {
	connector *sql.DB
}

func (handler ServerManagementHandler) Connect() (err error) {
	handler.connector, err = sql.Open("sqlite3", DB_NAME)
	return err
}

func (handler ServerManagementHandler) ExecuteQuery(sql string) (rows *sql.Rows, err error) {
	rows, err = handler.connector.Query(sql)
	return
}

func (handler ServerManagementHandler) ExecuteNonQuery(sql string) (res sql.Result, err error) {
	statement, err := handler.connector.Prepare(sql)
	if nil == err {
		res, err = statement.Exec()
	}

	return
}

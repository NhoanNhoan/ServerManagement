package model

import (
	"database/sql"
)

type DatabaseHandler interface {
	ExecuteQuery(sql string) *sql.Rows
	ExecuteNonQuery(sql string) (sql.Result, error)
}
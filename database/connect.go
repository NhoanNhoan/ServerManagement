package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type SQLComponents struct {
	tables []string
	columns []string
	selection string
	selectionArgs []string
	groupBy string
	having string
	orderBy string
	limit string
}

func MakeQuery(components SQLComponents) string {
	return concat("SELECT", 
					makeClause(components.columns, ", "), 
					"FROM", 
					makeClause(components.tables, ", "), 
					"WHERE",
					replace(components.selection, selectionArgs, "?", 1)))
}

func makeClause(values []string, delim string) string {
	return strings.Join(values, delim)
}

func concat(clause ...string) string {
	return strings.Join(clause, " ")
}

func replace(s string, values []string, delim string, times uint) string {
	var ans strings.Builder

	for s, _ := range values {
		s = strings.Replace(s, delim, s, times)
	}

	return s
}

func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	//dbPass := "kali"
	dbName := "server_management"
	db, err := sql.Open(dbDriver, dbUser + "@/" + dbName)
	if nil != err {
		panic (err.Error())
	}

	return db
}

func Query(components SQLComponents) (*sql.Rows, error) {
	sql := MakeQuery(components)
	db := DBConn()
	defer db.Close()
	return db.Query(sql)
}
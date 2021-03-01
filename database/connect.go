package database

import (
	"database/sql"
	_ "fmt"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func DBConn() (db *sql.DB) {
	dbDriver := "sqlite3"
	//dbUser := "root"
	//dbPass := "kali"@
	//dbName := "ServerManagement"
	//dbUser + "@/" + dbName)
	// mysqlConString := "root:root@tcp(127.0.0.1:3306)/ServerManagement"
	db, err := sql.Open(dbDriver, "database/ServerManagement.db")
	if nil != err {
		panic(err.Error())
	}

	return db
}

func Query(component QueryComponent) (*sql.Rows, error) {
	db := DBConn()
	defer db.Close()

	sql := MakeQuery(component)
	stmp, err := db.Prepare(sql)

	if nil != err {
		panic(err)
	}

	args := toInterfaceSplice(component.SelectionArgs)
	rows, err := stmp.Query(args...)

	if nil != err {
		panic(err)
	}

	return rows, err
}

func Insert(component InsertComponent) (err error) {
	sql := MakeInsert(component)

	for _, value := range component.Values {
		err = executeStatement(sql, value...)

		if nil != err {
			return err
		}
	}

	return err
}

func Update(component UpdateComponent) (err error) {
	sql := MakeUpdateStatement(component)
	concatenation := append(component.Values,
		component.SelectionArgs...)
	return executeStatement(sql, concatenation...)
}

func Delete(component DeleteComponent) (err error) {
	sql := MakeDeleteStatement(component)
	return executeStatement(sql, component.SelectionArgs...)
}

func executeStatement(statement string, values ...string) (err error) {
	db := DBConn()
	defer db.Close()

	stmp, err := db.Prepare(statement)

	if nil != err {
		panic(err)
	}

	args := toInterfaceSplice(values)
	_, err = stmp.Exec(args...)

	return err
}

func MakeQueryAll(tables []string, columns []string) QueryComponent {
	return QueryComponent{
		Tables:        tables,
		Columns:       columns,
		Selection:     "",
		SelectionArgs: nil,
		GroupBy:       "",
		Having:        "",
		OrderBy:       "",
		Limit:         "",
	}
}

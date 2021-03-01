package entity

import (
	"CURD/database"
)

type User struct {
	Id, Username, Password string
}

func (user *User) IsExists() bool {
	comp := user.makeQueryExistsComp()
	rows, err := database.Query(comp)
	defer rows.Close()

	return nil == err && rows.Next()
}

func (user *User) makeQueryExistsComp() database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"USER"},
		Columns: []string {"ID"},
		Selection: "USERNAME = ? and PASS = ?",
		SelectionArgs: []string {user.Username, user.Password},
	}
}
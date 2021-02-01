package error_entity

import (
	"CURD/database"
	"database/sql"
)

type Person struct {
	Id, Name string
}

func (p *Person) New(IdPerson string) error {
		p.Id = IdPerson
		row := p.queryPerson()
		return p.FetchPerson(row)
}

func (p *Person) queryPerson() *sql.Rows {
	comp := database.QueryComponent {
		Tables: []string {
				"PERSON",
			},
			
		Columns: []string {
				"ID",
				"NAME",
			},
			
		Selection: "ID = ?",
		
		SelectionArgs: []string {p.Id},
		
		GroupBy: "",
		Having: "",
		OrderBy: "",
		Limit: "",
	}
	
	row, err := database.Query(comp)
	
	if nil == err {
		return row
	}
	
	return nil
}

func (p *Person) FetchPerson(row *sql.Rows) (err error) {
	if nil != row && row.Next() {
		err = row.Scan(
				&p.Id,
				&p.Name,
			)
	}
	
	return
}

func FetchPersons() []Person {
	persons := make([]Person, 0)
	var err error
	var p Person
	rows := queryPersons()

	for rows.Next() {
		err = rows.Scan(&p.Id,
					&p.Name)
		if nil != err {
			panic (err)
		}

		persons = append(persons, p)
	}

	return persons
}

func queryPersons() *sql.Rows {
	comp := personsComp()
	rows, err := database.Query(comp)

	if nil != err {
		panic (err)
	}

	return rows
}

func personsComp() database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {
				"PERSON",
			},
			
		Columns: []string {
				"ID",
				"NAME",
			},
			
		Selection: "",
		SelectionArgs: nil,
		GroupBy: "",
		Having: "",
		OrderBy: "",
		Limit: "",
	}
}
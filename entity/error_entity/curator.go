package error_entity

import (
	"CURD/database"
	
	"database/sql"
)

type icomp = database.InsertComponent

type Curator struct {
	Id string
	IdError string
	IdPerson string
}

func (c *Curator) MakeByRows(row *sql.Rows) error {
	err := row.Scan(&c.Id,
					&c.IdError,
					&c.IdPerson)
	return err
}

func (c *Curator) Execute() (err error) {
	if c.IsExists(c.IdPerson) {
		return c.update()
	}

	return c.insert()
}

func (c *Curator) update() (err error) {
	comp := c.makeUComp()
	return database.Update(comp)
}

func (c *Curator) makeUComp() ucomp {
	return ucomp {
		Table: "CURATOR",
		SetClause: "ID_ERROR = ?, " +
					"ID_PERSON = ?",
		Values: []string {
			c.IdError,
			c.IdPerson,
		},
		Selection: "ID = ?",
		SelectionArgs: []string {c.Id},
	}
}

func (c *Curator) insert() (err error) {
	comp := c.makeInsComp()
	return database.Insert(comp)
}

func (c *Curator) makeInsComp() icomp {
	return icomp {
		Table: "CURATOR",
		Columns: []string {
			"ID",
			"ID_ERROR",
			"ID_PERSON",
		},
		Values: [][]string {
			[]string {
				c.Id,
				c.IdError,
				c.IdPerson,
			},
		},
	}
}
// CHECK EXITS AREA

func (c *Curator) IsExists(IdPerson string) bool {
	comp := c.makeExistsQueryComp(IdPerson)
	row, err := database.Query(comp)
	return nil == err && row.Next()
}

func (c *Curator) makeExistsQueryComp(IdPerson string) qcomp {
	return qcomp {
		Tables: []string {"CURATOR AS C",
					"PERSON AS P",
				},

		Columns: []string {
					"C.ID",
				},

		Selection: "C.ID = ? AND " +
				"C.ID_ERROR = ? AND " +
				"C.ID_PERSON = ?",

		SelectionArgs: []string {c.Id, c.IdError, IdPerson},
	}
}
// END OF CHECK EXISTS AREA

// FETCH CURATOR ARRAY AREA

func FetchCurators(IdError string) ([]Curator, error) {
	comp := makeCuratorsComp(IdError)
	rows, err := database.Query(comp)
	curators := make([]Curator, 0)
	var curator Curator

	for (nil == err) && rows.Next() {
		err = curator.MakeByRows(rows)
		curators = append(curators, curator)
	}

	if nil != err {
		panic (err)
	}

	return curators, err
}

func makeCuratorsComp(IdError string) qcomp {
	return qcomp {
		Tables: []string {"CURATOR"},

		Columns: []string {
					"ID",
					"ID_ERROR", 
					"ID_PERSON",
				},

		Selection: "ID_ERROR = ?",

		SelectionArgs: []string {IdError},
	}
}

// END OF FETCH CURATOR AREA
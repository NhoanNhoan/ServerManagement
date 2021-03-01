package error_entity

import (
	"CURD/database"
	
	"database/sql"
)

type Curator struct {
	//Id string
	IdError string
	IdPerson string
}

func (c *Curator) MakeByRows(row *sql.Rows) error {
	err := row.Scan(//&c.Id,
					&c.IdError,
					&c.IdPerson)
	return err
}

// Execute performs update or insert
// If id has been exists, it will be updated
// otherwise this curator will be inserted
func (c *Curator) Execute() (err error) {
	if WasExistedCurator(*c) {
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
		SetClause: "ID_PERSON = ?",
		Values: []string {
			c.IdPerson,
		},
		Selection: "ID_ERROR = ? AND ID_PERSON = ?",
		SelectionArgs: []string {c.IdError, 
								c.IdPerson},
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
			"ID_ERROR",
			"ID_PERSON",
		},
		Values: [][]string {
			[]string {
				c.IdError,
				c.IdPerson,
			},
		},
	}
}

// CHECK EXITS AREA

func (c *Curator) IsExistsPerson(IdPerson string) bool {
	comp := c.makeExistsQueryComp(IdPerson)
	row, err := database.Query(comp)
	defer row.Close()
	return nil == err && row.Next()
}

func (c *Curator) makeExistsQueryComp(IdPerson string) qcomp {
	return qcomp {
		Tables: []string {"CURATOR AS C",
				},

		Columns: []string {
					"C.ID_ERROR",
				},

		Selection: "C.ID_ERROR = ? AND " +
					"C.ID_PERSON = ?",

		SelectionArgs: []string {c.IdError, IdPerson},
	}
}
// END OF CHECK EXISTS AREA

// ===========================

// FETCH CURATOR ARRAY AREA

func FetchCurators(IdError string) ([]Curator, error) {
	comp := makeCuratorsComp(IdError)
	rows, err := database.Query(comp)
	defer rows.Close()
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
					"ID_ERROR", 
					"ID_PERSON",
				},

		Selection: "ID_ERROR = ?",

		SelectionArgs: []string {IdError},
	}
}

// END OF FETCH CURATOR AREA

// MAKE NEW PRIMARY FOR A CURATOR

func WasExistedCurator(curator Curator) bool {
	comp := qcomp {
		Tables: []string {"CURATOR"},
		Columns: []string {"ID_ERROR"},
		Selection: "ID_ERROR = ? AND ID_PERSON = ?",
		SelectionArgs: []string {curator.IdError, 
							curator.IdPerson},
	}

	row, err := database.Query(comp)
	defer row.Close()
	return (nil == err) && row.Next()
}

// 

func DeleteCuratorsByErrorId(ErrorId string) error {
	comp := database.DeleteComponent {
		Table: "CURATOR",
		Selection: "ID_ERROR = ?",
		SelectionArgs: []string {ErrorId},
	}

	return database.Delete(comp)
} 
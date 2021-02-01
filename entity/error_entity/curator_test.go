package error_entity

import (
	"CURD/database"
	"testing"
)

func TestMakeByRows(t *testing.T) {
	var curator Curator
	comp := makeCuratorsComp("ERR001")
	rows, err := database.Query(comp)

	if nil != err {
		panic (err)
	}
	rows.Next()
	err = curator.MakeByRows(rows)
	if nil != err {
		panic (err)
	}

	if "" == curator.Id || 
			"" == curator.IdError ||
			"" == curator.IdPerson {
		t.Error ("Fail", curator)
	}
}

func TestFetchCurators(t *testing.T) {
	curators, _ := FetchCurators("ERR001")
	if 0 == len(curators) {
		t.Error ("Fail")
	}
}
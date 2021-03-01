package database

import (
	"testing"
)

func TestDBConn(t *testing.T) {
	conn := DBConn()
	if nil == conn {
		t.Error("Can't connect")
	}
}

func TestQuery(t *testing.T) {
	comp := QueryComponent {
		Tables: []string {"DC"},
		Columns: []string {"id", "Description"},
	}

	rows, err := Query(comp)

	if nil != err {
		t.Error ("Query")
	}
}
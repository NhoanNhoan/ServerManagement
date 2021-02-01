package database

import (
	"testing"
	"strings"
)

func TestMakeUpdateStatement(t *testing.T) {
	components := []UpdateComponent {
		UpdateComponent {
			table: "DC",
			setClause: "description = ?",
			values: nil,
			selection: "id = (select s.id from server s where s.id = ?)",
			selectionArgs: nil,
		},

		UpdateComponent {
			table: "RACK",
			setClause: "description = ?",
			values: nil,
			selection: "id = ?",
			selectionArgs: nil,
		},

		UpdateComponent {
			table: "SERVER",
			setClause: "id_DC = (select d.id from dc d where d.description = ?)",
			values: nil,
			selection: "serial_number = ? AND id_port_type = ?",
			selectionArgs: nil,
		},
	}

	expected := []string {
		"UPDATE DC "+ 
		"SET DESCRIPTION = ? " +
		"WHERE id = (select s.id from server s where s.id = ?)",

		"UPDATE RACK " +
		"SET DESCRIPTION = ? " +
		"WHERE ID = ?",

		"UPDATE SERVER " +
		"SET id_DC = (select d.id from dc d where d.description = ?) " +
		"WHERE serial_number = ? AND id_port_type = ?",
	}

	for i := range components {
		ans := MakeUpdateStatement(components[i])
		if strings.ToUpper(ans) != strings.ToUpper(expected[i]) {
			t.Error(ans, " should equal ", expected[i])
		}
	}
}

func TestUpdate(t *testing.T) {
	components := []UpdateComponent {
		UpdateComponent {
			table: "DC",
			setClause: "description = ?",
			values: []string {"DC 9"},
			selection: "ID = (SELECT S.ID_DC FROM SERVER S WHERE S.ID = ?)",
			selectionArgs: []string {"SV0000000001"},
		},

		UpdateComponent {
			table: "RACK",
			setClause: "description = ?",
			values: []string {"RACK 01"},
			selection: "id = ?",
			selectionArgs: []string {"RK0001"},
		},

		UpdateComponent {
			table: "SERVER",
			setClause: "ID_DC = (SELECT D.ID FROM DC D WHERE D.DESCRIPTION = ?)",
			values: []string {"DATA CENTER 09"},
			selection: "SERIAL_NUMBER = ? AND ID_PORT_TYPE = ?",
			selectionArgs: []string {"AWERHadaADS", "PT01"},
		},
	}

	for i := range components {
		err := Update(components[i])

		if nil != err {
			panic (err)
		}
	}
}
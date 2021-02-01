package database

import (
	"testing"
	"strings"
)

func TestMakeDeleteStatement(t *testing.T) {
	components := []DeleteComponent {
		DeleteComponent {
			table: "RACK",
			selection: "description = ?",
			selectionArgs: nil,
		},

		DeleteComponent {
			table: "PORT_TYPE",
			selection: "id = ?",
			selectionArgs: nil,
		},

		DeleteComponent {
			table: "SERVER",
			selection: "SERIAL_NUMBER = ? AND " +
			"ID_U_START = (SELECT ID " +
				"FROM RACK_UNIT " + 
				"WHERE DESCRIPTION = ?",
			selectionArgs: nil,
		},
	}

	expected := []string {
		"DELETE FROM RACK WHERE DESCRIPTION = ?",

		"DELETE FROM PORT_TYPE WHERE ID = ?",

		"DELETE FROM SERVER WHERE " + 
				"SERIAL_NUMBER = ? AND " +
				"ID_U_START = (SELECT ID " +
				"FROM RACK_UNIT " + 
				"WHERE DESCRIPTION = ?",
	}

	for i := range components {
		ans := MakeDeleteStatement(components[i])

		if strings.ToUpper(ans) != strings.ToUpper(expected[i]) {
			t.Error (ans, " should be equal ", expected[i])
		}
	}
}

func TestDelete(t *testing.T) {
	components := []DeleteComponent {
		DeleteComponent {
			table: "SERVER",
			selection: "SERIAL_NUMBER = ? AND " +
			"ID_U_START IN (SELECT R.ID " +
				"FROM RACK_UNIT AS R " + 
				"WHERE DESCRIPTION = ?)",
			selectionArgs: []string {"AWERHgadaADS", "U 30"},
		},
	}

	for i := range components {
		err := Delete(components[i])

		if nil != err {
			t.Error(MakeDeleteStatement(components[i]))
			panic (err)
		}
	}
}
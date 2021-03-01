package database

import (
	"testing"
)

// func TestDuplicateString(t *testing.T) {
// 	testCases := []string {
// 		duplicate("?", ", ", 7),
// 	}

// 	expected := []string {
// 		"?, ?, ?, ?, ?, ?, ?",
// 	}

// 	for i := range testCases {
// 		ans := testCases[i]
// 		if ans != expected[i] {
// 			t.Error (ans, " should equal ", expected[i])
// 		}
// 	}
// }

func TestMakeInsert(t *testing.T) {
	components := []InsertComponent {
		InsertComponent {
			"SERVER",
			[]string {"id", "name", "status", "maker", "service"},
			[][]string {{"ABCD", "SERVER_NAME", "ACTIVE", "DELL", "NOTIFICATION"}},
		},

		InsertComponent {
			"SERVER",
			[]string {"id", "name", "status", "maker", "service"},
			[][]string {{"ABCD", "SERVER_NAME", "ACTIVE", "DELL", "NOTIFICATION"},
					{"ABCD", "SERVER_NAME", "ACTIVE", "DELL", "NOTIFICATION"},
					{"ABCD", "SERVER_NAME", "ACTIVE", "DELL", "NOTIFICATION"},
				{"ABCD", "SERVER_NAME", "ACTIVE", "DELL", "NOTIFICATION"}},
		},
	}

	expected := []string {
		"INSERT INTO SERVER(id, name, status, maker, service) VALUES(?, ?, ?, ?, ?)",
		"INSERT INTO SERVER(id, name, status, maker, service) VALUES(?, ?, ?, ?, ?)",
	}

	for i := 0; i < len(components); i++ {
		if MakeInsert(components[i]) != expected[i] {
			t.Error(MakeInsert(components[i]), " should equal ", expected[i])
		}
	}
}

func TestInsert(t *testing.T) {
	components := []InsertComponent {
		InsertComponent {
			"STATUS_ROW",
			[]string {"id", "description"},
			[][] string {{"1", "active"}, {"2", "stop"},},
		},

		InsertComponent {
			"DC",
			[]string {"id", "description"},
			[][] string {{"DC0001", "DC 08"}, {"DC0002", "DC 08"},},
		},

		InsertComponent {
			"RACK",
			[]string {"id", "description"},
			[][] string {{"RK0001", "RACK 01"}, {"RK0002", "RACK 02"},},
		},

		InsertComponent {
			"RACK_UNIT",
			[]string {"id", "description"},
			[][] string {{"RU0001", "U 30"}, {"RU0002", "U31"},},
		},

		InsertComponent {
			"PORT_TYPE",
			[]string {"id", "description"},
			[][] string {{"PT01", "IDRAC"}, {"PT02", "ILO"},},
		},

		InsertComponent {
			"SERVER_STATUS",
			[]string {"id", "description"},
			[][] string {{"SS0001", "ACTIVE"}, {"SS0002", "INTERACTIVE"},},
		},

		InsertComponent {
			"IP_NET",
			[]string {"id",  "value"},
			[][] string {{"IP0001", "42.118.242."}, {"IP0002", "192.168.1."},},
		},

		InsertComponent {
			"SERVER",
			[]string {"id", "id_DC", "id_RACK", 
						"id_U_start", "id_U_end", "num_disks",
						"id_PORT_TYPE", "serial_number", "id_SERVER_STATUS",
						"maker", "id_STATUS_ROW",
					},
			[][] string {
				{"SV0000000001", "DC0001", "RK0001", "RU0001", "RU0001", "20", "PT01", "AWERHgadaADS", "SS0001", "DELL", "1",},
				{"SV0000000002", "DC0002", "RK0002", "RU0001", "RU0002", "20", "PT02", "ADFEeawADFAF", "SS0002", "HP", "1",},
			},
		},
	}

	for idx := range components {
		err := Insert(components[idx])
		if err != nil {
			panic (err)
		}
	}
}
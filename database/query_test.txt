package database

import (
	"testing"
	"strings"
)

func TestMakeQuery(t *testing.T) {
	components := []QueryComponent {
			QueryComponent {
				tables: []string {"SERVER"},
				columns: []string {"id", "id_DC", "id_RACK", "num_disks", "serial_number"},
				selection: "id = ?",
				selectionArgs: []string {"SV0000000001"},
				groupBy: "",
				having: "",
				orderBy:"",
				limit: "",
			},

			QueryComponent {
				tables: []string {"SERVER"},
				columns: []string {"id_SERVER_STATUS"},
				selection:"id_STATUS_ROW = ?",
				selectionArgs: []string {"1"},
				groupBy: "",
				having: "",
				orderBy: "",
				limit: "",
			},

			QueryComponent {
				tables: []string {"SERVER as S", "DC as D", "RACK as R", "RACK_UNIT as UStart", "RACK_UNIT as UEnd", "PORT_TYPE as P", "SERVER_STATUS as SS", "STATUS_ROW as SR"},
				columns: []string {"S.id", "D.description", "R.description", "UStart.description", "UEnd.description", "P.description", "SS.description", "SR.description"},
				selection: "S.id = ? AND S.id_DC = D.id AND S.id_RACK = R.id AND S.id_U_start = UStart.id AND S.id_U_end = UEnd.id AND S.id_PORT_TYPE = P.id AND S.id_SERVER_STATUS = SS.id AND S.id_STATUS_ROW = SR.id",
				selectionArgs:[]string {"SV000000001"},
				groupBy: "",
				having: "",
				orderBy:"",
				limit: "",
			},

			QueryComponent {
				tables: []string {"DC"},
				columns: []string {"id", "description"},
				selection:"",
				selectionArgs: nil,
				groupBy: "",
				having: "",
				orderBy: "",
				limit: "",
			},
	}

	exptected := []string {
		"SELECT id, id_DC, id_RACK, num_disks, serial_number " +
		"FROM SERVER " +
		"WHERE id = ?",

		"SELECT id_SERVER_STATUS " +
		"FROM SERVER " +
		"WHERE id_STATUS_ROW = ?",

		"select s.id, d.description, r.description, " + 
		"ustart.description, uend.description, p.description, " +
		"ss.description, sr.description " +
		"from server as s, dc as d, rack as r, rack_unit as ustart," +
		" rack_unit as uend, port_type as p, server_status as ss, " + 
		"status_row as sr where S.id = ? AND S.id_DC = D.id " + 
		"AND S.id_RACK = R.id AND S.id_U_start = UStart.id AND " +
		"S.id_U_end = UEnd.id AND S.id_PORT_TYPE = P.id AND " + 
		"S.id_SERVER_STATUS = SS.id AND S.id_STATUS_ROW = SR.id",

		"select id, description from dc ",
	}

	for i := range components {
		ans := MakeQuery(components[i])

		if strings.ToUpper(ans) != strings.ToUpper(exptected[i]) {
			t.Error(strings.ToUpper(ans), " should equal ", strings.ToUpper(exptected[i]))
		}
	}
}

func TestQuery(t *testing.T) {
	components := []QueryComponent {
			QueryComponent {
				tables: []string {"SERVER"},
				columns: []string {"id", "id_DC", "id_RACK", "num_disks", "serial_number"},
				selection: "id = ?",
				selectionArgs: []string {"SV0000000001"},
				groupBy: "",
				having: "",
				orderBy:"",
				limit: "",
			},

			QueryComponent {
				tables: []string {"SERVER"},
				columns: []string {"id_SERVER_STATUS"},
				selection:"id_STATUS_ROW = ?",
				selectionArgs: []string {"1"},
				groupBy: "",
				having: "",
				orderBy: "",
				limit: "",
			},

			QueryComponent {
				tables: []string {"SERVER as S", "DC as D", "RACK as R", "RACK_UNIT as UStart", "RACK_UNIT as UEnd", "PORT_TYPE as P", "SERVER_STATUS as SS", "STATUS_ROW as SR"},
				columns: []string {"S.id", "D.description", "R.description", "UStart.description", "UEnd.description", "P.description", "SS.description", "SR.description"},
				selection: "S.id = ? AND S.id_DC = D.id AND S.id_RACK = R.id AND S.id_U_start = UStart.id AND S.id_U_end = UEnd.id AND S.id_PORT_TYPE = P.id AND S.id_SERVER_STATUS = SS.id AND S.id_STATUS_ROW = SR.id",
				selectionArgs:[]string {"SV000000001"},
				groupBy: "",
				having: "",
				orderBy:"",
				limit: "",
			},

	}

	for i := range components {
		_, err := Query(components[i])
		if nil != err {
			panic (err)
		}
	}
}
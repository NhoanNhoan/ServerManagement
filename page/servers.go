package page

import (
	"CURD/database"
	"CURD/entity"
)

type Servers struct {
	entity.DataCenter
	Items []entity.Server
	Tags []entity.Tag
}

func (s *Servers) New(IdDC string) {
	component := makeQueryComponent(IdDC)
	rows, err := database.Query(component)
	defer rows.Close()

	if nil != err {
		panic(err)
	}

	for rows.Next() {
		var id,
			rackName, 
			ustartName, 
			uendName, 
			ssd,
			hdd,
			portType, 
			serverStatus, 
			maker string

		err = rows.Scan(&id, 
			&rackName, 
			&ustartName, 
			&uendName, 
			&ssd, 
			&hdd,
			&portType, 
			&serverStatus, 
			&maker)

		var server entity.Server

		server.New(id)
		server.FetchIpAddrs()

		s.Items = append(s.Items, server)
	}

	s.Tags = entity.FetchAllTags()
}

// select s.id,
//     case when s.id_rack is not null and s.id_rack = r.id then r.description else null end,
//     case when s.id_u_start is not null and s.id_u_start = ustart.id then ustart.description else null end,
//     case when s.id_u_end is not null and s.id_u_end = uend.id then uend.description else null end,
//     case when s.id_port_type is not null and s.id_port_type = p.id then p.description else null end,
//     case when s.id_server_status is not null and s.id_server_status = ss.id then ss.description else null end,
//     s.num_disks, s.maker
//     from server as s,
//     rack as r,
//     rack_unit as ustart,
//     rack_unit as uend,
//     port_type as p,
//     server_status as ss
//     where s.id_dc = 'DC00001'

func makeQueryComponent(IdDC string) database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"SERVER", 
						"RACK", 
						"RACK_UNIT AS USTART", 
						"RACK_UNIT AS UEND", 
						"PORT_TYPE",
						"SERVER_STATUS"},
		Columns: []string {"SERVER.ID",
						"RACK.Description",
						"USTART.Description",
						"UEND.Description",
						"SSD",
						"HDD",
						"PORT_TYPE.Description",
						"SERVER_STATUS.Description",
						"SERVER.MAKER"},
		Selection: "SERVER.ID_DC = ? AND " +
				"SERVER.ID_RACK = RACK.ID AND " + 
				"SERVER.ID_U_START = USTART.ID AND " + 
				"SERVER.ID_U_END = UEND.ID AND " + 
				"SERVER.ID_PORT_TYPE = PORT_TYPE.ID AND " + 
				"SERVER.ID_SERVER_STATUS = SERVER_STATUS.ID",
		SelectionArgs: []string {IdDC},
	}
}

func (s *Servers) GetServersByTagId(tagId string) {
	comp := makeQueryCompByTagId(tagId, s.DataCenter.Id)
	rows, err := database.Query(comp)
	defer rows.Close()

	if nil != err {
		panic(err)
	}

	for rows.Next() {
		var id,
			rackName, 
			ustartName, 
			uendName, 
			ssd, 
			hdd,
			portType, 
			serialNumber, 
			serverStatus, 
			maker string

		err = rows.Scan(&id, 
			&rackName, 
			&ustartName, 
			&uendName, 
			&ssd, 
			&hdd,
			&portType, 
			&serialNumber, 
			&serverStatus, 
			&maker)

		var server entity.Server

		server.New(id)
		server.FetchIpAddrs()

		s.Items = append(s.Items, server)
	}

	s.Tags = entity.FetchAllTags()
}

func makeQueryCompByTagId(tagId string, dcId string) database.QueryComponent {
	return database.QueryComponent {
			Tables: []string {"SERVER", 
							"RACK", 
							"RACK_UNIT AS USTART", 
							"RACK_UNIT AS UEND", 
							"PORT_TYPE",
							"SERVER_STATUS",
							"STATUS_ROW",
							"SERVER_TAG AS ST"},
			Columns: []string {"SERVER.ID",
							"RACK.Description",
							"USTART.Description",
							"UEND.Description",
							"SERVER.SSD",
							"SERVER.HDD",
							"PORT_TYPE.Description",
							"SERVER.SERIAL_NUMBER",
							"SERVER_STATUS.Description",
							"SERVER.MAKER"},
			Selection: "SERVER.ID_DC = ? AND " +
					"SERVER.ID_RACK = RACK.ID AND " + 
					"SERVER.ID_U_START = USTART.ID AND " + 
					"SERVER.ID_U_END = UEND.ID AND " + 
					"SERVER.ID_PORT_TYPE = PORT_TYPE.ID AND " + 
					"SERVER.ID_SERVER_STATUS = SERVER_STATUS.ID AND " + 
					"STATUS_ROW.DESCRIPTION = ? AND SERVER.ID_STATUS_ROW = STATUS_ROW.ID AND " +
					"ST.TAGID = ? AND SERVER.ID = ST.SERVERID",
			SelectionArgs: []string {dcId, "available", tagId},
			GroupBy: "",
			Having: "",
			OrderBy: "",
			Limit: "",
		}
}

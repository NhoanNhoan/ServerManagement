package page

import (
	_ "fmt"
	"CURD/database"
	"CURD/entity"
)

type Servers struct {
	entity.DataCenter
	Items []entity.Server
}

func (s *Servers) New(IdDC string) {
	component := makeQueryComponent(IdDC)
	rows, err := database.Query(component)

	if nil != err {
		panic(err)
	}

	for rows.Next() {
		var id,
			rackName, 
			ustartName, 
			uendName, 
			numDisk, 
			portType, 
			serialNumber, 
			serverStatus, 
			maker string

		err = rows.Scan(&id, 
			&rackName, 
			&ustartName, 
			&uendName, 
			&numDisk, 
			&portType, 
			&serialNumber, 
			&serverStatus, 
			&maker)

		var server entity.Server

		server.New(id)
		server.FetchIpAddrs()

		s.Items = append(s.Items, server)
	}

}

func makeQueryComponent(IdDC string) database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"SERVER", 
						"RACK", 
						"RACK_UNIT AS USTART", 
						"RACK_UNIT AS UEND", 
						"PORT_TYPE",
						"SERVER_STATUS",
						"STATUS_ROW"},
		Columns: []string {"SERVER.ID",
						"RACK.Description",
						"USTART.Description",
						"UEND.Description",
						"NUM_DISKS",
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
				"STATUS_ROW.DESCRIPTION = ? AND SERVER.ID_STATUS_ROW = STATUS_ROW.ID;",
		SelectionArgs: []string {IdDC, "available"},
		GroupBy: "",
		Having: "",
		OrderBy: "",
		Limit: "",
	}
}

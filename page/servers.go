package page

import (
	"CURD/database"
	"CURD/model"
	"database/sql"
)

type Servers struct {
	model.DataCenter
	Items []model.Server
}

type ServerItem struct {
	model.Server
	IpAddrs []string
}

func (s *Servers) New(Id string) {
	db:= database.DBConn()
	defer db.Close()

	sql := makeQuery(Id)

	rows, err := db.Query(sql)

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

		dc := model.DataCenter {
			Name: s.DataCenter.Name,
		}

		rack := model.Rack {
			Name: rackName,
		}

		ustart := model.RackUnit {
			Name: ustartName,
		}

		uend := model.RackUnit {
			Name: uendName,
		}

		pType := model.PortType {
			Name: portType,
		}

		location := model.Location {
			DataCenter: dc,
			Rack: rack,
			UStart: ustart,
			UEnd: uend,
		}

		server := ServerItem {
			Id: id,
			Location: &location,
			SerialNumber: serialNumber,
			PortType: &pType,
			IpAddrs: s.fetchIPAddrs(db, Id),
		}

		s.Items = append(s.Items, server)
	}
}

func makeQuery(Id string) string {
	return "select SERVER.id, RACK.name, " +
					"ustart.name, uend.name, num_disks, " +
					"PORT_TYPE.name, serial_number, " +
					"SERVER_STATUS.status, SERVER.maker " + 
			"from SERVER, RACK, " +
				"RACK_UNIT as ustart, RACK_UNIT as uend, PORT_TYPE, " + 
				"SERVER_STATUS, STATUS_ROW " +
			"where SERVER.id_DC = '" + Id + "' and " + 
				"SERVER.id_Rack = RACK.id and " + 
				"SERVER.id_U_start = ustart.id and " + 
				"SERVER.id_U_end = uend.id and " + 
				"SERVER.id_PORT_TYPE = PORT_TYPE.id and " + 
				"SERVER.id_SERVER_STATUS = SERVER_STATUS.id and " + 
				"STATUS_ROW.status = 'available' and SERVER.id_STATUS_ROW = STATUS_ROW.id;";
}

func (s *Servers) fetchIPAddrs(db *sql.DB, IdServer string) []string {
	sql := makeQueryIPString(IdServer)
	rows, err := db.Query(sql)

	if nil != err {
		panic(err)
	}

	ipAddrs := scanRows(rows)

	return ipAddrs
}

func scanRows(rows *sql.Rows) (ipAddrs []string) {
	var ipNet, ipHost string

	for rows.Next() {
		var ipNet, ipHost string
		scanErr := rows.Scan(&ipNet, &ipHost)

		if nil != scanErr {
			panic(scanErr)
		}

		ipAddrs = append(ipAddrs, ipNet + ipHost)
	}

	return
}

func makeQueryIPString(IdServer string) {
	return "select id_IP_NET, ip_host " +
		"from IP_SERVER, STATUS_ROW as ST" +
		"where id_Server = '" + IdServer + "' AND " +
			"ST.status = 'available' AND IP_SERVER.id_STATUS_ROW = ST.id"
}

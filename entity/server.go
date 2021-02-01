package entity

import (
	"CURD/database"
	"database/sql"
)

type qcomp = database.QueryComponent
type ucomp = database.UpdateComponent

type Server struct {
	Id string
	DC DataCenter
	Rack
	UStart RackUnit
	UEnd RackUnit
	NumDisks string
	Maker string
	PortType
	SerialNumber string
	ServerStatus
	IpAddrs []IpAddress
	Services []string
	Events []string
}

func (obj *Server) New(Id string) (err error) {
	obj.SetId(Id)
	component := obj.makeQueryComponent(Id)
	rows, err := database.Query(component)

	if nil != err {
		return err
	}

	if rows.Next() {
		err = rows.Scan(
					&obj.DC.Id,
					&obj.DC.Description,
					&obj.Rack.Id,
					&obj.Rack.Description,
					&obj.UStart.Id,
					&obj.UStart.Description,
					&obj.UEnd.Id,
					&obj.UEnd.Description,
					&obj.NumDisks,
					&obj.Maker,
					&obj.PortType.Id,
					&obj.PortType.Description,
					&obj.SerialNumber,
					&obj.ServerStatus.Id,
					&obj.ServerStatus.Description,
			)
	}

	return err
}

func (obj *Server) SetId(Id string) {
	obj.Id = Id
}

func (obj *Server) makeQueryComponent(IdServer string) qcomp {
	return qcomp {
		Tables: []string {
				"SERVER AS S",
				"DC AS D",
				"RACK AS R",
				"RACK_UNIT AS USTART",
				"RACK_UNIT AS UEND",
				"PORT_TYPE AS PT",
				"SERVER_STATUS AS SS",
				"STATUS_ROW AS SR",
			},
			
		Columns: []string {
				"D.ID",
				"D.DESCRIPTION",
				"R.ID",
				"R.DESCRIPTION",
				"USTART.ID", 
				"USTART.DESCRIPTION",
				"UEND.ID",
				"UEND.DESCRIPTION",
				"S.NUM_DISKS",
				"S.MAKER",
				"PT.ID",
				"PT.DESCRIPTION",
				"S.SERIAL_NUMBER",
				"SS.ID",
				"SS.DESCRIPTION",
			},
			
		Selection: "S.ID = ? AND " +
					"S.ID_DC = D.ID AND " +
					"S.ID_RACK = R.ID AND " + 
					"S.ID_U_START = USTART.ID AND " + 
					"S.ID_U_END = UEND.ID AND " + 
					"SR.DESCRIPTION = ? AND S.ID_STATUS_ROW = SR.ID;",
						
		SelectionArgs: []string {IdServer, "available"},
		GroupBy: "",
		Having: "",
		OrderBy: "",
		Limit: "",
	}
}

func (obj *Server) FetchIpAddrs() {
	component := obj.makeIpQueryComponent()
	rows, err := database.Query(component)

	if nil != err {
		panic (err)
	}

	obj.initIpAddrs(rows)
}

func (obj Server) makeIpQueryComponent() qcomp {
	return qcomp {
		Tables: []string {"IP_SERVER", "IP_NET"},
		Columns: []string {"IP_NET.VALUE", "IP_SERVER.IP_HOST"},
		Selection: "IP_SERVER.ID_SERVER = ? AND IP_SERVER.ID_IP_NET = IP_NET.ID",
		SelectionArgs: []string {obj.Id},
		GroupBy: "",
		Having: "",
		OrderBy: "",
		Limit: "",
	}
}

func (obj *Server) initIpAddrs(rows *sql.Rows) {
	var ipNet, ipHost string
	var ipAddr IpAddress

	for rows.Next() && (nil == rows.Scan(&ipNet, &ipHost)) {
		ipAddr.New(ipNet, ipHost)
		obj.IpAddrs = append(obj.IpAddrs, ipAddr)
	}
}

func (obj *Server) FetchServices() {
	comp := obj.makeServicesQueryComponent()
	rows, err := database.Query(comp)

	if nil != err {
		panic (err)
	}

	obj.initServerServices(rows)
}

func (obj *Server) initServerServices(rows *sql.Rows) {
	var service string

	for rows.Next() && (nil == rows.Scan(&service)) {
		obj.Services = append(obj.Services, service)
	}
}

func (obj Server) makeServicesQueryComponent() database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"SERVICES"},
		Columns: []string {"SERVICE"},
		Selection: "id_SERVER = ?",
		SelectionArgs: []string {obj.Id},
		GroupBy: "",
		Having: "",
		OrderBy: "",
		Limit: "",
	}
}

func UpdateServer(server Server) (msg string) {
	comp := makeUpdateServerComponent(server)
	err := database.Update(comp)
	
	if nil != err {
		msg = "Can't update the server information"
		panic (err)
	} else {
		msg = "Success"
	}
	
	return
}

func makeUpdateServerComponent(server Server) ucomp {
	return ucomp {
		Table: "SERVER",
		SetClause: "id_DC = ?, " +
			"id_RACK = ?, " +
			"id_U_start = ?, " +
			"id_U_end  = ?, " +
			"num_disks = ?, " +
			"maker = ?, " +
			"id_PORT_TYPE = ?, " +
			"serial_number = ?, " +
			"id_SERVER_STATUS = ?",
		Values: []string {
				server.DC.Id,
				server.Rack.Id,
				server.UStart.Id,
				server.UEnd.Id,
				"12",
				server.Maker,
				server.PortType.Id,
				server.SerialNumber,
				server.ServerStatus.Id,
			},
		Selection: "id = ?",
		SelectionArgs: []string {
				server.Id,
			},
	}
}

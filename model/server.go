package model

type Server struct {
	Id string
	*Location
	IPAddr	string
	Service string
	SerialNumber string
	*PortType
}

type PortType struct {
	Id, Name string
}

type IPServer struct {
	IdServer string
	IPNet string
	IPHost string
}

//	select id, dc.name, rack.name, ustart.name, uend.name, num_disk, port_type.name, serial_number, server_status.name, maker
//	from server, dc, rack, u as ustart, u as uend, port_type, server_status, status_row
//	where server.id_DC = dc.id and
//		server.id_Rack = rack.id and
//		server.id_U_start = ustart.id and
//		server.id_U_end = uend.id and
//		server.id_PORT_TYPE = port_type.id and
//		server.id_SERVER_STATUS = server_status.id and
//		status_row.name = 'available' and server.id_STATUS_ROW = status_row.id
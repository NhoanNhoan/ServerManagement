package server

import (
	"CURD/database"
	"CURD/entity"
	"database/sql"
	"strings"
)

var (
	AllFieldsServer = []string {
		"ID",
		"ID_DC",
		"ID_RACK",
		"ID_USTART",
		"ID_UEND",
		"HDD",
		"SSD",
		"MAKER",
		"ID_PORT_TYPE",
		"REDFISH_IP",
		"ID_SERVE",
		"SERIAL_NUMBER",
		"ID_SERVER_STATUS",
	}
)

type ServerRepo struct {
	repo    SqliteRepo
	ipRepo  IpRepo
	tagRepo TagRepo
}

func (s ServerRepo) Fetch(comp qcomp,
	scan func(interface{}, *sql.Rows) (interface{}, error)) ([]entity.Server, error) {
	entities, err := s.repo.Query(comp, func() interface{} {return entity.Server{}}, scan)

	if nil != err {
		return nil, err
	}

	servers := make([]entity.Server, len(entities))
	for i := range servers {
		servers[i] = entities[i].(entity.Server)
	}

	return servers, nil
}

func (s ServerRepo) FetchAll() ([]interface{}, error) {
	comp := s.makeFetchAllComp()
	makeServer := func() interface{} {return entity.Server{}}
	scanServer := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		server := obj.(entity.Server)
		err := row.Scan(&server.DC.Id, &server.DC.Description,
			&server.Rack.Id, &server.Rack.Description,
			&server.UStart.Id, &server.UStart.Description,
			&server.UEnd.Id, &server.UEnd.Description,
			&server.SSD, &server.HDD,
			&server.Maker,
			&server.PortType.Id,&server.PortType.Description,
			&server.RedfishIp,
			&server.PortType.Id, &server.PortType.Description,
			&server.SerialNumber,
			&server.ServerStatus.Id, &server.ServerStatus.Description,
			&server.ServerStatus.Id, &server.ServeCustomer.Description,
		)

		return server, err
	}

	entities, err := s.repo.Query(comp, makeServer, scanServer)
	if nil != err {
		return nil, err
	}

	return entities, nil
}

func (s ServerRepo) makeFetchAllComp() qcomp {
	return qcomp{
		Tables: []string{
			"SERVER",
			"DC", "RACK", "RACK_UNIT AS USTART", "RACK_UNIT AS UEND",
			"PORT_TYPE",
			"IP_ADDRESS",
			"SERVER_STATUS",
			"SERVE",
		},

		Columns: []string{
			"SERVER.ID",
			"DC.ID", "DC.DESCRIPTION",
			"RACK.ID", "RACK.DESCRIPTION",
			"USTART.ID", "USTART.DESCRIPTION",
			"UEND.ID", "UEND.DESCRIPTION",
			"SERVER.SSD", "SERVER.HDD",
			"SERVER.MAKER",
			"PORT_TYPE.ID", "PORT_TYPE.DESCRIPTION",
			"SERVER.REDFISH_IP",
			"SERVER.SERIAL_NUMBER",
			"SERVER.ID_SERVER_STATUS", "SERVER_STATUS.DESCRIPTION",
			"SERVER.ID_SERVE", "SERVE.DESCRIPTION",
		},

		Selection: strings.Join(
			[]string{
				"SERVER.ID_DC = ?",
				"SERVER.ID_RACK = ?",
				"SERVER.ID_U_START = ?",
				"SERVER.ID_U_END = ?",
				"SERVER.ID_PORT_TYPE = ?",
				"SERVER.ID_SERVER_STATUS = ?",
				"SERVER.ID_SERVE = ?",
			},
			" AND ",
		),

		SelectionArgs: []string{"DC.ID", "RACK.ID", "USTART.ID", "UEND.ID",
			"PORT_TYPE.ID", "SERVER_STATUS.ID", "SERVE.ID",
		},
	}
}

func (s ServerRepo) IsExists(id string) bool {
	row, err := database.Query(qcomp{
		Tables: []string {"SERVER"},
		Columns: []string {"ID"},
		Selection: "ID = ?",
		SelectionArgs: []string {id},
	})

	if nil != err {
		return false
	}
	defer row.Close()

	return row.Next()
}

func (s ServerRepo) GenerateId() (id string) {
	for {
		id = database.GeneratePrimaryKey(
			true,
			true,
			true,
			false,
			"SV",
			12)
		if !s.IsExists(id) {
			break
		}
	}

	return id
}

func (s ServerRepo) FetchById(Id string) ([]interface{}, error) {
	comp := s.makeFetchServerByIdComp(Id)
	makeServer := func() interface{} {return entity.Server{}}
	scanServer := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		server := obj.(entity.Server)
		err := row.Scan(&server.DC.Id, &server.DC.Description,
			&server.Rack.Id, &server.Rack.Description,
			&server.UStart.Id, &server.UStart.Description,
			&server.UEnd.Id, &server.UEnd.Description,
			&server.SSD, &server.HDD,
			&server.Maker,
			&server.PortType.Id,&server.PortType.Description,
			&server.RedfishIp,
			&server.PortType.Id, &server.PortType.Description,
			&server.SerialNumber,
			&server.ServerStatus.Id, &server.ServerStatus.Description,
			&server.ServerStatus.Id, &server.ServeCustomer.Description,
			)

		return server, err
	}

	entities, err := s.repo.Query(comp, makeServer, scanServer)
	if nil != err {
		return nil, err
	}

	return entities, nil
}

func (s ServerRepo) makeFetchServerByIdComp(id string) qcomp {
	return qcomp{
		Tables: []string {
			"SERVER",
			"DC", "RACK", "RACK_UNIT AS USTART", "RACK_UNIT AS UEND",
			"PORT_TYPE",
			"IP_ADDRESS",
			"SERVER_STATUS",
			"SERVE",
		},

		Columns: []string {
			"SERVER.ID",
			"DC.ID", "DC.DESCRIPTION",
			"RACK.ID", "RACK.DESCRIPTION",
			"USTART.ID", "USTART.DESCRIPTION",
			"UEND.ID", "UEND.DESCRIPTION",
			"SERVER.SSD", "SERVER.HDD",
			"SERVER.MAKER",
			"PORT_TYPE.ID", "PORT_TYPE.DESCRIPTION",
			"SERVER.REDFISH_IP",
			"SERVER.SERIAL_NUMBER",
			"SERVER.ID_SERVER_STATUS", "SERVER_STATUS.DESCRIPTION",
			"SERVER.ID_SERVE", "SERVE.DESCRIPTION",
		},

		Selection: strings.Join(
			[]string {
			"SERVER.ID = ?",
			"SERVER.ID_DC = ?",
			"SERVER.ID_RACK = ?",
			"SERVER.ID_U_START = ?",
			"SERVER.ID_U_END = ?",
			"SERVER.ID_PORT_TYPE = ?",
			"SERVER.ID_SERVER_STATUS = ?",
			"SERVER.ID_SERVE = ?",
			},
			" AND ",
			),

		SelectionArgs: []string {
			id, "DC.ID", "RACK.ID", "USTART.ID", "UEND.ID",
			"PORT_TYPE.ID", "SERVER_STATUS.ID", "SERVE.ID",
		},
	}
}

func (s ServerRepo) Insert(servers... entity.Server) error {
	for _, server := range servers {
		if s.IsExists(server.Id) || "" == server.Id {
			server.Id = s.GenerateId()
		}

		err := database.Insert(icomp{
			Table: "SERVER",
			Columns: []string {
				"ID",
				"ID_DC",
				"ID_RACK",
				"ID_U_START",
				"ID_U_END",
				"ID_PORT_TYPE",
				"ID_SERVER_STATUS",
				"SSD",
				"HDD",
				"MAKER",
				"SERIAL_NUMBER",
				"id_STATUS_ROW",
			},
			Values: [][]string {
				[]string {server.Id,
					server.DC.Id,
					server.Rack.Id,
					server.UStart.Id,
					server.UEnd.Id,
					server.PortType.Id,
					server.ServerStatus.Id,
					server.SSD,
					server.HDD,
					server.Maker,
					server.SerialNumber,
					"1",
				},
			},
		})

		if nil != err {
			return err
		}
	}

	return nil
}

func (s ServerRepo) Update(servers... interface{}) error {
	for _, obj := range servers {
		server := obj.(entity.Server)
		if err := database.Update(ucomp{
					Table: "SERVER",
					SetClause: strings.Join(
						[]string {
							"ID_DC = ?",
							"ID_RACK = ?",
							"ID_U_START = ?",
							"ID_U_END = ?",
							"SSD = ?",
							"HDD = ?",
							"ID_PORT_TYPE = ?",
							"SERIAL_NUMBER = ?",
							"ID_SERVER_STATUS = ?",
							"REDFISH_IP = ?",
							"ID_SERVE = ?",
						}, " AND "),
					Values: []string {
							server.DC.Id,
							server.Rack.Id,
							server.UStart.Id,
							server.UEnd.Id,
							server.SSD,
							server.HDD,
							server.PortType.Id,
							server.SerialNumber,
							server.ServerStatus.Id,
							server.RedfishIp,
							server.ServeCustomer.Id,
						},
					Selection: "id = ?",
					SelectionArgs: []string {
							server.Id,
						},
				}); nil != err {
			return err
		}
	}
	return nil
}

func (s ServerRepo) Delete(servers... interface{}) error {
	for _, obj := range servers {
		server := obj.(entity.Server)
		if err := database.Delete(dcomp{
			Table: "SERVER",
			Selection: "ID = ?",
			SelectionArgs: []string {server.Id},
		}); nil != err {
			return err
		}
	}

	return nil
}

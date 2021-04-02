package entity

import (
	"CURD/database"
	"database/sql"
)

type Switch struct {
	Id string
	Name string
	DC DataCenter
	Rack
	UStart RackUnit
	UEnd RackUnit
	MaximumPort string
	IpAddrs []IpAddress
}

func (obj *Switch) New(Id string) (err error) {
	component := obj.makeQueryComponent(Id)
	rows, err := database.Query(component)
	defer rows.Close()

	if nil != err {
		return err
	}

	if rows.Next() {
		err = rows.Scan(&obj.Id,
					&obj.DC.Description,
					&obj.Rack.Description,
					&obj.UStart.Description,
					&obj.UEnd.Description,
					&obj.MaximumPort,
					)
	}

	return err
}

func (obj *Switch) makeQueryComponent(IdSwitch string) database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"SWITCH AS SW",
						"DC",
						"RACK", 
						"RACK_UNIT AS USTART", 
						"RACK_UNIT AS UEND", 
						"STATUS_ROW"},
		Columns: []string {"SW.ID",
							"DC.DESCRIPTION",
							"RACK.DESCRIPTION",
							"USTART.DESCRIPTION",
							"UEND.DESCRIPTION",
							"SW.MAXIMUM_PORT",
						},
		Selection: "SW.ID = ? AND " +
					"SW.ID_DC = DC.ID AND " +
					"SW.ID_RACK = RACK.ID AND " + 
					"SW.ID_U_START = USTART.ID AND " + 
					"SW.ID_U_END = UEND.ID",
		SelectionArgs: []string {IdSwitch},
		GroupBy: "",
		Having: "",
		OrderBy: "",
		Limit: "",

	}
}

// func (obj *Switch) makeSwitchConnectionQueryComponent(IdSwitch string) database.QueryComponent {
// 	return database.QueryComponent {
// 		Tables: []string {"SWITCH_CONNECTION AS SC",
// 						"SERVER AS S",
// 						"STATUS_ROW AS SR"},
// 		Columns: []string {"SC.ID_SWITCH, SC.ID_CABLE_TYPE"},
// 		Selection: "ID_SERVER = ? AND " + 
// 					"SR.DESCRIPTION = 'available' " + 
// 					"AND S.ID_STATUS_ROW = SR.ID",
// 		SelectionArgs: []string {IdSwitch, "available"},
// 		GroupBy: "",
// 		Having: "",
// 		OrderBy: "",
// 		Limit: "",
// 	}
// }

func (obj *Switch) FetchIpAddrs() {
	component := obj.makeIpQueryComponent()
	rows, err := database.Query(component)
	defer rows.Close()

	if nil != err {
		panic (err)
	}

	obj.initIpAddrs(rows)
}

func (obj Switch) makeIpQueryComponent() database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"IP_SWITCH", "IP_NET"},
		Columns: []string {"IP_NET.VALUE", "IP_SWITCH.IP_HOST"},
		Selection: "IP_SWITCH.ID_SWITCH = ? AND IP_SWITCH.ID_IP_NET = IP_NET.ID",
		SelectionArgs: []string {obj.Id},
	}
}

func (obj *Switch) initIpAddrs(rows *sql.Rows) {
	var ipNet, ipHost string
	var ipAddr IpAddress

	for rows.Next() && (nil == rows.Scan(&ipNet, &ipHost)) {
		ipAddr.New(ipNet, ipHost)
		obj.IpAddrs = append(obj.IpAddrs, ipAddr)
	}
}

func (obj *Switch) Insert() error {
	obj.GenerateSwitchId()
	comp := obj.InsertComponent()
	return database.Insert(comp)
}

func (obj *Switch) GenerateSwitchId() {
	obj.Id = database.GeneratePrimaryKey(true,
						true, true, false, "SW", 10)
	for obj.Exists() {
		obj.Id = database.GeneratePrimaryKey(true,
						true, true, false, "SW", 10)
	}
}

func (obj *Switch) Exists() bool {
	comp := obj.ExistsQueryComp()
	rows, err := database.Query(comp)
	defer rows.Close()
	return (nil == err) && (rows.Next())
}

func (obj *Switch) ExistsQueryComp() qcomp {
	return qcomp {
		Tables: []string {"SWITCH"},
		Columns: []string {"ID"},
		Selection: "ID = ?",
		SelectionArgs: []string {obj.Id},
	}
}

func (obj *Switch) InsertComponent() icomp {
	return icomp {
		Table: "SWITCH",
		Columns: []string {"ID", 
						"ID_DC", 
						"ID_RACK", 
						"ID_U_START", 
						"ID_U_END", 
						"MAXIMUM_PORT",
					},
		Values: [][]string {
			[]string {
				obj.Id,
				obj.DC.Id,
				obj.Rack.Id,
				obj.UStart.Id,
				obj.UEnd.Id,
				obj.MaximumPort,
			},
		},
	}
}

func (obj *Switch) InsertIps() error {
	var comp icomp
	var err error

	for i := range obj.IpAddrs {
		comp = obj.InsertIpComp(obj.IpAddrs[i])
		err = database.Insert(comp)
		if nil != err {
			break
		}
	}

	return err
}

func (obj *Switch) InsertIpComp(ip IpAddress) icomp {
	return icomp {
		Table: "IP_SWITCH",
		Columns: []string {"ID_SWITCH", 
						"ID_IP_NET", 
						"IP_HOST",
					},
		Values: [][]string {
			[]string {obj.Id},
		},
	}
}
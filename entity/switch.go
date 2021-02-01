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

	if nil != err {
		return err
	}

	if rows.Next() {
		err = rows.Scan(&obj.Id,
					&obj.Name,
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
							"SW.NAME",
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
					"SW.ID_U_END = UEND.ID AND " + 
					"STATUS_ROW.DESCRIPTION = ? AND SW.ID_STATUS_ROW = STATUS_ROW.ID;",
		SelectionArgs: []string {IdSwitch, "available"},
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
		GroupBy: "",
		Having: "",
		OrderBy: "",
		Limit: "",
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
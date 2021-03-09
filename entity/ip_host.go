package entity

import (
	"CURD/database"

)

const (
	AVAILABLE = "AVAILABLE"
	USED = "USED"
)

type IpHost struct {
	IpNet
	Host string
	State string
}

func (ipHost *IpHost) Insert() error {
	comp := ipHost.insertComp()
	return database.Insert(comp)
}

func (ipHost *IpHost) insertComp() icomp {
	return icomp {
		Table: "IP_HOST",
		Columns: []string {"ID_NET", "HOST", "STATE"},
		Values: [][]string {
			[]string {
					ipHost.IpNet.Id, 
					ipHost.Host, 
					ipHost.State,
				},
		},
	}
}

func (ipHost *IpHost) Update() error {
	comp := ipHost.updateComp()
	return database.Update(comp)
}

func (ipHost *IpHost) updateComp() ucomp {
	return ucomp {
		Table: "IP_HOST",
		SetClause: "STATE = ?",
		Values: []string {ipHost.State},
		Selection: "ID_NET = ? AND HOST = ?",
		SelectionArgs: []string {ipHost.IpNet.Id, ipHost.Host},
	}
}

func QueryIpHost(comp qcomp) IpHost {
	rows, err := database.Query(comp)
	defer rows.Close()

	if nil != err {
		panic (err)
	}

	var ipHost IpHost
	if rows.Next() {
		err := rows.Scan(
			&ipHost.IpNet.Id,
			&ipHost.Host,
			&ipHost.State)

		if nil != err {
			panic (err)
		}
	}

	return ipHost
}

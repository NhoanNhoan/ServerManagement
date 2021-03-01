package entity

import (
	"CURD/database"
)

type IpAddress struct {
	IpNet
	IpHost string
}

func (obj *IpAddress) New(ipNetValue string, ipHost string) {
	obj.IpNet = IpNet {Value: ipNetValue}
	obj.IpHost = ipHost
}

func (ip *IpAddress) Insert(Id string) error {
	comp := ip.makeServerIComp(Id)
	return database.Insert(comp)
}

// Make insert component for IP_SERVER
func (ip *IpAddress) makeServerIComp(Id string) icomp {
	return icomp {
		Table: "IP_SERVER",
		Columns: []string {
			"ID_SERVER",
			"ID_IP_NET",
			"IP_HOST",
		},
		Values: [][]string {
			[]string {
				Id,
				ip.IpNet.Id,
				ip.IpHost,
			},
		},
	}
}

func (obj *IpAddress) String() string {
	return obj.IpNet.Value + obj.IpHost
}


func (obj *IpAddress) IsExistsServerIp(ServerId string) bool {
	comp := obj.makeCheckExistsServerIpQueryComponent(ServerId)
	rows, err := database.Query(comp)
	defer rows.Close()
	return (nil == err) && rows.Next()
}

func (obj *IpAddress) makeCheckExistsServerIpQueryComponent(ServerId string) database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"IP_SERVER"},
		Columns: []string {"ID_IP_NET"},
		Selection: "ID_SERVER = ? AND ID_IP_NET = ? AND IP_HOST = ?",
		SelectionArgs: []string {ServerId, obj.IpNet.Id, obj.IpHost},
	}
}

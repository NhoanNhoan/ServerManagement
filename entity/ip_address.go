package entity

import (
	"CURD/database"
	"strconv"
	"strings"
)

type IpAddress struct {
	Netmask int
	Octet1, Octet2, Octet3, Octet4 string
	State string
}

func (ip IpAddress) String() string {
	return strings.Join([]string {
		ip.Octet1,
		ip.Octet2,
		ip.Octet3,
		ip.Octet4,
	}, ".")
}

func (ip *IpAddress) Parse(content string) bool {
	octets := strings.Split(content, ".")
	if len(octets) < 4 {
		return false
	}

	splitNetmask := strings.Split(octets[3], "/")
	if len(splitNetmask) < 2 {
		return false
	}

	ip.set(octets...)
	return true
}

func (ip *IpAddress) set(octets ...string) {
	ip.Octet1 = octets[0]
	ip.Octet2 = octets[1]
	ip.Octet3 = octets[2]

	values := strings.Split(octets[3], "/")
	ip.Octet4 = values[0]
	ip.Netmask, _ = strconv.Atoi(values[1])
}

func (obj *IpAddress) New(ipNetValue string, ipHost string) {
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
			},
		},
	}
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
		SelectionArgs: []string {ServerId},
	}
}

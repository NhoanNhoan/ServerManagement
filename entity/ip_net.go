package entity

import (
	"database/sql"
	"math"
	"strconv"
	"strings"
	"net"

	"CURD/database"
)

type IpNet struct {
	Id string
	Value string
	Netmask int
}

func (ipNet IpNet) ToInstance(args ...string) Entity {
	Id, Value := args[0], args[1]
	return IpNet{Id: Id, Value: Value}
}

func GetIpNets() []IpNet {
	comp := database.MakeQueryAll([]string {"IP_NET"},
							[]string {"Id", "value"})
	ipNets := getEntities(comp, func (rows *sql.Rows) Entity {
		var id, des string
		err := rows.Scan(&id, &des)

		if nil != err {
			panic (err)
		}

		return IpNet {Id: id, Value: des}
	})

	return toIpNetSplice(ipNets)
}

func toIpNetSplice(entities []Entity) []IpNet {
	ipNets := make([]IpNet, len(entities))

	for i := range entities {
		ipNets[i] = entities[i].(IpNet)
	}

	return ipNets
}

func (net *IpNet) Insert() error {
	net.GenerateId()
	comp := net.makeInsertComp()
	err := database.Insert(comp)

	if nil != err {
		panic (err)
	}

	return net.insertIpHosts()
}

func (net *IpNet) GenerateId() {
	net.Id = database.GeneratePrimaryKey(true,
					true, true, false, "NET", 6)
	for net.Exists() {
		net.Id = database.GeneratePrimaryKey(true,
					true, true, false, "NET", 6)
	}
}

func (net *IpNet) Exists() bool {
	comp := net.existsComp()
	rows, err := database.Query(comp)
	defer rows.Close()
	return nil == err && rows.Next()
}

func (net *IpNet) existsComp() qcomp {
	return qcomp {
		Tables: []string {"IP_NET"},
		Columns: []string {"ID"},
		Selection: "ID = ?",
		SelectionArgs: []string {net.Id},
	}
}

func (net *IpNet) makeInsertComp() icomp {
	return icomp {
		Table: "IP_NET",
		Columns: []string {"ID", "VALUE", "NETMASK"},
		Values: [][]string {
					[]string {net.Id, 
							net.Value, 
							strconv.Itoa(net.Netmask),
							},
				},
	}
}

func (net *IpNet) insertIpHosts() error {
	var ipHost IpHost
	hosts := net.CalculateHostRange()
	var err error

	for _, host := range hosts {
		ipHost.IpNet = *net
		ipHost.Host = host
		ipHost.State = AVAILABLE
		err = ipHost.Insert()
		if nil != err {
			break
		}
	}

	return err
}

func (net *IpNet) CalculateHostRange() []string {
	netmask := strconv.Itoa(net.Netmask)
	ips, _, _ := Hosts(net.Value + "/" + netmask)
	return ips
}

func (net *IpNet) CalculateHosts() int {
	return int(math.Pow(2, float64(32 - net.Netmask))) - 2 // except gateway, net
}

func (net *IpNet) GetOctets() []int {
	octetStrings := strings.Split(net.Value, ".")
	octetIntegers := make([]int, len(octetStrings))

	for i := range octetIntegers {
		octetIntegers[i], _ = strconv.Atoi(octetStrings[i])
	}

	for len(octetIntegers) < 4 {
		octetIntegers = append(octetIntegers, 0)
	}

	return octetIntegers
}

func FetchIpHostArray(net IpNet) []IpHost {
	comp := queryIpHostArrayComp(net)
	rows, err := database.Query(comp)
	defer rows.Close()

	ipHost := IpHost{IpNet: net}
	arr := make([]IpHost, 0)
	for rows.Next() && nil == err {
		err = rows.Scan(&ipHost.Host, &ipHost.State)
		if nil == err {
			arr = append(arr, ipHost)
		}
	}

	return arr
}

func queryIpHostArrayComp(net IpNet) qcomp {
	return qcomp {
		Tables: []string {"IP_HOST"},
		Columns: []string {"HOST", "STATE"},
		Selection: "ID_NET = ? ORDER BY HOST ASC",
		SelectionArgs: []string {net.Id},
	}
}

func Hosts(cidr string) ([]string, int, error) {
    ip, ipnet, err := net.ParseCIDR(cidr)
    if err != nil {
        return nil, 0, err
    }

    var ips []string
    for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
        ips = append(ips, ip.String())
    }

    // remove network address and broadcast address
    lenIPs := len(ips)
    switch {
    case lenIPs < 2:
        return ips, lenIPs, nil

    default:
    return ips[1 : len(ips)-1], lenIPs - 2, nil
    }
}

func inc(ip net.IP) {
    for j := len(ip) - 1; j >= 0; j-- {
        ip[j]++
        if ip[j] > 0 {
            break
        }
    }
}

func FetchAllIpNets() []IpNet {
	rows := makeRowsAllIpNets()
	defer rows.Close()
	return parseIpNets(rows)
}

func makeRowsAllIpNets() *sql.Rows {
	comp := makeAllIpNetsComp()
	rows, err := database.Query(comp)
	if nil != err {
		panic (err)
	}
	return rows
}

func makeAllIpNetsComp() qcomp {
	return qcomp {
		Tables: []string {"IP_NET"},
		Columns: []string {"ID", "VALUE", "NETMASK"},
	}
}

func parseIpNets(rows *sql.Rows) []IpNet {
	ipNets := make([]IpNet, 0)
	for rows.Next() {
		ipNets = append(ipNets, parseIpNet(rows))
	}
	return ipNets
}

func parseIpNet(row *sql.Rows) IpNet {
	var net IpNet
	err := row.Scan(&net.Id, &net.Value, &net.Netmask)
	if nil != err {
		panic (err)
	}
	return net
}
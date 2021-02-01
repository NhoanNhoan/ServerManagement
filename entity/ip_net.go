package entity

import (
	"database/sql"
	"CURD/database"
)

type IpNet struct {
	Id string
	Value string
}

func (ipNet IpNet) ToInstance(args ...string) Entity {
	Id, Value := args[0], args[1]
	return IpNet{Id, Value}
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

		return IpNet {id, des}
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
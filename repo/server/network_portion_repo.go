package server

import (
	"CURD/database"
	"CURD/entity"
	"database/sql"
	"errors"
	"strconv"
	"strings"
)

type NetworkPortionRepo struct {
	SqliteRepo
}

func (repo NetworkPortionRepo) Fetch(comp qcomp,
	scan func(obj interface{}, row *sql.Rows) (interface{}, error)) ([]entity.NetworkPortion, error) {
	makeNetworkPortion := func() interface{} {return entity.NetworkPortion{}}
	entities, err := repo.SqliteRepo.Query(comp, makeNetworkPortion, scan)
	if nil != err {
		return nil, err
	}

	portions := make([]entity.NetworkPortion, len(entities))
	for i := range portions {
		portions[i] = entities[i].(entity.NetworkPortion)
	}

	return portions, err
}

func (repo NetworkPortionRepo) FetchAll() ([]entity.NetworkPortion, error) {
	comp := qcomp {
		Tables: []string {"NETWORK_PORTION"},
		Columns: []string {"ID", "VALUE", "NETMASK"},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		portion := obj.(entity.NetworkPortion)
		err := row.Scan(&portion.Id, &portion.Value, &portion.Netmask)
		return portion, err
	}

	return repo.Fetch(comp, scan)
}

func (repo NetworkPortionRepo) FetchHosts(portion entity.NetworkPortion) ([]entity.IpAddress, error) {
	comp := qcomp {
		Tables: []string {"NETWORK_PORTION AS N", "IP_ADDRESS AS I"},
		Columns: []string {"I.OCTET_1", "I.OCTET_2", "I.OCTET_3", "I.OCTET_4", "I.STATE"},
		Selection: "I.NETWORK_PORTION_ID = ?",
		SelectionArgs: []string {portion.Id},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		ip := obj.(entity.IpAddress)
		err := row.Scan(&ip.Octet1, &ip.Octet2, &ip.Octet3, &ip.Octet4, &ip.State)
		return ip, err
	}

	return IpRepo{}.Fetch(comp, scan)
}

func (repo NetworkPortionRepo) FetchHostsByState(portion entity.NetworkPortion,
	state string) ([]entity.IpAddress, error) {
	comp := qcomp {
		Tables: []string {"NETWORK_PORTION AS N", "IP_ADDRESS AS I"},
		Columns: []string {"I.OCTET_1", "I.OCTET_2", "I.OCTET_3", "I.OCTET_4", "I.STATE"},
		Selection: "I.NETWORK_PORTION_ID = ? AND i.STATE = ?",
		SelectionArgs: []string {portion.Id, state},
	}

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		ip := obj.(entity.IpAddress)
		err := row.Scan(&ip.Octet1, &ip.Octet2, &ip.Octet3, &ip.Octet4, &ip.State)
		return ip, err
	}

	return IpRepo{}.Fetch(comp, scan)
}

// Insert network portion
// it's will create ip addresses
// ex: network portion: 42.118.242.128/26 create ip range: 42.118.242.129 - 42.118.242.190
// so need insert into ip address
// the values network portion make it with available state
func (repo NetworkPortionRepo) Insert(portions ...entity.NetworkPortion) error {
	values := make([][]string, len(portions))
	for i := range portions {
		values[i] = []string {portions[i].Id, portions[i].Value, strconv.Itoa(portions[i].Netmask)}
	}

	comp := icomp {
		Table: "NETWORK_PORTION",
		Columns: []string {"ID", "VALUE", "NETMASK"},
		Values: values,
	}

	err := repo.SqliteRepo.Insert(comp)
	if nil != err {
		return err
	}

	for i := range portions {
		hosts, err := portions[i].GenerateHosts()
		if nil != err {
			return err
		}

		ipRepo := IpRepo{}
		if nil != ipRepo.Insert(hosts...) {
			return err
		}
	}

	return nil
}

func (repo NetworkPortionRepo) IsExists(portion entity.NetworkPortion) bool {
	comp := qcomp {
		Tables: []string {"NETWORK_PORTION"},
		Columns: []string {"ID"},
		Selection: "ID = ?",
		SelectionArgs: []string {portion.Id},
	}

	row, err := database.Query(comp)
	if nil != err {
		 return false
	}
	defer row.Close()

	return row.Next()
}

func (repo NetworkPortionRepo) GenerateId() string {
		Id := database.GeneratePrimaryKey(true,
		true, true,
		false, "NP", 6)

		for repo.IsExists(entity.NetworkPortion{Id: Id}) {
		Id = database.GeneratePrimaryKey(true,
		true, true,
		false, "NP", 6)
	}

		return Id
}

func (repo NetworkPortionRepo) Delete(NetworkPortionId string) error {
	comp := dcomp {
		Table: "NETWORK_PORTION",
		Selection: "ID = ?",
		SelectionArgs: []string {NetworkPortionId},
	}

	if err := repo.SqliteRepo.Delete(comp); nil != err {
		return err
	}

	return IpRepo{}.DeleteByNetworkPortionId(NetworkPortionId)
}

func StadardlizeNetworkPortion(raw string) (string, error) {
	octets := strings.Split(raw, ".")
	if len(octets) > 4 {
		return strings.Join(octets[:5], "."), nil
	}

	// Check whether octets is number
	for i := range octets {
		if "" == octets[i] {
			octets[i] = "0"
		}

		if num, err := strconv.Atoi(octets[i]); nil != err || num < 0 || num > 255 {
			return "", errors.New(raw + " is not like format of ip address")
		}
	}

	// Insert 0 if octet is empty
	for len(octets) < 4 {
		octets = append(octets, "0")
	}

	return strings.Join(octets, "."), nil
}
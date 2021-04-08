package server

import (
	"CURD/database"
	"CURD/entity"
	"database/sql"
	"strconv"
)

type IpRepo struct {
	repo SqliteRepo
}

func (ipRepo IpRepo) Fetch(comp qcomp,
	scan func(obj interface{}, rows *sql.Rows) (interface{}, error)) ([]entity.IpAddress, error) {
	entities, err := ipRepo.repo.Query(comp, func() interface{} {return entity.IpAddress{}}, scan)
	if nil != err {
		return nil, err
	}

	ips := make([]entity.IpAddress, len(entities))
	for i := range ips {
		ips[i] = entities[i].(entity.IpAddress)
	}

	return ips, nil
}

func (ipRepo IpRepo) FetchById(id string) (interface{}, error) {
	return nil, nil
}

func (ipRepo IpRepo) IsExists(id string) bool {
	return false
}

func (ipRepo IpRepo) GenerateId() string {
	return ""
}

func (ipRepo IpRepo) Insert(ipArr... entity.IpAddress) error {
	values := make([][]string, len(ipArr))
	for i := range values {
		values[i] = []string {ipArr[i].Octet1,
					ipArr[i].Octet2,
					ipArr[i].Octet3,
					ipArr[i].Octet4,
					strconv.Itoa(ipArr[i].Netmask),
					ipArr[i].NetworkPortion.Id,
					ipArr[i].State}
	}

	comp := icomp {
		Table: "IP_ADDRESS",
		Columns: []string {"OCTET_1", "OCTET_2",
			"OCTET_3", "OCTET_4",
			"NETMASK", "NETWORK_PORTION_ID",
			"STATE"},
		Values: values,
	}

	return ipRepo.repo.Insert(comp)
}

func (ipRepo IpRepo) Update(ipArr... entity.IpAddress) error {
	return nil
}

func (repo IpRepo) UpdateState(ip entity.IpAddress, state string) error {
	comp := ucomp {
		Table: "IP_ADDRESS",
		SetClause: "STATE = ?",
		Values: []string {state},
		Selection: "OCTET_1 = ? AND OCTET_2 = ? AND OCTET_3 = ? AND OCTET_4 = ? AND NETMASK = ?",
		SelectionArgs: []string {ip.Octet1, ip.Octet2, ip.Octet3, ip.Octet4, strconv.Itoa(ip.Netmask)},
	}

	return repo.repo.Update(comp)
}

func (ipRepo IpRepo) Delete(ipArr... entity.IpAddress) error {
	return nil
}

func (repo IpRepo) DeleteByNetworkPortionId(NetworkPortionId string) error {
	comp := dcomp {
		Table: "IP_ADDRESS",
		Selection: "NETWORK_PORTION_ID = ?",
		SelectionArgs: []string {NetworkPortionId},
	}

	return repo.repo.Delete(comp)
}

func (repo IpRepo) FetchState(ip entity.IpAddress) (string, error) {
	comp := database.QueryComponent{
		Tables: []string {"IP_ADDRESS"},
		Columns: []string {"STATE"},
		Selection: "OCTET_1 = ? AND OCTET_2 = ? AND OCTET_3 = ? AND OCTET_4 = ? AND NETMASK = ?",
		SelectionArgs: []string {ip.Octet1, ip.Octet2, ip.Octet3, ip.Octet4, strconv.Itoa(ip.Netmask)},
	}

	row, err := database.Query(comp)
	if nil != err {
		return "", err
	}
	defer row.Close()

	var state string
	if row.Next() {
		err = row.Scan(&state)
	}

	return state, err
}

func (repo IpRepo) CountByState(NetId string) (map[string]int, error) {
	comp := qcomp{
		Tables:  []string{"IP_ADDRESS AS I1"},
		Columns: []string{"STATE", "COUNT(SELECT COUNT(I2.STATE) FROM IP_ADDRESS AS I2 WHERE I1.STATE = I2.STATE"},
	}

	rows, err := database.Query(comp)
	if nil != err {
		return nil, err
	}
	defer rows.Close()

	stateCount := map[string]int{}
	for rows.Next() {
		var state string
		var count int
		err = rows.Scan(&state, &count)
		if nil != err {
			return nil, err
		}

		stateCount[state] = count
	}

	return stateCount, err
}

func (repo IpRepo) CountStates(NetId string, State string) (int, error) {
	comp := qcomp {
		Tables: []string {"IP_ADDRESS"},
		Columns: []string {"COUNT(STATE)"},
		Selection: "STATE = ? AND NETWORK_PORTION_ID = ?",
		SelectionArgs: []string {State, NetId},
	}

	var err error
	var rows *sql.Rows
	if rows, err = database.Query(comp); nil == err {
		defer rows.Close()
		if rows.Next() {
			var count int
			err := rows.Scan(&count)
			return count, err
		}

		return 0, nil
	}

	return 0, err
}
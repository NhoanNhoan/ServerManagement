package server

import (
	"CURD/database"
	"CURD/entity"
	"database/sql"
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
	return nil
}

func (ipRepo IpRepo) Update(ipArr... entity.IpAddress) error {
	return nil
}

func (ipRepo IpRepo) Delete(ipArr... entity.IpAddress) error {
	return nil
}

func (repo IpRepo) FetchState(ip entity.IpAddress) (string, error) {
	comp := database.QueryComponent{
		Tables: []string {"IP_ADDRESS"},
		Columns: []string {"STATE"},
		Selection: "OCTET_1 = ? AND OCTET_2 = ? AND OCTET_3 = ? AND OCTET_4 = ? AND NETMASK = ?",
		SelectionArgs: []string {ip.Octet1, ip.Octet2, ip.Octet3, ip.Octet4, ip.Octet4},
	}

	row, err := database.Query(comp)
	if nil != err {
		return "", err
	}
	defer row.Close()

	var state string
	err = row.Scan(&state)

	return state, err
}
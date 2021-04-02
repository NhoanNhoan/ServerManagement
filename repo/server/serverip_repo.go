package server

import (
	"CURD/database"
	"CURD/entity"
	"database/sql"
	"strconv"
	"strings"
)

type ServerIpRepo struct {
	SqliteRepo
}

func (repo ServerIpRepo) Insert(comp icomp) error {
	return repo.SqliteRepo.Insert(comp)
}

func (repo ServerIpRepo) MakeIcomp(ServerId string, list ...entity.IpAddress) icomp {
	return icomp{
		Table: "IP_ADDRESS",
		Columns: []string{
			"ServerId",
			"Octet_1",
			"Octet_2",
			"Octet_3",
			"Octet_4",
			"Netmask",
		},
		Values: repo.makeValues(ServerId, list...),
	}
}

func (repo ServerIpRepo) makeValues(ServerId string, list ...entity.IpAddress) [][]string {
	values := make([][]string, len(list))
	for i := range list {
		values = append(values,
			[]string{ServerId,
				list[i].Octet1,
				list[i].Octet2,
				list[i].Octet3,
				list[i].Octet4,
				strconv.Itoa(list[i].Netmask),
			})
	}
	return values
}

func (repo ServerIpRepo) Delete(serverId string, listIp ... entity.IpAddress) error {
	if len(listIp) == 0 {
		return nil
	}

	comp := dcomp{
		Table: "SERVER_IP",
		Selection: strings.Join([]string {"SERVER_ID = ?",
			"OCTET_1 = ?", "OCTET_2 = ?",
			"OCTET_3 = ?", "OCTET_4 = ?", "NETMASK = ?"},
			" AND "),
	}

	for _, ip := range listIp {
		comp.SelectionArgs = []string {serverId,
			ip.Octet1, ip.Octet2,
			ip.Octet3, ip.Octet4, string(ip.Netmask)}
		if err := repo.SqliteRepo.Delete(comp); nil != err {
			return err
		}
	}

	return nil
}

func (s ServerIpRepo) FetchServerIpAddrs(ServerId string) ([]entity.IpAddress, error) {
	comp := qcomp{
		Tables: []string {"SERVER_IP", "IP_TYPE"},
		Columns: []string {"OCTET_1", "OCTET_2", "OCTET_3", "OCTET_4", "NETMASK"},
		Selection: "SERVER_ID = ? AND IP_TYPE.DES = ? AND SERVER_IP.IP_TYPE_ID = IP_TYPE.ID",
		SelectionArgs: []string {ServerId, "NORMAL"},
	}
	return IpRepo{}.Fetch(comp, scanIp)
}

func (s ServerIpRepo) FetchRedfishIp(ServerId string) ([]entity.IpAddress, error) {
	comp := qcomp{
		Tables: []string {"SERVER_IP", "IP_TYPE"},
		Columns: []string {"OCTET_1", "OCTET_2", "OCTET_3", "OCTET_4", "NETMASK"},
		Selection: "SERVER_ID = ? AND IP_TYPE.DES = ? AND SERVER_IP.IP_TYPE_ID = IP_TYPE.ID",
		SelectionArgs: []string {ServerId, "REDFISH"},
	}
	return IpRepo{}.Fetch(comp, scanIp)
}

func (s ServerIpRepo) fetchListIpByServerIdComp(ServerId string) qcomp {
	return qcomp{
		Tables: []string {"SERVER_IP"},
		Columns: []string {"OCTET_1", "OCTET_2", "OCTET_3", "OCTET_4", "NETMASK"},
		Selection: "SERVER_ID = ?",
		SelectionArgs: []string {ServerId},
	}
}

func (s ServerIpRepo) fetchListRedfishIpComp(ServerId string) qcomp {
	comp := s.fetchListIpByServerIdComp(ServerId)
	comp.Selection = "SERVER_ID = ? AND IP_TYPE.DES = ? AND SERVER.IP_TYPE_ID = ?"
	comp.SelectionArgs = []string {ServerId, "REDFISH", "IP_TYPE.ID"}
	return comp
}

func scanIp(obj interface{}, row *sql.Rows) (interface{}, error) {
	ip := obj.(entity.IpAddress)
	err := row.Scan(&ip.Octet1, &ip.Octet2, &ip.Octet3, &ip.Octet4, &ip.Netmask)
	return ip, err
}

func (s ServerIpRepo) FetchServerIdByIp(ip entity.IpAddress) (string, error) {
	comp := qcomp{
		Tables: []string {"SERVER_IP"},
		Columns: []string {"SERVER_ID"},
		Selection: "OCTET_1 = ? AND OCTET_2 = ? AND OCTET_3 = ? AND OCTET_4 = ? AND NETMASK=?",
		SelectionArgs: []string {ip.Octet1, ip.Octet2, ip.Octet3, ip.Octet4, strconv.Itoa(ip.Netmask)},
	}

	row, err := database.Query(comp)
	if nil != err {
		return "", err
	}
	defer row.Close()

	if row.Next() {
		var serverId string
		err = row.Scan(&serverId)
		return serverId, err
	}

	return "", nil
}
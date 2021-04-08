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

// When insert a new ip address into server_ip table
// the state of ip will change used state
func (repo ServerIpRepo) InsertNormalIpAddresses(ServerId string, listIp ...entity.IpAddress) error {
	comp := icomp{
		Table: "SERVER_IP",
		Columns: []string{
			"SERVER_ID",
			"Octet_1",
			"Octet_2",
			"Octet_3",
			"Octet_4",
			"Netmask",
			"IP_TYPE_ID",
		},
		Values: repo.makeNormalValues(ServerId, listIp...),
	}

	ipRepo := IpRepo{}
	for i := range listIp {
		if err := ipRepo.UpdateState(listIp[i], "used"); nil != err {
			return err
		}
	}
	return repo.Insert(comp)
}

func (repo ServerIpRepo) InsertRedfishIpAddresses(ServerId string, list ...entity.IpAddress) error {
	comp := icomp{
		Table: "SERVER_IP",
		Columns: []string{
			"SERVER_ID",
			"Octet_1",
			"Octet_2",
			"Octet_3",
			"Octet_4",
			"Netmask",
			"IP_TYPE_ID",
		},
		Values: repo.makeRedfishValues(ServerId, list...),
	}

	return repo.Insert(comp)
}

func (repo ServerIpRepo) MakeIcomp(ServerId string, list ...entity.IpAddress) icomp {
	return icomp{
		Table: "SERVER_IP",
		Columns: []string{
			"SERVER_ID",
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
		values[i] = []string{ServerId,
			list[i].Octet1,
			list[i].Octet2,
			list[i].Octet3,
			list[i].Octet4,
			strconv.Itoa(list[i].Netmask),
		}
	}
	return values
}

func (repo ServerIpRepo) makeRedfishValues(ServerId string, list ...entity.IpAddress) [][]string {
	values := make([][]string, len(list))
	for i := range list {
		values[i] = []string{ServerId,
			list[i].Octet1,
			list[i].Octet2,
			list[i].Octet3,
			list[i].Octet4,
			strconv.Itoa(list[i].Netmask),
			"2",
		}
	}
	return values
}

func (repo ServerIpRepo) makeNormalValues(ServerId string, listIp ...entity.IpAddress) [][]string {
	values := make([][]string, len(listIp))
	for i := range listIp {
		values[i] = []string{ServerId,
			listIp[i].Octet1,
			listIp[i].Octet2,
			listIp[i].Octet3,
			listIp[i].Octet4,
			strconv.Itoa(listIp[i].Netmask),
			"1",
		}
	}
	return values
}

// Delete a ip address of server then set state of ip is available
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

		ipRepo := IpRepo{}
		if err := ipRepo.UpdateState(ip, "available"); nil != err {
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

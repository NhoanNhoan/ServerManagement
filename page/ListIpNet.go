package page

import (
	"CURD/database"
	"CURD/entity"
	"database/sql"
)

type IpNetItem struct {
	entity.IpNet
	AvailableIpNets int
	UsedIpNets int
}

func (item *IpNetItem) parseByRow(row *sql.Rows) error {
	return row.Scan(
		&item.IpNet.Id,
		&item.IpNet.Value,
		&item.IpNet.Netmask,
		&item.AvailableIpNets,
		&item.UsedIpNets)
}

func (item *IpNetItem) fetchAvailableAndUsedIpNets() error {
	row := item.fetchAvailableAndUsedIpNetsRow()
	defer row.Close()

	var err error
	if row.Next() {
		err = row.Scan(&item.AvailableIpNets,
			&item.UsedIpNets)
	}

	return err
}

func (item *IpNetItem) fetchAvailableAndUsedIpNetsRow() *sql.Rows {
	comp := item.fetchAvailableAndUsedIpNetsComp()
	row, err := database.Query(comp)

	if nil != err {
		panic (err)
	}

	return row
}

func (item *IpNetItem) fetchAvailableAndUsedIpNetsComp() database.QueryComponent {
	return database.QueryComponent{
		Tables: []string {"IP_HOST"},
		Columns: []string {"COUNT(CASE WHEN STATE='AVAILABLE' THEN 1 ELSE NULL END)",
			"COUNT(CASE WHEN STATE='USED' THEN 1 ELSE NULL END)"},
		Selection: "ID_NET = ?",
		SelectionArgs: []string {item.IpNet.Id},
	}
}

type ListIpNet struct {
	Items []IpNetItem
}

func (obj *ListIpNet) New() {
	obj.FetchAllIpNetItems()
}

func (obj *ListIpNet) FetchAllIpNetItems() error {
	nets := entity.FetchAllIpNets()
	obj.Items = make([]IpNetItem, len(nets))

	var err error
	for i := range nets {
		obj.Items[i].IpNet = nets[i]
		err = obj.Items[i].fetchAvailableAndUsedIpNets()
		if nil != err {
			break
		}
	}

	return err
}

func (obj *ListIpNet) makeRowsQuery() *sql.Rows {
	comp := obj.makeFetchAllIpNetItemComp()
	rows, err := database.Query(comp)

	if nil != err {
		panic (err)
	}

	return rows
}

func (obj *ListIpNet) makeFetchAllIpNetItemComp() database.QueryComponent {
	return database.QueryComponent{
		Tables: []string {"IP_NET", "IP_HOST"},
		Columns: []string {"ID", "VALUE", "NETMASK",
							"COUNT(CASE WHEN STATE='AVAILABLE' THEN 1 ELSE NULL END)",
							"COUNT(CASE WHEN STATE='USED' THEN 1 ELSE NULL END)"},
		Selection: "ID = ID_NET",
	}
}

func (obj *ListIpNet) fetchItemsByRows(rows *sql.Rows) error {
	var err error
	var item IpNetItem

	for rows.Next() && nil == err {
		if err = item.parseByRow(rows); nil == err {
			obj.Items = append(obj.Items, item)
		}
	}

	return err
}
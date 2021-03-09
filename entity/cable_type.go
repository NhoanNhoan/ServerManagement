package entity

import (
	"CURD/database"
)

type CableType struct {
	Id string
	Name string
	SignPort string
}

func QueryAllCabs() []CableType {
	comp := queryAllCabsComp()
	rows, err := database.Query(comp)
	defer rows.Close()

	if nil != err {
		panic (err)
	}

	var cable CableType
	cabs := make([]CableType, 0)
	for rows.Next() && nil == err {
		err = rows.Scan(&cable.Id, &cable.Name, &cable.SignPort)
		cabs = append(cabs, cable)
	}

	if nil != err {
		panic (err)
	}

	return cabs
}


func queryAllCabsComp() qcomp {
	return qcomp {
		Tables: []string {"CABLE_TYPE"},
		Columns: []string {"ID", "NAME", "SIGN_PORT"},
	}
}
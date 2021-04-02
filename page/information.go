package page

import (
	"CURD/entity"
)

type Information struct {
	entity.Server
	ConnectedSwitches []SwitchInfo
	Tagged            []entity.Tag
	Untagged          []entity.Tag
}

//func (s *Information) FetchServerByIp(Net, Host string) {
//	comp := s.makeQueryServerByIpComp(Net, Host)
//	var err error
//	s.Server, err = entity.FetchServer(comp)
//	s.Server.FetchServices()
//	if nil != err {
//		panic (err)
//	}
//}

//func (info *SwitchInfo) ParseFromRow(row *sql.Rows) error {
//	return row.Scan(&info.Switch.Id,
//				&info.Switch.DC.Description,
//				&info.Switch.Rack.Description,
//				&info.Switch.UStart.Description,
//				&info.Switch.UEnd.Description,
//				&info.Switch.MaximumPort,
//				&info.SwitchConnection.Port,
//				&info.CableType.Name,
//				&info.CableType.SignPort)
//}
//
//func (s *Information) FetchServerByIp(Net, Host string) {
//	comp := s.makeQueryServerByIpComp(Net, Host)
//	var err error
//	s.Server, err = entity.FetchServer(comp)
//	s.Server.FetchServices()
//	if nil != err {
//		panic (err)
//	}
//}
//
//func (s *Information) makeQueryServerByIpComp(Net string, Host string) database.QueryComponent {
//	return database.QueryComponent {
//		Tables: []string {"SERVER AS S",
//			"IP_NET",
//			"IP_SERVER",
//			"DC AS D",
//			"RACK AS R",
//			"RACK_UNIT AS USTART",
//			"RACK_UNIT AS UEND",
//			"PORT_TYPE AS PT",
//			"SERVER_STATUS AS SS",
//			"STATUS_ROW AS SR",
//		},
//		Columns: []string {
//			"S.ID",
//			"D.ID",
//			"D.DESCRIPTION",
//			"R.ID",
//			"R.DESCRIPTION",
//			"USTART.ID",
//			"USTART.DESCRIPTION",
//			"UEND.ID",
//			"UEND.DESCRIPTION",
//			"S.SSD",
//			"S.HDD",
//			"S.MAKER",
//			"PT.ID",
//			"PT.DESCRIPTION",
//			"S.SERIAL_NUMBER",
//			"SS.ID",
//			"SS.DESCRIPTION",
//		},
//		Selection: "S.ID = IP_SERVER.ID_SERVER AND " +
//			"IP_SERVER.ID_IP_NET = IP_NET.ID AND " +
//			"? = IP_NET.Id AND ? = IP_SERVER.IP_HOST AND " +
//			"S.ID_DC = D.ID AND " +
//			"S.ID_RACK = R.ID AND " +
//			"S.ID_U_START = USTART.ID AND " +
//			"S.ID_U_END = UEND.ID AND " +
//			"SR.DESCRIPTION = ? AND S.ID_STATUS_ROW = SR.ID AND " +
//			"S.ID_PORT_TYPE = PT.ID AND " +
//			"S.ID_SERVER_STATUS = SS.ID",
//		SelectionArgs: []string {Net, Host, "available"},
//	}
//}
//
//func (obj *Information) Prepare(Net, Host string) {
//	obj.FetchServerByIp(Net, Host)
//	obj.FetchTagged()
//	obj.FetchUntagged()
//	obj.Server.FetchIpAddrs()
//	obj.Server.FetchEvents()
//	obj.Server.FetchServices()
//	err := obj.initConnectedSwitches(obj.Server.Id)
//
//	if nil != err {
//		panic (err)
//	}
//}
//
//func (obj *Information) New(IdServer string) {
//	obj.initServer(IdServer)
//	obj.FetchTagged()
//	obj.FetchUntagged()
//	err := obj.initConnectedSwitches(IdServer)
//	if nil != err {
//		panic (err)
//	}
//}
//
//func (obj *Information) initServer(IdServer string) {
//	err := obj.Server.New(IdServer)
//	obj.Server.FetchIpAddrs()
//	obj.Server.FetchServices()
//	obj.Server.FetchEvents()
//
//	if nil != err {
//		panic (err)
//	}
//}
//
//func (obj *Information) initConnectedSwitches(ServerId string) error {
//	comp := obj.queryConnectedSwitchesComp(ServerId)
//	rows, err := database.Query(comp)
//	defer rows.Close()
//
//	var switchInfo SwitchInfo
//	for rows.Next() && nil == err {
//		err = switchInfo.ParseFromRow(rows)
//		switchInfo.Switch.FetchIpAddrs()
//		obj.ConnectedSwitches = append(obj.ConnectedSwitches,
//									switchInfo)
//	}
//
//	return err
//}
//
//func (obj *Information) queryConnectedSwitchesComp(ServerId string) database.QueryComponent {
//	return database.QueryComponent {
//		Tables: []string {
//			"SWITCH AS SW",
//			"SWITCH_CONNECTION AS SC",
//			"DC AS D",
//			"RACK AS R",
//			"RACK_UNIT AS USTART",
//			"RACK_UNIT AS UEND",
//			"CABLE_TYPE AS CT",
//		},
//
//		Columns: []string {
//			"SW.ID",
//			"D.DESCRIPTION",
//			"R.DESCRIPTION",
//			"USTART.DESCRIPTION",
//			"UEND.DESCRIPTION",
//			"SW.MAXIMUM_PORT",
//			"SC.PORT",
//			"CT.NAME",
//			"CT.SIGN_PORT",
//		},
//
//		Selection: "SC.ID_SERVER = ? AND " +
//				"SC.ID_SWITCH = SW.ID AND " +
//				"SW.ID_DC = D.ID AND " +
//				"SW.ID_RACK = R.ID AND " +
//				"SW.ID_U_START = USTART.ID AND " +
//				"SW.ID_U_END = UEND.ID AND " +
//				"SC.ID_CABLE_TYPE = CT.ID",
//
//		SelectionArgs: []string {ServerId},
//	}
//}
//
//func (obj *Information) FetchTagged() {
//	comp := obj.makeTaggedQueryComponent()
//	obj.Tagged = entity.FetchTags(comp)
//}
//
//func (obj *Information) FetchUntagged() {
//	comp := obj.makeUntaggedQueryComponent()
//	obj.Untagged = entity.FetchTags(comp)
//}
//
//func (obj *Information) makeTaggedQueryComponent() database.QueryComponent {
//	return database.QueryComponent {
//		Tables: []string {"TAG", "SERVER_TAG"},
//		Columns: []string {"TAG.TAGID", "TAG.TITLE"},
//		Selection: "SERVER_TAG.SERVERID = ? AND " +
//				"SERVER_TAG.TAGID = TAG.TAGID",
//		SelectionArgs: []string {obj.Server.Id},
//	}
//}
//
//func (obj *Information) makeUntaggedQueryComponent() database.QueryComponent {
//	selection := "TAGID NOT IN (SELECT S.TAGID " +
//				"FROM SERVER_TAG AS S " +
//				"WHERE S.SERVERID = ?)"
//	return database.QueryComponent {
//		Tables: []string {"TAG"},
//		Columns: []string {"TAGID", "TITLE"},
//		Selection: selection,
//		SelectionArgs: []string {obj.Server.Id},
//	}
//}

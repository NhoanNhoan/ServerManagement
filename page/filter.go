package page

import (
	"CURD/entity"
)

type Filter struct {
	entity.Tag
	Tags []entity.Tag
	Servers []entity.Server
}

func (f *Filter) New(tag entity.Tag) {
	//f.SetTag(tag)
	//err := f.initializeServers(f.Tag.TagId)
	//if nil != err {
	//	panic (err)
	//}
	//
	//f.initializeTags()
}
//
//func (f *Filter) SetTag(tag entity.Tag) {
//	f.Tag = tag
//}
//
//func (f *Filter) initializeServers(TagId string) error {
//	comp := f.makeQueryServersByTagComponent(TagId)
//	var err error
//	err, f.Servers = entity.FetchServers(comp)
//	return err
//}
//
//func (f *Filter) initializeTags() {
//	f.Tags = entity.FetchAllTags()
//}
//
//func (f *Filter) fetchServersIpAddr() {
//	for i := range f.Servers {
//		f.Servers[i].FetchIpAddrs()
//	}
//}
//
//func (f *Filter) makeQueryServersByTagComponent(TagId string) database.QueryComponent {
//	return database.QueryComponent {
//		Tables: []string {
//				"SERVER AS S",
//				"DC AS D",
//				"RACK AS R",
//				"RACK_UNIT AS USTART",
//				"RACK_UNIT AS UEND",
//				"PORT_TYPE AS PT",
//				"SERVER_STATUS AS SS",
//				"STATUS_ROW AS SR",
//				"SERVER_TAG AS ST",
//			},
//
//		Columns: []string {
//				"D.ID",
//				"D.DESCRIPTION",
//				"R.ID",
//				"R.DESCRIPTION",
//				"USTART.ID",
//				"USTART.DESCRIPTION",
//				"UEND.ID",
//				"UEND.DESCRIPTION",
//				"S.MAKER",
//				"PT.ID",
//				"PT.DESCRIPTION",
//				"S.SERIAL_NUMBER",
//				"SS.ID",
//				"SS.DESCRIPTION",
//			},
//
//		Selection:	"ST.TAGID = ? AND " +
//					"ST.SERVERID = S.ID AND " +
//					"S.ID_DC = D.ID AND " +
//					"S.ID_RACK = R.ID AND " +
//					"S.ID_U_START = USTART.ID AND " +
//					"S.ID_U_END = UEND.ID AND " +
//					"SR.DESCRIPTION = ? AND S.ID_STATUS_ROW = SR.ID AND " +
//					"S.ID_PORT_TYPE = PT.ID AND " +
//					"S.ID_SERVER_STATUS = SS.ID",
//
//		SelectionArgs: []string {TagId, "available"},
//	}
//}
//
//func (f *Filter) SearchServersByMultiTags(tags []string) error {
//	comp := f.makeQueryServerByMultiTagsComp(tags)
//	rows, err := database.Query(comp)
//	defer rows.Close()
//
//	for rows.Next() {
//		var server entity.Server
//		err = rows.Scan(&server.Id)
//		server.FetchIpAddrs()
//		f.Servers = append(f.Servers, server)
//	}
//
//	for i := range f.Servers {
//		f.Servers[i].New(f.Servers[i].Id)
//	}
//
//	return err
//}
//
//func (f *Filter) makeQueryServerByMultiTagsComp(tags []string) database.QueryComponent {
//	selection := ""
//	if len(tags) > 1 {
//		selection = f.makeSelection(tags)
//	} else {
//		selection = "TAG.TITLE = ? AND SERVER_TAG.TAGID = TAG.TAGID AND SERVER.ID = SERVER_TAG.SERVERID"
//	}
//
//	comp := database.QueryComponent {
//		Tables: []string {"SERVER", "SERVER_TAG", "TAG"},
//		Columns: []string {"SERVER.ID"},
//		Selection: selection,
//		SelectionArgs: tags,
//	}
//
//	fmt.Println(database.GetQueryStatement(comp))
//	return comp
//
//}
//
//func (f *Filter) makeSelection(tags []string) string {
//	selections := make([]string, len(tags) - 1)
//
//	for i := range selections {
//		clause := "SELECT SERVER.ID FROM SERVER, SERVER_TAG, TAG WHERE TAG.TITLE = ? AND SERVER_TAG.TAGID = TAG.TAGID AND SERVER.ID = SERVER_TAG.SERVERID "
//		selections[i] = clause
//	}
//
//	return "TAG.TITLE = ? AND SERVER_TAG.TAGID = TAG.TAGID AND SERVER.ID = SERVER_TAG.SERVERID intersect " + strings.Join(selections, " intersect ")
//}

package page

import (
	"CURD/database"
	"CURD/entity"
	"CURD/repo/server"
	"database/sql"
)

type Servers struct {
	entity.DataCenter
	Items []entity.Server
	Tags []entity.Tag
}

func (s *Servers) FetchServerByDCId(DCId string) error {
	comp := s.makeQueryServersComp(DCId)
	return s.fetchServers(comp)
}

func (s Servers) makeQueryServersComp(DCId string) database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"SERVER", 
						"RACK", 
						"RACK_UNIT AS USTART", 
						"RACK_UNIT AS UEND", 
						"PORT_TYPE",
						"SERVER_STATUS",
						"SERVE"},
		Columns: []string {"SERVER.ID",
						"RACK.Description",
						"USTART.Description",
						"UEND.Description",
						"SSD",
						"HDD",
						"PORT_TYPE.Description",
						"SERVER_STATUS.Description",
						"SERVER.MAKER",
						"SERVE.DESCRIPTION"},
		Selection: "SERVER.ID_DC = ? AND " +
				"SERVER.ID_RACK = RACK.ID AND " + 
				"SERVER.ID_U_START = USTART.ID AND " + 
				"SERVER.ID_U_END = UEND.ID AND " + 
				"SERVER.ID_PORT_TYPE = PORT_TYPE.ID AND " + 
				"SERVER.ID_SERVER_STATUS = SERVER_STATUS.ID AND " +
				"SERVER.ID_SERVE = SERVE.ID",
		SelectionArgs: []string {DCId},
	}
}

func (s *Servers) FetchServerByTagId(tagId string) error {
	comp := s.makeQueryCompByTagId(tagId, s.DataCenter.Id)
	return s.fetchServers(comp)
}

func (s Servers) makeQueryCompByTagId(tagId string, dcId string) database.QueryComponent {
	return database.QueryComponent {
			Tables: []string {"SERVER", 
							"RACK", 
							"RACK_UNIT AS USTART", 
							"RACK_UNIT AS UEND", 
							"PORT_TYPE",
							"SERVER_STATUS",
							"STATUS_ROW",
							"SERVER_TAG AS ST"},
			Columns: []string {"SERVER.ID",
							"RACK.Description",
							"USTART.Description",
							"UEND.Description",
							"SERVER.SSD",
							"SERVER.HDD",
							"PORT_TYPE.Description",
							"SERVER.SERIAL_NUMBER",
							"SERVER_STATUS.Description",
							"SERVER.MAKER",
							"SERVER.REDFISH_IP"},
			Selection: "SERVER.ID_DC = ? AND " +
					"SERVER.ID_RACK = RACK.ID AND " + 
					"SERVER.ID_U_START = USTART.ID AND " + 
					"SERVER.ID_U_END = UEND.ID AND " + 
					"SERVER.ID_PORT_TYPE = PORT_TYPE.ID AND " + 
					"SERVER.ID_SERVER_STATUS = SERVER_STATUS.ID AND " + 
					"STATUS_ROW.DESCRIPTION = ? AND SERVER.ID_STATUS_ROW = STATUS_ROW.ID AND " +
					"ST.TAGID = ? AND SERVER.ID = ST.SERVERID",
			SelectionArgs: []string {dcId, "available", tagId},
			GroupBy: "",
			Having: "",
			OrderBy: "",
			Limit: "",
		}
}

func (s *Servers) fetchServers(comp database.QueryComponent) error {
	scanServer := func (e interface{}, rows *sql.Rows) (interface{}, error) {
		server := e.(entity.Server)
		err := rows.Scan(
			&server.Id,
			&server.Rack.Description,
			&server.UStart.Description,
			&server.UEnd.Description,
			&server.SSD,
			&server.HDD,
			&server.PortType.Description,
			&server.ServerStatus.Description,
			&server.Maker,
			&server.ServeCustomer.Description)
		return server, err
	}

	var err error
	s.Items, err = server.ServerRepo{}.Fetch(comp, scanServer)

	if nil == err {
		err = s.FetchServerIpsItems()
	}

	if nil == err {
		err = s.fetchRedfishIp()
	}

	if nil == err {
		err = s.initTags()
	}

	return err
}

func (s *Servers) FetchServerIpsItems() (err error) {
	for i := range s.Items {
		s.Items[i].IpAddrs, err = server.ServerIpRepo{}.FetchServerIpAddrs(s.Items[i].Id)
		if nil != err {
			return
		}
	}

	return
}

func (s *Servers) fetchRedfishIp() (err error) {
	for i := range s.Items {
		ips, err := server.ServerIpRepo{}.FetchRedfishIp(s.Items[i].Id)
		if nil != err {
			return err
		}

		if len(ips) > 0 {
			redfishIp := ips[0]
			s.Items[i].RedfishIp = redfishIp.String()
		}
	}

	return nil
}

func (s *Servers) initTags() error {
	comp := database.QueryComponent{
		Tables: []string {"TAG"},
		Columns: []string {"TAGID", "TITLE"},
	}

	scanTag := func (obj interface{}, rows *sql.Rows) (interface{}, error) {
		tag := obj.(entity.Tag)
		err := rows.Scan(&tag.TagId, &tag.Title)
		return tag, err
	}

	var err error
	s.Tags, err = server.TagRepo{}.Fetch(comp, scanTag)
	return err
}
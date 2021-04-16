package page

import (
	"CURD/database"
	"CURD/entity"
	"CURD/repo/server"
	"CURD/repo/server/hardware_repo"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	InfoLogger *log.Logger
)

type Filter struct {
	entity.Tag
	ClusterServers []string
	Tags []entity.Tag
	Servers []entity.Server
	Stats FilterStats
}

func (f *Filter) New(tag entity.Tag) error {
	f.SetTag(tag)
	err := f.initializeServers(f.Tag.TagId)
	if nil != err {
		panic (err)
	}

	if err = f.initializeTags(); nil != err {
		return err
	}
	if err = f.initClusterServers(); nil != err {return err}

	f.Stats.Gather(*f)
	return nil
}

func (f *Filter) initClusterServers() (err error) {
	f.ClusterServers = make([]string, len(f.Servers))
	repo := hardware_repo.HardwareConfigRepo{}
	for i := range f.Servers {
		f.ClusterServers[i], err = repo.FetchClusterServer(f.Servers[i].Id)
		if nil != err {return}
	}

	return
}

func (f *Filter) SetTag(tag entity.Tag) {
	f.Tag = tag
}

func (f *Filter) initializeServers(TagId string) error {
	comp := f.makeQueryServersByTagComponent(TagId)
	var err error
	f.Servers, err = server.ServerRepo{}.Fetch(comp, f.scanServer)
	if nil != err {
		return err
	}

	if err = f.fetchServersIpAddr(); nil != err {
		return err
	}

	if err = f.fetchServersRedfishIpAddr(); nil != err {
		return err
	}

	return err
}

func (f *Filter) initializeTags() (err error) {
	f.Tags, err = server.TagRepo{}.FetchAll()
	return err
}

func (f *Filter) fetchServersIpAddr() (err error) {
	serverIpRepo := server.ServerIpRepo{}
	for i := range f.Servers {
		if f.Servers[i].IpAddrs, err = serverIpRepo.FetchServerIpAddrs(f.Servers[i].Id); nil != err {
			return err
		}
	}
	return nil
}

func (f *Filter) fetchServersRedfishIpAddr() (err error) {
	serverIpRepo := server.ServerIpRepo{}
	for i := range f.Servers {
		redfish, err := serverIpRepo.FetchRedfishIp(f.Servers[i].Id)
		if nil != err {
			return err
		}

		if len(redfish) > 0 {
			f.Servers[i].RedfishIp = redfish[0]
		}
	}
	return nil
}

func (f *Filter) makeQueryServersByTagComponent(TagId string) database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {
				"SERVER AS S",
				"DC AS D",
				"RACK AS R",
				"RACK_UNIT AS USTART",
				"RACK_UNIT AS UEND",
				"PORT_TYPE AS PT",
				"SERVER_STATUS AS SS",
				"STATUS_ROW AS SR",
				"SERVER_TAG AS ST",
				"SERVE",
			},

		Columns: []string{
			"D.ID",
			"D.DESCRIPTION",
			"R.ID",
			"R.DESCRIPTION",
			"USTART.ID",
			"USTART.DESCRIPTION",
			"UEND.ID",
			"UEND.DESCRIPTION",
			"PT.ID",
			"PT.DESCRIPTION",
			"S.SERIAL_NUMBER",
			"SS.ID",
			"SS.DESCRIPTION",
			"SERVE.ID",
			"SERVE.DESCRIPTION",
		},

		Selection:	"ST.TAGID = ? AND " +
					"ST.SERVERID = S.ID AND " +
					"S.ID_DC = D.ID AND " +
					"S.ID_RACK = R.ID AND " +
					"S.ID_U_START = USTART.ID AND " +
					"S.ID_U_END = UEND.ID AND " +
					"SR.DESCRIPTION = ? AND S.ID_STATUS_ROW = SR.ID AND " +
					"S.ID_PORT_TYPE = PT.ID AND " +
					"S.ID_SERVER_STATUS = SS.ID AND " +
					"S.ID_SERVE = SERVE.ID",

		SelectionArgs: []string {TagId, "available"},
	}
}

func (obj Filter) scanServer(content interface{}, row *sql.Rows) (interface{}, error) {
	server := content.(entity.Server)
	err := row.Scan(
		&server.DC.Id,
		&server.DC.Description,
		&server.Rack.Id,
		&server.Rack.Description,
		&server.UStart.Id,
		&server.UStart.Description,
		&server.UEnd.Id,
		&server.UEnd.Description,
		&server.PortType.Id,
		&server.PortType.Description,
		&server.SerialNumber,
		&server.ServerStatus.Id,
		&server.ServerStatus.Description,
		&server.ServeCustomer.Id,
		&server.ServeCustomer.Description,
	)

	return server, err
}

func (f *Filter) SearchServersByMultiTags(tags []string) (err error) {
	serverRepo := server.ServerRepo{}
	scanOnlyServerId := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		s := obj.(entity.Server)
		err := row.Scan(&s.Id)
		return s, err
	}

	f.Servers, err = serverRepo.Fetch(
		f.makeQueryServerByMultiTagsComp(tags),
		scanOnlyServerId)

	for i := range f.Servers {
		listServer, err := serverRepo.FetchById(f.Servers[i].Id)

		if nil != err {
			return err
		}

		if len(listServer) > 0 {
			f.Servers[i] = listServer[0].(entity.Server)
		}
	}

	if err = f.initClusterServers(); nil != err {return err}

	return f.fetchAllIps()
}

func (f *Filter) fetchAllIps() (err error) {
	serverIpRepo := server.ServerIpRepo{}

	for i := range f.Servers {
		f.Servers[i].IpAddrs, err = serverIpRepo.FetchServerIpAddrs(f.Servers[i].Id)
		if nil != err {
			return err
		}

		redfishIps, err := serverIpRepo.FetchRedfishIp(f.Servers[i].Id)
		if nil != err {
			return err
		}

		fmt.Println ("Redfish IP: ", redfishIps)
		if len(redfishIps) > 0 {
			f.Servers[i].RedfishIp = redfishIps[0]
		}
	}

	return nil
}

func (f *Filter) makeQueryServerByMultiTagsComp(tags []string) database.QueryComponent {
	selection := ""
	if len(tags) > 1 {
		selection = f.makeSelection(tags)
	} else {
		selection = "TAG.TITLE = ? AND SERVER_TAG.TAGID = TAG.TAGID AND SERVER.ID = SERVER_TAG.SERVERID"
	}

	comp := database.QueryComponent {
		Tables: []string {"SERVER", "SERVER_TAG", "TAG"},
		Columns: []string {"SERVER.ID"},
		Selection: selection,
		SelectionArgs: tags,
	}

	return comp

}

func (f *Filter) makeSelection(tags []string) string {
	selections := make([]string, len(tags) - 1)

	for i := range selections {
		clause := "SELECT SERVER.ID FROM SERVER, SERVER_TAG, TAG WHERE TAG.TITLE = ? AND SERVER_TAG.TAGID = TAG.TAGID AND SERVER.ID = SERVER_TAG.SERVERID "
		selections[i] = clause
	}

	return "TAG.TITLE = ? AND SERVER_TAG.TAGID = TAG.TAGID AND SERVER.ID = SERVER_TAG.SERVERID intersect " + strings.Join(selections, " intersect ")
}

type FilterStats struct {
	ServersDC8Count int
	ServersDC9Count int
	ServersHNICount int
}

func (f *FilterStats) Gather(filter Filter) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666)
	if nil != err {
		panic (err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate | log.Ltime | log.Lshortfile)

	for i := range filter.Servers {
		if filter.Servers[i].DC.Description == "HCM-DC8" {
			f.ServersDC8Count++
			InfoLogger.Println ("DC8 count: ", f.ServersDC8Count)
		} else if filter.Servers[i].DC.Description == "HCM-DC9" {
			f.ServersDC9Count++
			InfoLogger.Println ("DC9 count: ", f.ServersDC9Count)
		} else {
			f.ServersHNICount++
			InfoLogger.Println ("HNI count: ", f.ServersHNICount)
		}
	}
}


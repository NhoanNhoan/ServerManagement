package server

import (
	"CURD/database"
	"CURD/entity"
	"testing"
)

func TestServerRepo_Delete(t *testing.T) {
	serverId := "SV999"
	if err := database.Insert(icomp{
		Table: "SERVER",
		Columns: []string {"ID"},
		Values: [][]string {[]string {serverId}},
	}); nil != err {
		t.Errorf("Can't insert into server table, error: '%s'", err)
	}

	var server entity.Server = entity.Server {Id: serverId}
	var serverRepo ServerRepo = ServerRepo{}
	if err := serverRepo.Delete(server); nil != err {
		t.Errorf("Can't delete the server row, error: '%s'", err)
	}
}

func TestServerRepo_IsExists(t *testing.T) {
	serverId := "SV999"
	if err := database.Insert(icomp{
		Table: "SERVER",
		Columns: []string {"ID"},
		Values: [][]string {[]string {serverId}},
	}); nil != err {
		t.Errorf("Can't insert into server table, error: '%s'", err)
	}

	var serverRepo ServerRepo = ServerRepo{}
	if !serverRepo.IsExists(serverId) {
		t.Error("Server id is not exists")
	}

	serverRepo.Delete(entity.Server{Id: serverId})
}

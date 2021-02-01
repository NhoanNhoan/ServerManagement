package entity

import (
	"testing"
)

func TestNew(t *testing.T) {
	var server Server

	if nil != server.New("SV0000000001") {
		t.Error ("Couldn't make a new server entity")
	}
}

func TestFetchIp(t *testing.T) {
	server := Server {Id: "SV0000000001"}
	server.FetchIpAddrs()

	if (nil == server.IpAddrs) || (0 == len(server.IpAddrs)) {
		t.Error ("Couldn't get IP of this server")
	}
}

func TestFetchServices(t *testing.T) {
	server := Server {Id: "SV0000000001"}
	server.FetchServices()

	if (nil == server.Services || (0 == len(server.Services))) {
		t.Error ("Couldn't get services for server loading")
	}
}


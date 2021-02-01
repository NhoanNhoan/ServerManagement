package entity

import (
	"testing"
)

func TestNewSwitch(t *testing.T) {
	var device Switch

	if nil != device.New("SW00000001") {
		t.Error ("Couldn't make switch data")
	}
}

func TestSwitchFetchIp(t *testing.T) {
	var device Switch

	device.New("SW00000001")
	device.FetchIpAddrs()

	if nil == device.IpAddrs || 0 == len(device.IpAddrs) {
		t.Error ("Couldn't get ip", device.IpAddrs)
	}
}
package entity

import (
	"testing"
)

func TestGetIpNets(t *testing.T) {
	ipNets := GetIpNets()

	if 0 == len(ipNets) {
		t.Error ("Fail")
	}
}
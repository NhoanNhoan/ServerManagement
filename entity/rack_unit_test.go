package entity

import (
	"testing"
)

func TestGetRackUnits(t *testing.T) {
	rackUnits := GetRackUnits()
	
	if 0 == len(rackUnits) {
		t.Error ("Fail")
	}
}
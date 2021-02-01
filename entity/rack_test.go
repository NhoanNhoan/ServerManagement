package entity

import (
	"testing"
)

func TestGetRacks(t *testing.T) {
	racks := GetRacks()

	if 0 == len(racks) {
		t.Error ("Fail")
	}
}
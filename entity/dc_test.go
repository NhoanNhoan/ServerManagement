package entity

import (
	"testing"
)

func TestGetDCs(t *testing.T) {
	DCs := GetDCs()
	
	if 0 == len(DCs) {
		t.Error ("Fail")
	}
}
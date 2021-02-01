package entity

import (
	"testing"
)

func TestGetServerStates(t *testing.T) {
	states := GetServerStates()

	if 0 == len(states) {
		t.Error ("Fail")
	}
}
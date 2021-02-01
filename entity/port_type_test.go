package entity

import (
	"testing"
)

func TestGetPortTypes(t *testing.T) {
	portTypes := GetPortTypes()

	if 0 == len(portTypes) {
		t.Error ("Fail")
	}
}
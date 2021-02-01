package page

import (
	"testing"
)

func TestNewUpdateServerPage(t *testing.T) {
	var updatePage UpdateServer

	updatePage.New("SV0000000001")

	if "SV0000000001" != updatePage.Server.Id {
		t.Error ("Couldn't query server information")
	}
}

func TestInitAttributesOfUpdateServerObject(t *testing.T) {
	var updatePage UpdateServer
	var id, description string

	updatePage.init("DC", []string {"Id", "Description"}, updatePage.DCs, id, description)

	if 0 == len(updatePage.DCs) || nil == updatePage.DCs {
		t.Error ("Couldn't init DC array")
	}
}
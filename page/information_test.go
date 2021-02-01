package page

import (
	"testing"
)

func TestFindSwitchId(t *testing.T) {
	var info Information

	if "SW00000001" != info.findSwitchId("SV0000000002") {
		t.Error ("Id not found")
	}
}

func TestNewInformation(t *testing.T) {
	var info Information

	info.New("SV0000000002")

	if "SV0000000002" != info.Server.Id {
		t.Error ("Couldn't query server")
	}

	if "SW00000001" != info.Switch.Id {
		t.Error ("Couldn't query switch")
	}

	t.Error (info)
}
package server_parse

import (
	"CURD/entity"
	"testing"
	"net/http"
	"time"
)

var unit entity.Entity

func TestParseRackUnit(t *testing.T) {
	time.Sleep (4 * time.Second)
	http.Get("http://localhost:9000/unit?txtUStartId=1&txtUEndDes=2")
	ustart := e.(entity.RackUnit)

	if ("1" != ustart.Id && "2" != ustart.Description) {
		t.Error ("Fail")
	}
}
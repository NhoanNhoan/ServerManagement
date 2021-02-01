package server_parse

import (
	"CURD/entity"
	"testing"
	"net/http"
	"time"
)

var rackEntity entity.Entity

func TestParseRack(t *testing.T) {
	time.Sleep (3 * time.Second)
	http.Get("http://localhost:9000/rack?txtRackId=1&txtRackDes=2")
	rack := e.(entity.Rack)

	if ("1" != rack.Id && "2" != rack.Description) {
		t.Error ("Fail")
	}
}
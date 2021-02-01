package server_parse

import (
	"CURD/entity"
	"testing"
	"net/http"
	"time"
)

func TestParsePortType(t *testing.T) {
	time.Sleep(5 * time.Second)
	http.Get("http://localhost:9000/porttype?txtPTId=1&txtPTDes=2")
	portType := e.(entity.PortType)

	if ("1" != portType.Id && "2" != portType.Description) {
		t.Error ("Fail")
	}
}

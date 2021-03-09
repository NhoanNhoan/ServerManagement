package page

import (
	"CURD/database"
	"CURD/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchIp struct {
	entity.IpHost
}

func (obj *SearchIp) New(c *gin.Context) {
	obj.IpHost.Host = c.Query("txtIp")
	obj.IpHost = entity.QueryIpHost(obj.queryIpHostComp())

	if "" != obj.IpHost.State {
		c.String(http.StatusOK, obj.IpHost.State)
	} else {
		c.String(http.StatusNotFound, "No Ip Found")
	}
}

func (obj *SearchIp) queryIpHostComp() database.QueryComponent {
	return database.QueryComponent{
		Tables: []string {"IP_HOST"},
		Columns: []string {"ID_NET", "HOST", "STATE"},
		Selection: "HOST = ?",
		SelectionArgs: []string {obj.IpHost.Host},
	}
}

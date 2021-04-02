package page

import (
	"CURD/entity"
	"CURD/repo/server"
	"errors"
	"github.com/gin-gonic/gin"
)

type SearchIp struct {
	IpState string
}

func (obj *SearchIp) New(c *gin.Context) (err error) {
	ipStr := c.Query("txtIpSearchState")
	var ip entity.IpAddress

	if !ip.Parse(ipStr) {
		return errors.New(ipStr + " is not like format ip address")
	}

	obj.IpState, err = server.IpRepo{}.FetchState(ip)
	return err
}

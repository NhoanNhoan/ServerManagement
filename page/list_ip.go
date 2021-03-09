package page

import (
	"CURD/entity"

	"github.com/gin-gonic/gin"
)

type ListIp struct {
	IpArr []entity.IpHost
}

func (obj *ListIp) New(c *gin.Context) {
	netId := c.Query("txtNetId")
	net := entity.IpNet {Id: netId}
	obj.IpArr = entity.FetchIpHostArray(net)
}

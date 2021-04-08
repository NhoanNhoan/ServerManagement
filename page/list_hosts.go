package page

import (
	"CURD/entity"
	"CURD/repo/server"

	"github.com/gin-gonic/gin"
)

type ListIp struct {
	IpArr []entity.IpAddress
}

func (obj *ListIp) New(c *gin.Context) (err error) {
	repo := server.NetworkPortionRepo{}
	filterBy := c.Query("txtFilter")
	portion := entity.NetworkPortion{Id: c.Query("txtNetId")}

	if "" == filterBy {
		obj.IpArr, err = repo.FetchHosts(portion)
		return
	}

	obj.IpArr, err = repo.FetchHostsByState(portion, filterBy)

	return
}
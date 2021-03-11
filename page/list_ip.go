package page

import (
	"CURD/database"
	"CURD/entity"

	"github.com/gin-gonic/gin"
)

type ListIp struct {
	IpArr []entity.IpHost
}

func (obj *ListIp) New(c *gin.Context) {
	netId := c.Query("txtNetId")
	filterBy := c.Query("txtFilter")

	if "" == filterBy {
		net := entity.IpNet{Id: netId}
		obj.IpArr = entity.FetchIpHostArray(net)
	} else {
		if filterBy == "available" {
			obj.availableIpArr(netId)
		} else {
			obj.usedIpArr(netId)
		}
	}
}

func (obj *ListIp) availableIpArr(netId string) {
	comp := obj.stateIpComp(netId, "available")
	obj.IpArr = entity.FetchIpHostArr(comp)
}

func (obj *ListIp) usedIpArr(netId string) {
	comp := obj.stateIpComp(netId, "used")
	obj.IpArr = entity.FetchIpHostArr(comp)
}

func (obj *ListIp) stateIpComp(netId string, state string) database.QueryComponent {
	return database.QueryComponent{
		Tables: []string {"IP_HOST"},
		Columns: []string {"ID_NET", "HOST", "STATE"},
		Selection: "ID_NET = ? AND STATE = ?",
		SelectionArgs: []string {netId, state},
	}
}
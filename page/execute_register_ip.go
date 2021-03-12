package page

import (
	"CURD/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExecuteRegisterIp struct {
	Msg string
	AvailableHosts int
	entity.IpNet
}

func (ex *ExecuteRegisterIp) New(c *gin.Context) {
	ex.IpNet.Value = c.PostForm("txtIpNet")
	ex.IpNet.Netmask, _ = strconv.Atoi(c.PostForm("txtNetmask"))
	err := ex.IpNet.Insert()

	if nil != err {
		ex.Msg = "Can't insert new ip net"
		panic (err)
	}

	ex.AvailableHosts = ex.IpNet.CalculateHosts()
	ex.Msg = "Success"
}


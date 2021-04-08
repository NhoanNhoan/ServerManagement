package page

import (
	"CURD/entity"
	"github.com/gin-gonic/gin"
	"strings"
)

type UpdateStateIp struct {
	Msg string
	Hosts []entity.IpHost
}

func (obj *UpdateStateIp) New(c *gin.Context) {
	netId := c.PostForm("txtNetId")
	reqStr := c.PostForm("txtIpStates")

	obj.parseIpAddress(reqStr, netId)
	if err := obj.ExecuteUpdate(); nil != err  {
		panic (err)
	}
	obj.Msg = "Success"
}

func (obj *UpdateStateIp) parseIpAddress(reqStr string, netId string) {
	content := strings.Split(reqStr, ";")
	obj.Hosts = make([]entity.IpHost, len(content))

	for idx := range content {
		if values := strings.Split(content[idx], "-"); len(values) == 2 {
			obj.Hosts[idx] = entity.IpHost {
				IpNet: entity.IpNet{Id: netId},
				Host: values[0],
				State: values[1],
			}
		}
	}
}

func (obj *UpdateStateIp) ExecuteUpdate() error {
	for _, host := range obj.Hosts {
		if err := host.Update(); nil != err {
			return err
		}
	}

	return nil
}

package page

import (
	"strings"

	"CURD/entity"

	"github.com/gin-gonic/gin"
)

type ExecuteRegisterSwitch struct {
	Msg string
	entity.Switch
}

type eswitch = ExecuteRegisterSwitch

func (e *eswitch) Execute(c *gin.Context) error {
	e.parseSwitch(c)
	var err error

	err = e.Switch.Insert()
	if nil != err  {
		e.Msg = "Can't insert switch"
		return err
	}

	err = e.InsertIps()
	if nil != err {
		e.Msg = "Can't insert ipaddress"
	}

	e.Msg = "Success"

	return err
}

func (e *eswitch) parseSwitch(c *gin.Context) {
	e.Switch.DC = e.parseDC(c)
	e.Switch.Rack = e.parseRack(c)
	e.Switch.UStart = e.parseUStart(c)
	e.Switch.UEnd = e.parseUEnd(c)
	e.Switch.IpAddrs = e.parseIpAddrs(c)
	e.Switch.MaximumPort = c.PostForm("txtMaxPorts")
}

func (e *eswitch) parseDC(c *gin.Context) entity.DataCenter {
	return entity.DataCenter {Id: c.PostForm("cbDC")}
}

func (e *eswitch) parseRack(c *gin.Context) entity.Rack {
	return entity.Rack {Id: c.PostForm("cbRack")}
}

func (e *eswitch) parseUStart(c *gin.Context) entity.RackUnit {
	return entity.RackUnit {Id: c.PostForm("cbUStart")}
}

func (e *eswitch) parseUEnd(c *gin.Context) entity.RackUnit {
	return entity.RackUnit {Id: c.PostForm("cbUEnd")}
}

func (e *eswitch) parseIpAddrs(c *gin.Context) []entity.IpAddress {
	ipStr := c.PostForm("txtAllIp")
	ips := strings.Split(ipStr, ",")
	ipArr := make([]entity.IpAddress, len(ips))

	for i := range ipArr {
		items := strings.Split(ips[i], "-")
		ipArr[i] = entity.IpAddress {
						IpNet: entity.IpNet{Id: items[0]}, 
						IpHost: items[1],
					}
	}

	return ipArr
}
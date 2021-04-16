package page

import "CURD/entity"

type ServerContextParser interface {
	ParseDC() entity.DataCenter
	ParseRack() entity.Rack
	ParseUSTart() entity.RackUnit
	ParseUEnd() entity.RackUnit
	ParsePortType() entity.PortType
	ParseServerState() entity.ServerStatus
	ParseServeCustomer() entity.ServeCustomer
	ParseRedfishIp() entity.IpAddress
	ParseIpAddresses() []entity.IpAddress
	ParseTags() []entity.Tag

}

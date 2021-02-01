package page

import (
	"CURD/entity"
)

type UpdateServer struct {
	DCs []entity.DataCenter
	Racks []entity.Rack
	RackUnits []entity.RackUnit
	PortTypes []entity.PortType
	ServerStates []entity.ServerStatus
	IpNets []entity.IpNet
	entity.Server
}

func (obj *UpdateServer) New(server entity.Server) {
	obj.Server = server
	obj.Server.FetchIpAddrs()
	obj.Server.FetchServices()

	obj.DCs = entity.GetDCs()
	obj.Racks = entity.GetRacks()
	obj.RackUnits = entity.GetRackUnits()
	obj.PortTypes = entity.GetPortTypes()
	obj.ServerStates = entity.GetServerStates()
	obj.IpNets = entity.GetIpNets()
}
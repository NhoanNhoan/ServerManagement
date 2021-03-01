package page


import (
	"CURD/entity"
)


type RegistrationServer struct {
	AllDataCenters []entity.DataCenter
	AllRacks []entity.Rack
	AllRackUnits []entity.RackUnit
	AllPortTypes []entity.PortType
	AllServerStates []entity.ServerStatus
	AllIpNets []entity.IpNet
	Tags []entity.Tag
}

func (r *RegistrationServer) New() {
	r.initAllDataCenters()
	r.initAllRacks()
	r.initAllRackUnits()
	r.initAllPortType()
	r.initAllServerStates()
	r.initAllIpNets()
	r.initTags()
}


func (r *RegistrationServer) initAllDataCenters() {
	r.AllDataCenters = entity.GetDCs()
}

func (r *RegistrationServer) initAllRacks() {
	r.AllRacks = entity.GetRacks()
}

func (r *RegistrationServer) initAllRackUnits() {
	r.AllRackUnits = entity.GetRackUnits()
}

func (r *RegistrationServer) initAllPortType() {
	r.AllPortTypes = entity.GetPortTypes()
}

func (r *RegistrationServer) initAllServerStates() {
	r.AllServerStates = entity.GetServerStates()
}

func (r *RegistrationServer) initAllIpNets() {
	r.AllIpNets = entity.GetIpNets()
}

func (r *RegistrationServer) initTags() {
	r.Tags = entity.FetchAllTags()
}
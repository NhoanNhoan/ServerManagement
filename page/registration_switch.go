package page


import (
	"CURD/entity"
)


type RegistrationSwitch struct {
	AllDataCenters []entity.DataCenter
	AllRacks []entity.Rack
	AllRackUnits []entity.RackUnit
	AllIpNets []entity.IpNet
}

func (r *RegistrationSwitch) New() {
	//r.initAllDataCenters()
	//r.initAllRacks()
	//r.initAllRackUnits()
	//r.initAllIpNets()
}
//
//func (r *RegistrationSwitch) initAllDataCenters() {
//	r.AllDataCenters = entity.GetDCs()
//}
//
//func (r *RegistrationSwitch) initAllRacks() {
//	r.AllRacks = entity.GetRacks()
//}
//
//func (r *RegistrationSwitch) initAllRackUnits() {
//	r.AllRackUnits = entity.GetRackUnits()
//}
//
//func (r *RegistrationSwitch) initAllIpNets() {
//	r.AllIpNets = entity.GetIpNets()
//}
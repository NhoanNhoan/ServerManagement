package page


import (
	"CURD/entity"
	"CURD/repo/server"
)


type RegistrationServer struct {
	AllDataCenters []entity.DataCenter
	AllRacks []entity.Rack
	AllRackUnits []entity.RackUnit
	AllPortTypes []entity.PortType
	AllServerStates []entity.ServerStatus
	AllServes []entity.ServeCustomer
	Tags []entity.Tag
	HardwareData
}

func (r *RegistrationServer) New() (err error) {
	dcRepo := server.DCRepo{}
	if r.AllDataCenters, err = dcRepo.FetchAll(); nil != err {
		return err
	}

	rackRepo := server.RackRepo{}
	if r.AllRacks, err = rackRepo.FetchAll(); nil != err {
		return err
	}

	rackUnitRepo := server.RackUnitRepo{}
	if r.AllRackUnits, err = rackUnitRepo.FetchAll(); nil != err {
		return err
	}

	portTypeRepo := server.PortTypeRepo{}
	if r.AllPortTypes, err = portTypeRepo.FetchAll(); nil != err {
		return err
	}

	serverStateRepo := server.ServerStatusRepo{}
	if r.AllServerStates, err = serverStateRepo.FetchAll(); nil != err {
		return err
	}

	tagRepo := server.TagRepo{}
	if r.Tags, err = tagRepo.FetchAll(); nil != err {
		return err
	}

	serveCustomerRepo := server.ServeCustomerRepo{}
	if r.AllServes, err = serveCustomerRepo.FetchAll(); nil != err {
		return err
	}

	return r.HardwareData.New()

	//r.initAllDataCenters()
	//r.initAllRacks()
	//r.initAllRackUnits()
	//r.initAllPortType()
	//r.initAllServerStates()
	//r.initAllIpNets()
	//r.initTags()
}

//
//func (r *RegistrationServer) initAllDataCenters() {
//	r.AllDataCenters = entity.GetDCs()
//}
//
//func (r *RegistrationServer) initAllRacks() {
//	r.AllRacks = entity.GetRacks()
//}
//
//func (r *RegistrationServer) initAllRackUnits() {
//	r.AllRackUnits = entity.GetRackUnits()
//}
//
//func (r *RegistrationServer) initAllPortType() {
//	r.AllPortTypes = entity.GetPortTypes()
//}
//
//func (r *RegistrationServer) initAllServerStates() {
//	r.AllServerStates = entity.GetServerStates()
//}
//
//func (r *RegistrationServer) initAllIpNets() {
//	r.AllIpNets = entity.GetIpNets()
//}
//
//func (r *RegistrationServer) initTags() {
//	r.Tags = entity.FetchAllTags()
//}
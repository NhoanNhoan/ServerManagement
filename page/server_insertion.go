package page

import (
	"CURD/entity"
	"CURD/entity/server/hardware"
	"CURD/repo/server"
)

type ServerInsertion struct {
	entity.Server
	hardware.HardwareConfig
	ServerTags []entity.Tag
	HardwareExecution
}

func (ins ServerInsertion) ExecuteServer() (err error) {
	if err = ins.insertServerRedfishIp(); nil != err {return}
	if err = ins.insertServerIpAddresses(); nil != err {return}
	return ins.insertServer()
}

func (ins ServerInsertion) insertServerRedfishIp() error {
	repo := server.ServerIpRepo{}
	return repo.InsertRedfishIpAddresses(ins.Server.Id, ins.Server.RedfishIp)
}

func (ins ServerInsertion) insertServerIpAddresses() error {
	repo := server.ServerIpRepo{}
	return repo.InsertNormalIpAddresses(ins.Server.Id, ins.Server.IpAddrs...)
}

func (ins ServerInsertion) insertServer() error {
	repo := server.ServerRepo{}
	return repo.Insert(ins.Server)
}

func (ins ServerInsertion) ExecuteTags() error {
	repo := server.ServerTagRepo{}
	return repo.Insert(ins.Server.Id, ins.ServerTags...)
}
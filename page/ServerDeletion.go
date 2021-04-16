package page

import (
	"CURD/entity"
	"CURD/entity/server/hardware"
	"CURD/repo/server"
)

type ServerDeletion struct {
	entity.Server
	hardware.HardwareConfig
	ServerTags []entity.Tag
	HardwareExecution
}

func (del ServerDeletion) ExecuteServer() (err error) {
	if err = del.deleteServerRedfishIp(); nil != err {return}
	if err = del.deleteServerIpAddresses(); nil != err {return}
	return del.deleteServer()
}

func (del ServerDeletion) deleteServerRedfishIp() error {
	repo := server.ServerIpRepo{}
	return repo.Delete(del.Server.Id, del.Server.RedfishIp)
}

func (del ServerDeletion) deleteServerIpAddresses() error {
	repo := server.ServerIpRepo{}
	return repo.Delete(del.Server.Id, del.Server.IpAddrs...)
}

func (del ServerDeletion) deleteServer() error {
	repo := server.ServerRepo{}
	return repo.Delete(del.Server)
}

func (del ServerDeletion) ExecuteTags() error {
	repo := server.ServerTagRepo{}
	return repo.Delete(repo.MakeDeleteComp(del.Server.Id))
}
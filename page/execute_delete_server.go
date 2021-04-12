package page

import (
	"CURD/entity"
	"CURD/repo/server"
	"CURD/repo/server/hardware_repo"
	"errors"
	"github.com/gin-gonic/gin"
)

type ExecuteDeleteServer struct {
	ServerId string
}

func (ex *ExecuteDeleteServer) New(c *gin.Context) error {
	ex.ServerId = c.Query("txtServerId")
	repo := server.ServerRepo{}
	if !repo.IsExists(ex.ServerId) {
		return errors.New("The server is not found")
	}
	return nil
}

// Delete a server must:
// delete ip of the server
// delete switch connection
// delete hardware config of server
// delete cpus, rams,.. config
// ips that server dont use will set state 'available'
// delete switch connect info like port type, sig
func (ex *ExecuteDeleteServer) Execute(c *gin.Context) (err error) {
	if err = ex.changeStateServerIps(); nil != err {return}
	if err = ex.deleteHardwareConfig(); nil != err {return}
	repo := server.ServerRepo{}
	return repo.Delete(entity.Server{Id: ex.ServerId})
}

func (ex ExecuteDeleteServer) changeStateServerIps() (err error) {
	ips, err := ex.fetchIps()
	if nil != err {return}
	return ex.updateIpsState(ips...)
}

func (ex ExecuteDeleteServer) fetchIps() ([]entity.IpAddress, error) {
	repo := server.ServerIpRepo{}
	return repo.FetchServerIpAddrs(ex.ServerId)
}

func (ex ExecuteDeleteServer) updateIpsState(ips ...entity.IpAddress) (err error) {
	repo := server.IpRepo{}
	for i := range ips {
		err := repo.UpdateState(ips[i], "available");
		if nil != err {break}
	}
	return
}

func (ex ExecuteDeleteServer) fetchHardwareConfigId() (string, error) {
	repo := hardware_repo.HardwareConfigRepo{}
	return repo.FetchConfigId(ex.ServerId)
}

func (ex ExecuteDeleteServer) deleteHardwareConfig() error {
	hwId, err := ex.fetchHardwareConfigId()
	if nil != err {return err}
	if err = ex.deleteHardwareCpu(hwId); nil != err {return err}
	if err = ex.deleteHardwareRam(hwId); nil != err {return err}
	if err = ex.deleteHardwareDisk(hwId); nil != err {return err}
	if err = ex.deleteHardwareNic(hwId); nil != err {return err}
	if err = ex.deleteHardwareRaid(hwId); nil != err {return err}
	if err = ex.deleteHardwarePsu(hwId); nil != err {return err}
	if err = ex.deleteHardwareMnt(hwId); nil != err {return err}
	return err
}

func (ex ExecuteDeleteServer) deleteHardwareCpu(hwId string) error {
	repo := hardware_repo.HardwareCPURepo{}
	return repo.Delete(hwId)
}

func (ex ExecuteDeleteServer) deleteHardwareRam(hwId string) error {
	repo := hardware_repo.HardwareRamRepo{}
	return repo.Delete(hwId)
}

func (ex ExecuteDeleteServer) deleteHardwareDisk(hwId string) error {
	repo := hardware_repo.HardwareDiskRepo{}
	return repo.Delete(hwId)
}

func (ex ExecuteDeleteServer) deleteHardwareNic(hwId string) error {
	repo := hardware_repo.HardwareNicRepo{}
	return repo.Delete(hwId)
}

func (ex ExecuteDeleteServer) deleteHardwareRaid(hwId string) error {
	repo := hardware_repo.HardwareRaidRepo{}
	return repo.Delete(hwId)
}

func (ex ExecuteDeleteServer) deleteHardwarePsu(hwId string) error {
	repo := hardware_repo.HardwarePsuRepo{}
	return repo.Delete(hwId)
}

func (ex ExecuteDeleteServer) deleteHardwareMnt(hwId string) error {
	repo := hardware_repo.HardwareManagementRepo{}
	return repo.Delete(hwId)
}

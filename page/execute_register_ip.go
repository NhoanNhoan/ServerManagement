package page

import (
	"CURD/entity"
	"CURD/repo/server"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NetworkPortionRegistration struct {
	entity.NetworkPortion
	NumAvailableHosts int
	Msg string
}

func (regis *NetworkPortionRegistration) New(c *gin.Context) (err error) {
	if err = regis.initNetworkPortion(c); nil == err {
		regis.initNumHosts()
		return
	}
	return
}

func (regis *NetworkPortionRegistration) initNetworkPortion(c *gin.Context) (err error) {
	raw := c.PostForm("txtNetwork")
	if regis.NetworkPortion.Value, err = server.StadardlizeNetworkPortion(raw); nil != err {
		return err
	}
	regis.NetworkPortion.Netmask, err = strconv.Atoi(c.PostForm("txtNetmask"))
	if nil != err {
		err = errors.New("Netmask is wrong!")
	}
	return
}

func (regis *NetworkPortionRegistration) initNumHosts() {
	regis.NumAvailableHosts = regis.NetworkPortion.CalculateNumHosts()
}

func (regis *NetworkPortionRegistration) Execute() (err error) {
	repo := server.NetworkPortionRepo{}
	regis.NetworkPortion.Id = repo.GenerateId()
	return repo.Insert(regis.NetworkPortion)
}

func (regis *NetworkPortionRegistration) SetMsg(msg string) {
	regis.Msg = msg
}
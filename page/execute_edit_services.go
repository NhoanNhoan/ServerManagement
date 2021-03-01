package page

import (
	"strings"

	"CURD/entity"

	"github.com/gin-gonic/gin"
)

type ExecuteEditServices struct {
	Msg string
	EditServices
}

func (obj *ExecuteEditServices) Execute(c *gin.Context) {
	obj.initialize(c)
	err := entity.DeleteServicesByServerId(obj.EditServices.ServerId)
	if nil != err {
		obj.Msg = "Error"
		panic (err)
	}

	if !obj.AreEmptyServices() {
		err = entity.InsertServices(obj.EditServices.ServerId, obj.EditServices.Services)
		if nil != err {
			obj.Msg = "Can't insert new services"
			panic (err)
		}

		obj.Msg = "Insert services successfully!"
	} else {
		obj.Msg = "Services are empty"
	}
}

func (obj *ExecuteEditServices) initialize(c *gin.Context) {
	obj.EditServices.ServerId = c.PostForm("txtServerId")
	serviceStr := c.PostForm("txtServices")
	obj.parseStringToServices(serviceStr)
}

func (obj *ExecuteEditServices) parseStringToServices(s string) {
	if len(s) > 0 {
		obj.EditServices.Services = strings.Split(s, ",")
	}
}

func (obj *ExecuteEditServices) AreEmptyServices() bool {
	return (0 == len(obj.EditServices.Services))
}
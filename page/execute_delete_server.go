package page

import (
	"CURD/entity"
	"github.com/gin-gonic/gin"
)

type ExecuteDeleteServer struct {
	Msg string
	entity.Server
}

func (obj *ExecuteDeleteServer) ExecuteDelete(c *gin.Context) string {
	//obj.Server.Id = c.PostForm("txtServerId")
	//err := obj.Server.Delete()
	//
	//if nil !=  err {
	//	return "Can't delete"
	//}

	return "Success"

}

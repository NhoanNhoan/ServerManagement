package page

import (
	"CURD/entity"

	"github.com/gin-gonic/gin"
)

type EditServices struct {
	ServerId string
	Services []string
}

func (s *EditServices) New(c *gin.Context) {
	s.initServerIdByPostForm(c)
	s.initServices()
}

func (s *EditServices) initServerIdByPostForm(c *gin.Context) {
	s.ServerId = c.PostForm("txtServerId")
}

func (s *EditServices) initServices() {
	s.Services = entity.FetchServicesByServerId(s.ServerId)
}
package controllers

import (
	"CURD/entity"
	"CURD/page"
	"CURD/page/error_page"
	"CURD/entity/error_entity"
	"github.com/gin-gonic/gin"
	"net/http"

	"fmt"
)

type Router struct {
	router *gin.Engine
}

func (obj *Router) New() {
	obj.router = gin.Default()
	obj.router.Static("/public", "./public")

	obj.handleHome()
	obj.handleInfo()
	obj.handleServers()
	obj.handleUpdateServer()
	obj.handleExecuteModify()
	obj.handleListError()
	obj.handleViewError()
}
func (obj *Router) Run(ip string) {
	obj.router.Run(ip)
}

func (obj *Router) handleHome() {
	obj.router.GET("/home", func(c *gin.Context) {
		home := page.Home{}
		home.New()
		obj.router.LoadHTMLFiles("templates/home.html")
		c.HTML(http.StatusOK, "templates/home.html", gin.H{
			"info": home.DCs,
		})
	})
}

func (obj *Router) handleServers() {
	obj.router.POST("/server/list", func(c *gin.Context) {
		DC := entity.DataCenter{
			c.PostForm("txtId"),
			c.PostForm("txtName"),
		}

		fmt.Println (DC)
		server := page.Servers{DC, nil}
		server.New(DC.Id)

		obj.router.LoadHTMLFiles("templates/server/list.html")
		c.HTML(http.StatusOK, "templates/server/list.html", server)
	})
}

func (obj *Router) handleInfo() {
	obj.router.POST("/server/information", func(c *gin.Context) {
		idServer := c.PostForm("txtIdServer")

		var infoPage page.Information
		infoPage.New(idServer)

		obj.router.LoadHTMLFiles("templates/server/information.html")
		c.HTML(http.StatusOK, "templates/server/information.html", infoPage)
	})
}

func (obj *Router) handleUpdateServer() {
	obj.router.POST("/server/modify", func(c *gin.Context) {
		id := c.PostForm("txtIdServer")
		idDC := c.PostForm("txtDCId")
		idRack := c.PostForm("txtRackId")
		idUStart := c.PostForm("txtUStartId")
		idUEnd := c.PostForm("txtUEndId")
		numDisks := c.PostForm("txtNumDisks")
		maker := c.PostForm("txtMaker")
		serialNumber := c.PostForm("txtSerialNumber")

		server := entity.Server {
			Id: id,
			DC: entity.DataCenter {
				Id: idDC,
			},
			Rack: entity.Rack {
				Id: idRack,
			},
			UStart: entity.RackUnit {
				Id: idUStart,
			},
			UEnd: entity.RackUnit {
				Id: idUEnd,
			},
			NumDisks: numDisks,
			Maker: maker,
			SerialNumber: serialNumber,
		}

		var updatePage page.UpdateServer
		updatePage.New(server)

		obj.router.LoadHTMLFiles("templates/server/modify.html")
		c.HTML(http.StatusOK, "templates/server/modify.html", updatePage)
	})
}

func (obj *Router) handleExecuteModify() {
	obj.router.POST("server/execute_modify", func(c *gin.Context) {
			server := getServerFromPostForm(c)
			var executeModify page.ExecuteModify
			executeModify.New(server)

			obj.router.LoadHTMLFiles("templates/server/execute_modify.html")
			c.HTML(http.StatusOK,
				"templates/server/execute_modify.html",
				executeModify)
		})
}

func getServerFromPostForm(c *gin.Context) entity.Server {
	return entity.Server{
		Id:           c.PostForm("txtIdServer"),
		DC:   		  getDCFromPostForm(c),
		Rack:         getRackFromPostForm(c),
		UStart:       getRackUnitFromPostForm("cbUStartId", c),
		UEnd:         getRackUnitFromPostForm("cbUEndId", c),
		NumDisks:     c.PostForm("txtNumDisks"),
	//	Maker:        c.PostForm("txtMaker"),
		PortType:     getPortTypeFromPostForm(c),
		SerialNumber: c.PostForm("txtSerialNumber"),
		ServerStatus: getServerStatusFromPostForm(c),
	}
}

func getDCFromPostForm(c *gin.Context) entity.DataCenter {
	return entity.DataCenter{
		Id: c.PostForm("cbDCId"),
	}
}

func getRackFromPostForm(c *gin.Context) entity.Rack {
	return entity.Rack{
		Id: c.PostForm("cbRackId"),
	}
}

func getRackUnitFromPostForm(name string, c *gin.Context) entity.RackUnit {
	return entity.RackUnit{
		Id: c.PostForm(name),
	}
}

func getPortTypeFromPostForm(c *gin.Context) entity.PortType {
	return entity.PortType{
		Id: c.PostForm("cbPortTypeId"),
	}
}

func getServerStatusFromPostForm(c *gin.Context) entity.ServerStatus {
	return entity.ServerStatus{
		Id: c.PostForm("cbStatusId"),
	}
}

func (obj *Router) handleListError() {
	obj.router.GET("/error/list", func (c *gin.Context) {
		var errorsPage error_page.Errors
		errorsPage.New()

		obj.router.LoadHTMLFiles("templates/error/list.html")
		c.HTML(http.StatusOK, "templates/error/list.html", errorsPage)
	})
}

func (obj *Router) handleViewError() {
	obj.router.POST("/error/view", func (c *gin.Context) {
		var viewPage error_page.ErrorView
		var errData error_entity.Error = ErrorFromPostForm(c)

		viewPage.New(errData)

		obj.router.LoadHTMLFiles("templates/error/view.html")
		c.HTML(http.StatusOK, "templates/error/view.html", viewPage)
	})
}

func ErrorFromPostForm(c *gin.Context) error_entity.Error {
	id := c.PostForm("txtIdError")
	summary := c.PostForm("txtSummary")
	des := c.PostForm("txtDescription")
	sol := c.PostForm("txtSolution")
	occurs := c.PostForm("txtOccurs")
	idServer := c.PostForm("txtIdServer")
	idState := c.PostForm("txtIdState")
	errStateDes := c.PostForm("txtErrStateDes")

	return error_entity.Error {
		Id: id,
		Summary: summary,
		Description: des,
		Solution: sol,
		Occurs: occurs,
		Server: entity.Server {Id: idServer},
		ErrorState: error_entity.ErrorState{Id: idState, Description: errStateDes},
	}
}
package main

import (
	_ "CURD/controllers"
	"CURD/entity"
	"CURD/entity/error_entity"
	"CURD/page"
	"CURD/page/error_page"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "html/template"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	// client := r.Group("/api")
	// {
	// 	client.GET("/story", func (c *gin.Context) {
	// 		c.HTML(http.StatusOK, "templates/index.html", gin.H {
	// 			"title": "Page",
	// 		})
	// 	})

	// 	client.GET("/story/hello", func (c *gin.Context) {
	// 		c.HTML(http.StatusOK, "templates/hello.html", gin.H {
	// 			"title": "Hello",
	// 		})
	// 	})

	// 	client.GET("/server/:name", func (c *gin.Context) {
	// 		serverName := c.Query("name")
	// 		c.HTML(http.StatusOK, "templates/details.html", gin.H {
	// 			"title": serverName,
	// 			"Name": server.Name,
	// 			"Position": server.Position,/vendor/md
	// 			"Status": server.Status,
	// 			"Thermal": server.Thermal,
	// 			"Temperature": server.Temperature,
	// 		})
	// 	})

	// 	client.POST("/story", func (c * gin.Context) {
	// 		file, _ := c.FormFile("file")
	// 		c.SaveUploadedFile(file, file.Filename)
	// 		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	// 	})

	// 	client.PATCH("/story/update/:id", controllers.Update)
	// 	client.DELETE("/story/:id", controllers.Delete)
	// }

	// servers := []model.Server {
	// 				makeServer(),
	// 				makeServer(),
	// 				makeServer(),
	// 				makeServer(),
	// 				makeServer(),
	// 			}

	// dc1 := model.DataCenter{"ABC", "Data Center 1"}
	// dc2 := model.DataCenter{"XYZ", "Data Center 2"}
	// dc3 := model.DataCenter{"PQS", "Data Center 3"}
	// dc4 := model.DataCenter{"DEF", "Data Center 4"}

	// home := page.Home{[]model.DataCenter{dc1, dc2, dc3, dc4}}

	HandleHome(r)
	HandleServers(r)
	HandleInfo(r)	
	HandleUpdateServer(r)
	HandleExecuteModify(r)
	HandleRegisterServer(r)

	HandleListError(r)
	HandleViewError(r)
	HandleErrorExecuteUpdate(r)

	// HandleLogin(r)
	// HandleRegisterServer(r, nil)
	// HandleSummary(r)
	// HandleErrorRegister(r)
	// HandleModifyError(r)
	// HandleViewError(r)

	// using bind json
	// r.POST("/login", func (c *gin.Context) {
	// 	var json Login
	// 	if err := c.BindJSON(&json); nil != err {
	// 		c.JSON(http.StatusBadRequest, gin.H {"error": err.Error()})
	// 		return
	// 	}

	// 	if json.User != "manu" || json.Password != "123" {
	// 		c.JSON(http.StatusUnauthorized, gin.H {"status": "unauthorized"})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H {"status": "you are logged in"})
	// 	fmt.Println(json)
	// })

	return r
}

func main() {
	r := setupRouter()
	r.Static("/templates", "./templates")

	r.Run(":8080")
}

func HandleHome(router *gin.Engine) {
	router.GET("/home", func(c *gin.Context) {
		home := page.Home{}
		home.New()
		router.LoadHTMLFiles("templates/home.html")
		c.HTML(http.StatusOK, "templates/home.html", gin.H{
			"info": home.DCs,
		})
	})
}

func HandleServers(r *gin.Engine) {
	r.POST("/server/list", func(c *gin.Context) {
		DC := entity.DataCenter{
			c.PostForm("txtId"),
			c.PostForm("txtName"),
		}

		server := page.Servers{DC, nil}
		server.New(DC.Id)

		r.LoadHTMLFiles("templates/server/list.html")
		c.HTML(http.StatusOK, "templates/server/list.html", server)
	})
}

func HandleInfo(r *gin.Engine) {
	r.POST("/server/information", func(c *gin.Context) {
		idServer := c.PostForm("txtIdServer")

		var infoPage page.Information
		infoPage.New(idServer)

		r.LoadHTMLFiles("templates/server/information.html")
		c.HTML(http.StatusOK, "templates/server/information.html", infoPage)
	})
}

func HandleUpdateServer(r *gin.Engine) {
	r.POST("/server/modify", func(c *gin.Context) {
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

		r.LoadHTMLFiles("templates/server/modify.html")
		c.HTML(http.StatusOK, "templates/server/modify.html", updatePage)
	})
}

func HandleExecuteModify(r *gin.Engine) {
	r.POST("server/execute_modify", func(c *gin.Context) {
		server := getServerFromPostForm(c)
		var executeModify page.ExecuteModify
		executeModify.New(server)

		r.LoadHTMLFiles("templates/server/execute_modify.html")
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

func HandleRegisterServer(r *gin.Engine) {
	r.GET("/server/register", func (c *gin.Context) {
		r.LoadHTMLFiles("templates/server/register.html")
		c.HTML(http.StatusOK, "templates/server/register.html", nil)
	})
}

func getServerStatusFromPostForm(c *gin.Context) entity.ServerStatus {
	return entity.ServerStatus{
		Id: c.PostForm("cbStatusId"),
	}
}

func HandleListError(r *gin.Engine) {
	r.GET("/error/list", func (c *gin.Context) {
		var errorsPage error_page.Errors
		errorsPage.New()

		r.LoadHTMLFiles("templates/error/list.html")
		c.HTML(http.StatusOK, "templates/error/list.html", errorsPage)
	})
}

func HandleViewError(r *gin.Engine) {
	r.POST("/error/view", func (c *gin.Context) {
		var viewPage error_page.ErrorView
		var errData error_entity.Error = ErrorFromPostForm(c)

		viewPage.New(errData)

		r.LoadHTMLFiles("templates/error/view.html")
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

func HandleErrorExecuteUpdate(r *gin.Engine) {
	r.POST("error/update_error", func (c *gin.Context) {
		update := error_page.ExecuteUpdateError {}
		update.New(c)
		err := update.Execute()
		if nil != err {
			panic (err)
		}

		r.LoadHTMLFiles("templates/error/execute_update.html")
		c.HTML(http.StatusOK, "templates/error/execute_update.html", update)
	})
}



/*func HandleLogin(r *gin.Engine) {
	r.GET("/user/login", func (c *gin.Context) {
		r.LoadHTMLFiles("templates/user/login.html")
		c.HTML(http.StatusOK, "templates/user/login.html", nil)
	})
}

func HandleSummary(r *gin.Engine) {
	r.GET("/server/summary", func (c *gin.Context) {
		r.LoadHTMLFiles("templates/server/summary.html")
		c.HTML(http.StatusOK, "templates/server/summary.html", nil)
	})
}

func HandleErrorRegister(r *gin.Engine) {
	r.GET("/error/register", func (c *gin.Context) {
		r.LoadHTMLFiles("templates/error/register.html")
		c.HTML(http.StatusOK, "templates/error/register.html", nil)
	})
}

func HandleModifyError(r *gin.Engine) {
	r.GET("/error/modify", func (c *gin.Context) {
		r.LoadHTMLFiles("templates/error/modify.html")
		c.HTML(http.StatusOK, "templates/error/modify.html", makeError())
	})
}

// List Errors testing area
func makeListErrorPage() *page.ListError {
	return &page.ListError {
		"Server Name",
		[]model.Error {
			makeError(),
			makeError(),
			makeError(),
			makeError(),
			makeError(),
			makeError(),
			makeError(),
			makeError(),
			makeError(),
			makeError(),
			makeError(),
			makeError(),
			makeError(),
			makeError(),
			makeError(),
		},
	}

}*/

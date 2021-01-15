package main

import (
	"github.com/gin-gonic/gin"
	_ "CURD/controllers"
	"net/http"
	"CURD/model"	
	"CURD/page"

	_ "fmt"
	_ "html/template"
	"time"

	"CURD/database"
)

// Get Parameters in path: Param
// Get Parameters in GET method: DefaultQuery("name", "value") for default
// Query("name") for value of parameter
// QueryMap("name") for map
// Get Parameters in POST method: DefaultPostForm, PostForm, PostFormMap, respectively
// Get file upload, using PostFile
// HTML Reponse
// r.LoadHTMLGlob("templates/**")
// File html: {{ define "templates/index.html"}}{{end}}
// Response by json, html, string: gin.Context object supports JSON, HTML, String methods.
// String method of gin.Context get http status code, content, value. Ex: c.String(http.StatusOK, "hello %s", value)
// JSON method of gin.Context get http status code, g.H object. Ex: c.JSON(http.StatusOK, gin.H{"status":  "posted","message": message,"nick": nick})

// Every Context have Keys attributes type map[string]interface{}.
// How using? Ex: if every user have id, Let store it by c.Set("username", id)
// It's same global dictionary in php
// It's store custom data of user

// Bind: When client send a request that contains some info
// info will be unmarshall to json(xml, yml)
// must set json:"fieldname" in custom struct.

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

	info := makeInfoPage()
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
	HandleLogin(r)
	HandleRegisterServer(r, nil)
	HandleInfo(r, &info)
	HandleSummary(r)
	HandleServers(r)
	HandleErrorRegister(r)
	HandleUpdateServer(r)
	HandleModifyError(r)
	HandleViewError(r)
	HandleListError(r)

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

func Read() model.Server {
	db := database.DBConn()
	sql := makeQuery()

	rows, err := db.Query(sql)

	if nil != err {
		panic(err)
	}

	for rows.Next() {
		var id, dcName, rackName, ustartName, uendName, numDisk, portType, serialNumber, serverStatus, maker string
		err = rows.Scan(&id, &dcName, &rackName, &ustartName, &uendName, &numDisk, &portType, &serialNumber, &serverStatus, &maker)


		if nil != err {
			panic(err.Error())
		}

		dc := model.DataCenter {
			Name: dcName,
		}

		rack := model.Rack {
			Name: rackName,
		}

		ustart := model.RackUnit {
			Name: ustartName,
		}

		uend := model.RackUnit {
			Name: uendName,
		}

		pType := model.PortType {
			Name: portType,
		}

		location := model.Location {
			DataCenter: dc,
			Rack: rack,
			UStart: ustart,
			UEnd: uend,
		}

		server := model.Server {
			Id: id,
			Location: &location,
			SerialNumber: serialNumber,
			PortType: &pType,
		}

		return server
	}

	defer db.Close()
	return model.Server{}
}

func makeQuery() string {
	return "select SERVER.id, DC.name, RACK.name, " +
					"ustart.name, uend.name, num_disks, " +
					"PORT_TYPE.name, serial_number, " +
					"SERVER_STATUS.status, SERVER.maker " + 
			"from SERVER, DC, RACK, " +
				"RACK_UNIT as ustart, RACK_UNIT as uend, PORT_TYPE, " + 
				"SERVER_STATUS, STATUS_ROW " +
			"where SERVER.id_DC = DC.id and " + 
				"SERVER.id_Rack = RACK.id and " + 
				"SERVER.id_U_start = ustart.id and " + 
				"SERVER.id_U_end = uend.id and " + 
				"SERVER.id_PORT_TYPE = PORT_TYPE.id and " + 
				"SERVER.id_SERVER_STATUS = SERVER_STATUS.id and " + 
				"STATUS_ROW.status = 'available' and SERVER.id_STATUS_ROW = STATUS_ROW.id;";
}

type Login struct {
	User string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func makeInfoPage() page.Information {
	return page.Information {
		makeServer(),
		makeSwitch(),
	}
}

func makeLocation() model.Location {
	DC := model.DataCenter{"123hfjdfa", "Data Center 01"}
	Rack := model.Rack{"je383jds324", "F500"}
	Ustart := model.RackUnit{"klajf", "U30"}
	Uend := model.RackUnit{"klajf", "U31"}
	return model.Location{
		DC,
		Rack,
		Ustart,
		Uend,
	}
}

func makeServer() model.Server {
	// loc := makeLocation()
	// return model.Server {
	// 	"SV001ACD",
	// 		&loc,
	// 		"192.168.99.12",
	// 		"Notification",
	// 		"fj32hgf",
	// 		&model.PortType{"a", "idrac"},
	// }
	return Read()
}

func makeSwitch() model.Switch {
	return model.Switch {
			"SWdds22112",
			"Switch 01",
			makeLocation(),
			"192.168.99.12",
			50,
		}
}

// func makeServerItems() []page.ServerItem {
// 	item := page.ServerItem {
// 		"fhadsjk235",
// 		"S3 Gateway",
// 		makeLocation(),
// 		"2324kj431",
// 		"fine",
// 	}

// 	items := make([]page.ServerItem, 0)

// 	for i := 0; i < 12; i++ {
// 		items = append(items, item)
// 	}

// 	return items
// }

func HandleHome(router *gin.Engine,) {
	router.GET("/home", func (c *gin.Context) {
		home := page.Home{}
		home.Init()
		router.LoadHTMLFiles("templates/home.html")
		c.HTML(http.StatusOK, "templates/home.html", gin.H {
			"info": home.DCs,
		})
	})
}

func HandleLogin(r *gin.Engine) {
	r.GET("/user/login", func (c *gin.Context) {
		r.LoadHTMLFiles("templates/user/login.html")
		c.HTML(http.StatusOK, "templates/user/login.html", nil)
	})
}

func HandleRegisterServer(r *gin.Engine, server *model.Server) {
	r.GET("/server/register", func (c *gin.Context) {
		r.LoadHTMLFiles("templates/server/register.html")
		c.HTML(http.StatusOK, "templates/server/register.html", nil)
	})
}

func HandleInfo(r *gin.Engine, info *page.Information) {
	r.POST("/server/information", func (c *gin.Context) {
		//id := c.PostForm("id")
		r.LoadHTMLFiles("templates/server/information.html")
		c.HTML(http.StatusOK, "templates/server/information.html", info)
	})
}

func HandleSummary(r *gin.Engine) {
	r.GET("/server/summary", func (c *gin.Context) {
		r.LoadHTMLFiles("templates/server/summary.html")
		c.HTML(http.StatusOK, "templates/server/summary.html", nil)
	})
}

func HandleServers(r *gin.Engine) {
	r.POST("/server/list", func (c *gin.Context) {
		DC := model.DataCenter{
			c.PostForm("txtId"),
			c.PostForm("txtName"),
		}

		sv := page.Servers{DC, nil}
		sv.Init(DC.Id)

		r.LoadHTMLFiles("templates/server/list.html")
		c.HTML(http.StatusOK, "templates/server/list.html", sv)
	})
}

func HandleUpdateServer(r *gin.Engine) {
	r.POST("/server/modify", func (c *gin.Context) {
		//id := c.PostForm("id")
		update_page := page.UpdateServer{makeLocation(), makeServer()}
		r.LoadHTMLFiles("templates/server/modify.html")
		c.HTML(http.StatusOK, "templates/server/modify.html", update_page)
	})
}

// Error modify testing area
func makeError() model.Error {
	status := makeErrorStatus()
	return model.Error{
		"error1",
		"summary",
		"describe",
		"solution",
		time.Now(),
		"Server1",
		&status,
	}
}

func makeErrorStatus() model.ErrorStatus {
	return model.ErrorStatus{
		"status1",
		"Active",
	}
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

func HandleViewError(r *gin.Engine) {
	r.POST("/error/view", func (c *gin.Context) {
		r.LoadHTMLFiles("templates/error/view.html")
		c.HTML(http.StatusOK, "templates/error/view.html", makeError())
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
}

func HandleListError(r *gin.Engine) {
	r.GET("/error/list", func (c *gin.Context) {
		r.LoadHTMLFiles("templates/error/list.html")
		c.HTML(http.StatusOK, "templates/error/list.html", makeListErrorPage())
	})
}
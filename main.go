package main

import (
	"CURD/entity"
	"CURD/entity/error_entity"
	"CURD/page"
	"CURD/page/error_page"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	_ "html/template"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysessions", store))	

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

	Login(r)
	Logout(r)
	Authen(r)
	HandleHome(r)
	HandleFilter(r)
	HandleSearch(r)
	HandleSearchTags(r)
	HandleServers(r)
	HandleInfo(r)	
	HandleUpdateServer(r)
	HandleEditServices(r)
	HandleExecuteEditServices(r)
	HandleExecuteModify(r)
	HandleRegisterServer(r)
	HandleExecuteRegister(r)
	HandleErrorRegister(r)
	HandleListError(r)
	HandleViewError(r)
	HandleErrorExecuteUpdate(r)

	HandleRegisterSwitch(r)
	HandleExecuteRegisterSwitch(r)

	HandleRegistrationIp(r)
	HandleExecuteRegisterIp(r)
	HandleViewIp(r)
	HandleUpdateStateIp(r)
	HandleSearchIp(r)
	HandleListIpNet(r)

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

func Login(router *gin.Engine) {
	router.GET("/login", func (c *gin.Context) {
		router.LoadHTMLFiles("templates/user/login.html")
		c.HTML(http.StatusOK, "templates/user/login.html", nil)
	})
}

func Logout(router *gin.Engine) {
	router.GET("/logout", func (c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.Redirect(http.StatusFound, "/login")
	})
}

func Authen(router *gin.Engine) {
	router.POST("/auth", func (c *gin.Context) {
		user := entity.User {
			Username: c.PostForm("username"),
			Password: c.PostForm("password"),
		}
		authen := page.Authen{User: user}
		isValid := authen.IsExistsUser()

		if isValid {
			session := sessions.Default(c)
			session.Set("id", user.Username)
			session.Save()
			c.Redirect(http.StatusFound, "/")
		} else {
			c.Redirect(http.StatusFound, "/login")
		}
	})
}

func CheckAuthen(c *gin.Context) {
	if session := sessions.Default(c); session.Get("id") == nil {
			c.Redirect(http.StatusFound, "/login")
	}
}

func HandleHome(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		CheckAuthen(c)
		home := page.Home{}
		home.New()
		router.LoadHTMLFiles("templates/AdminLTE/index.html")
		c.HTML(http.StatusOK, "templates/AdminLTE/index.html", home)
	})
}

func HandleFilter(router *gin.Engine) {
	router.GET("/filter", func (c *gin.Context) {
		CheckAuthen(c)
		tag := entity.Tag {
			TagId: c.Query("txtTagId"),
			Title: c.Query("txtTagTitle"),
		}

		filterPage := page.Filter{}
		filterPage.New(tag)

		router.LoadHTMLFiles("templates/server/filter.html")
		c.HTML(http.StatusOK, "templates/server/filter.html", filterPage)
	})
}

func HandleSearch(router *gin.Engine) {
	router.POST("/search", func (c *gin.Context) {
		CheckAuthen(c)
		ip := c.PostForm("txtIpAddr")

		infoPage := page.Information{}
		infoPage.Prepare(ip)

		if "" != infoPage.Server.Id {
			router.LoadHTMLFiles("templates/server/information.html")
			c.HTML(http.StatusOK, "templates/server/information.html", infoPage)
		} else {
			router.LoadHTMLFiles("templates/server/not_found_server.html")
			c.HTML(http.StatusOK, "templates/server/not_found_server.html", nil)
		}
	})
}

func HandleSearchTags(router *gin.Engine) {
	router.POST("/tags", func (c *gin.Context) {
		CheckAuthen(c)
		values := c.PostForm("txtTags")
		tags := strings.Split(values, ",")

		f := page.Filter{}
		f.SearchServersByMultiTags(tags)

		router.LoadHTMLFiles("templates/server/filter.html")
		c.HTML(http.StatusOK, "templates/server/filter.html", f)
	})
}

func HandleServers(r *gin.Engine) {
	r.POST("/server/list", func(c *gin.Context) {
		CheckAuthen(c)
		tagId := c.PostForm("txtTagId")
		DC := entity.DataCenter{
			c.PostForm("txtId"),
			c.PostForm("txtName"),
		}

		server := page.Servers{DC, nil, nil}

		if "" != tagId {
			server.GetServersByTagId(tagId)
		} else {
			server.New(DC.Id)
		}

		r.LoadHTMLFiles("templates/server/list.html")
		c.HTML(http.StatusOK, "templates/server/list.html", server)
	})
}

func HandleInfo(r *gin.Engine) {
	r.POST("/server/information", func(c *gin.Context) {
		CheckAuthen(c)
		idServer := c.PostForm("txtIdServer")

		var infoPage page.Information
		infoPage.New(idServer)

		r.LoadHTMLFiles("templates/server/information.html")
		c.HTML(http.StatusOK, "templates/server/information.html", infoPage)
	})
}

func HandleUpdateServer(r *gin.Engine) {
	r.POST("/server/modify", func(c *gin.Context) {
		CheckAuthen(c)
		var updatePage page.UpdateServer
		updatePage.New(c)

		r.LoadHTMLFiles("templates/server/modify.html")
		c.HTML(http.StatusOK, "templates/server/modify.html", updatePage)
	})
}

func HandleEditServices(r *gin.Engine) {
	r.POST("/server/services", func (c *gin.Context) {
		CheckAuthen(c)
		var edit page.EditServices
		edit.New(c)

		r.LoadHTMLFiles("templates/server/services.html")
		c.HTML(http.StatusOK, "templates/server/services.html", edit)
	})
}

func HandleExecuteEditServices(r *gin.Engine) {
	r.POST("/server/execute_edit_services", func (c *gin.Context) {
		CheckAuthen(c)
		var exServices page.ExecuteEditServices
		exServices.Execute(c)

		r.LoadHTMLFiles("templates/server/execute_modify.html")
		c.HTML(http.StatusOK, "templates/server/execute_modify.html", exServices)
	})
}

func HandleExecuteModify(r *gin.Engine) {
	r.POST("server/execute_modify", func(c *gin.Context) {
		CheckAuthen(c)
		// server := getServerFromPostForm(c)
		var executeModify page.ExecuteModify
		executeModify.New(c)

		r.LoadHTMLFiles("templates/server/execute_modify.html")
		c.HTML(http.StatusOK,
			"templates/server/execute_modify.html",
			executeModify)
	})
}

func HandleRegisterServer(r *gin.Engine) {
	r.GET("/server/register", func (c *gin.Context) {
		var registration page.RegistrationServer
		registration.New()
		r.LoadHTMLFiles("templates/server/register.html")
		c.HTML(http.StatusOK, "templates/server/register.html", registration)
	})
}

func HandleRegisterSwitch(r *gin.Engine) {
	r.GET("/switch/register", func (c *gin.Context) {
		var registration page.RegistrationSwitch
		registration.New()
		r.LoadHTMLFiles("templates/switch/switch_register.html")
		c.HTML(http.StatusOK, "templates/switch/switch_register.html", registration)
	})
}

func HandleExecuteRegisterSwitch(r *gin.Engine) {
	r.POST("/switch/execute_register_switch", func (c *gin.Context) {
		var execution page.ExecuteRegisterSwitch
		err := execution.Execute(c)
		if nil != err {
			panic (err)
		}

		r.LoadHTMLFiles("templates/switch/execute_register_switch.html")
		c.HTML(http.StatusOK, "templates/switch/execute_register_switch.html", execution)
	})
}

func HandleListError(r *gin.Engine) {
	r.GET("/error/list", func (c *gin.Context) {
		CheckAuthen(c)
		var errorsPage error_page.Errors
		errorsPage.New()

		r.LoadHTMLFiles("templates/error/list.html")
		c.HTML(http.StatusOK, "templates/error/list.html", errorsPage)
	})
}

func HandleViewError(r *gin.Engine) {
	r.POST("/error/view", func (c *gin.Context) {
		CheckAuthen(c)
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
		CheckAuthen(c)
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

func HandleErrorRegister(r *gin.Engine) {
	r.GET("/error/register", func (c *gin.Context) {
		r.LoadHTMLFiles("templates/error/register.html")
		c.HTML(http.StatusOK, "templates/error/register.html", nil)
	})
}

func HandleExecuteRegister(r *gin.Engine) {
	r.POST("/server/execute_register", func (c *gin.Context) {
		CheckAuthen(c)
		var registrationPage page.ExecuteRegister
		registrationPage.New(c)
		r.LoadHTMLFiles("templates/server/execute_register.html")
		c.HTML(http.StatusOK, "templates/server/execute_register.html", registrationPage)
	})
}

func HandleRegistrationIp(r *gin.Engine) {
	r.GET("/server/register_ip", func (c *gin.Context) {
		CheckAuthen(c)

		r.LoadHTMLFiles("templates/server/register_ip.html")
		c.HTML(http.StatusOK, "templates/server/register_ip.html", nil)
	})
}

func HandleExecuteRegisterIp(r *gin.Engine) {
	r.POST("/server/execute_register_ip", func (c *gin.Context) {
		CheckAuthen(c)

		ex := page.ExecuteRegisterIp{}
		ex.New(c)

		r.LoadHTMLFiles("templates/server/execute_modify.html")
		c.HTML(http.StatusOK, "templates/server/execute_modify.html", ex)
	})
}

func HandleViewIp(r *gin.Engine) {
	r.GET("/server/list_ip", func (c *gin.Context) {
		CheckAuthen(c)

		listIp := page.ListIp{}
		listIp.New(c)
		//
		//r.LoadHTMLFiles("templates/server/list_ip.html")
		//c.HTML(http.StatusOK, "templates/server/list_ip.html", listIp)
		c.JSON(http.StatusOK, listIp.IpArr)
	})
}

func HandleUpdateStateIp(r *gin.Engine) {
	r.POST("/server/update_state_ip", func (c *gin.Context) {
		CheckAuthen(c)

		updIp := page.UpdateStateIp {}
		updIp.New(c)

		r.LoadHTMLFiles("templates/server/execute_modify.html")
		c.HTML(http.StatusOK, "templates/server/execute_modify.html", updIp)
	})
}

func HandleSearchIp(r *gin.Engine) {
	r.GET("/server/search_ip", func (c *gin.Context) {
		CheckAuthen(c)

		ip := page.SearchIp{}
		ip.New(c)
	})
}

func HandleListIpNet(r *gin.Engine) {
	r.GET("/server/ip", func (c *gin.Context) {
		CheckAuthen(c)

		netView := page.ListIpNet{}
		netView.New()

		r.LoadHTMLFiles("templates/server/list_ip_net.html")
		c.HTML(http.StatusOK, "templates/server/list_ip_net.html", netView)
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

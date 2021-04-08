package main

import (
	"CURD/entity"
	"CURD/entity/error_entity"
	"CURD/page"
	"CURD/page/error_page"
	"CURD/repo/server"
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


	Login(r)
	Logout(r)
	Authen(r)

	// Server routing Area
	HandleHome(r)
	//HandleFilter(r)
	HandleSearch(r)
	HandleSearchTags(r)
	HandleServers(r)
	//HandleInfo(r)
	HandleEditServer(r)
	//HandleEditServices(r)
	//HandleExecuteEditServices(r)
	HandleExecuteModify(r)
	HandleRegisterServer(r)
	HandleExecuteRegister(r)
	// End of server routing area

	// Error routing area
	HandleErrorRegister(r)
	HandleListError(r)
	HandleViewError(r)
	HandleErrorExecuteUpdate(r)
	// end of error routing area

	// Switch routing area
	HandleRegisterSwitch(r)
	HandleExecuteRegisterSwitch(r)
	// end of switch routing area

	// Ip Management routing area
	HandleRegistrationIp(r)
	HandleExecuteRegisterIp(r)
	HandleViewIp(r)
	HandleUpdateStateIp(r)
	HandleSearchIp(r)
	HandleListIpNet(r)
	HandleDeleteIpNet(r)
	// end of IpManagement routing area

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
		err := home.New()

		if nil != err {
			panic (err)
		}

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
		ip := c.PostForm("txtIp")

		var resultPage page.Servers
		if err := resultPage.FetchServersByIpAddress(ip); nil != err {
			c.String(http.StatusOK, err.Error())
			return
		}

		router.LoadHTMLFiles("templates/server/list.html")
		c.HTML(http.StatusOK, "templates/server/list.html", resultPage)
	})
}

func HandleSearchTags(router *gin.Engine) {
	router.POST("/tags", func (c *gin.Context) {
		CheckAuthen(c)
		values := c.PostForm("txtTags")
		tags := strings.Split(values, ",")

		f := page.Filter{}
		err := f.SearchServersByMultiTags(tags)
		if nil != err {
			c.String(http.StatusOK, err.Error())
			return
		}

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
		var err error

		if "" != tagId {
			err = server.FetchServerByTagId(tagId)
		} else {
			err = server.FetchServerByDCId(DC.Id)
		}

		if nil != err {
			panic (err)
		}

		r.LoadHTMLFiles("templates/server/list.html")
		c.HTML(http.StatusOK, "templates/server/list.html", server)
	})
}
//
//func HandleInfo(r *gin.Engine) {
//	r.POST("/server/information", func(c *gin.Context) {
//		CheckAuthen(c)
//		idServer := c.PostForm("txtIdServer")
//
//		var infoPage page.Information
//		infoPage.New(idServer)
//
//		r.LoadHTMLFiles("templates/server/information.html")
//		c.HTML(http.StatusOK, "templates/server/information.html", infoPage)
//	})
//}

func HandleDeleteServer(r *gin.Engine) {
	r.POST("/server/delete", func (c *gin.Context) {
		CheckAuthen(c)

		deletionPage := page.ExecuteDeleteServer{}
		msg := deletionPage.ExecuteDelete(c)

		c.String(http.StatusOK, msg)
	})
}

func HandleEditServer(r *gin.Engine) {
	r.POST("/server/edit", func(c *gin.Context) {
		CheckAuthen(c)

		serverId := c.PostForm("txtIdServer")
		var updatePage page.UpdateServer
		err := updatePage.New(serverId)
		if nil != err {
			panic (err)
		}

		r.LoadHTMLFiles("templates/server/edit.html")
		c.HTML(http.StatusOK, "templates/server/edit.html", updatePage)
	})
}

//func HandleEditServices(r *gin.Engine) {
//	r.POST("/server/services", func (c *gin.Context) {
//		CheckAuthen(c)
//		var edit page.EditServices
//		edit.New(c)
//
//		r.LoadHTMLFiles("templates/server/services.html")
//		c.HTML(http.StatusOK, "templates/server/services.html", edit)
//	})
//}

//func HandleExecuteEditServices(r *gin.Engine) {
//	r.POST("/server/execute_edit_services", func (c *gin.Context) {
//		CheckAuthen(c)
//		var exServices page.ExecuteEditServices
//		exServices.Execute(c)
//
//		r.LoadHTMLFiles("templates/server/execute_modify.html")
//		c.HTML(http.StatusOK, "templates/server/execute_modify.html", exServices)
//	})
//}

func HandleExecuteModify(r *gin.Engine) {
	r.POST("server/execute_modify", func(c *gin.Context) {
		CheckAuthen(c)
		// server := getServerFromPostForm(c)
		var executeModify page.ExecuteModify
		executeModify.New(c)
		if err := executeModify.Execute(); nil != err {
			panic (err)
		}

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
		var err error
		if err = registrationPage.New(c); nil != err {
			c.String(http.StatusOK, err.Error())
		}

		if err = registrationPage.Execute(); nil != err {
			c.String(http.StatusOK, err.Error())
		}
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

		regis := page.NetworkPortionRegistration{}
		regis.SetMsg("success")
		if err := regis.New(c); nil != err {
			regis.SetMsg(err.Error())
		}

		if err := regis.Execute(); nil != err {
			regis.SetMsg(err.Error())
		}

		c.JSON(http.StatusOK, regis)
	})
}

func HandleViewIp(r *gin.Engine) {
	r.GET("/server/list_ip", func (c *gin.Context) {
		CheckAuthen(c)

		listIp := page.ListIp{}
		if err := listIp.New(c); nil != err {
			c.String(http.StatusOK, err.Error())
			return
		}

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
		err := ip.New(c)

		if nil != err {
			c.String(http.StatusOK, err.Error())
			return
		}

		if "" != ip.IpState {
			c.String(http.StatusOK, ip.IpState)
		} else {
			c.String(http.StatusOK, "Not found IP")
		}
	})
}

func HandleListIpNet(r *gin.Engine) {
	r.GET("/server/ip", func (c *gin.Context) {
		CheckAuthen(c)

		netView := page.ListNetworkPortions{}
		if err := netView.New(); nil != err {
			c.String(http.StatusOK, err.Error())
			return
		}

		r.LoadHTMLFiles("templates/server/list_ip_net.html")
		c.HTML(http.StatusOK, "templates/server/list_ip_net.html", netView)
	})
}

func HandleDeleteIpNet(r *gin.Engine) {
	r.POST("/server/delete_ip", func (c *gin.Context) {
		CheckAuthen(c)

		netId := c.PostForm("txtNetId")
		repo := server.NetworkPortionRepo{}
		msg :="success"
		if err := repo.Delete(netId); nil != err {
			msg = err.Error()
		}

		c.String(http.StatusOK, msg)
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

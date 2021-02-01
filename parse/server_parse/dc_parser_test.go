package server_parse

import (
	"CURD/entity"
	"testing"
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
)

var e entity.Entity

func init() {
	r := gin.Default()
	r.Static("/public", "./public")

	r.GET("/dc", func (c *gin.Context) {
		parser := DCParser{}
		parser.SetId("txtDCId")
		parser.SetDescription("txtDCDes")
		e = parser.ParseQuery(c)
	})

	r.GET("/rack", func (c *gin.Context) {
		parser := RackParser{}
		parser.SetId("txtRackId")
		parser.SetDescription("txtRackDes")
		e = parser.ParseQuery(c)
	})

	r.GET("/unit", func (c *gin.Context) {
		parser := RackUnitParser{}
		parser.SetId("txtUStartId")
		parser.SetDescription("txtUEndId")
		e = parser.ParseQuery(c)
	})

	r.GET("/porttype", func (c *gin.Context) {
		parser := PortTypeParser{}
		parser.SetId("txtPTId")
		parser.SetDescription("txtPTDes")
		e = parser.ParseQuery(c)
	})

	go r.Run(":9000")
}

func TestParseDC(t *testing.T) {
	time.Sleep (2 * time.Second)
	http.Get("http://localhost:9000/dc?txtDCId=1&txtDCDes=2")
	dc := e.(entity.DataCenter)

	if ("1" != dc.Id && "2" != dc.Description) {
		t.Error ("Fail")
	}
}
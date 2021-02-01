package server_parse

import (
	"CURD/entity"
	"github.com/gin-gonic/gin"
)

type PortTypeParser struct {
	Id, Description string
}

func (parser *PortTypeParser) SetId(Id string) {
	parser.Id = Id
}

func (parser *PortTypeParser) SetDescription(Description string) {
	parser.Description = Description
}

func (parser PortTypeParser) ParsePostForm(context *gin.Context) entity.PortType {
	Id := context.PostForm(parser.Id)
	Des := context.PostForm(parser.Description)
	return entity.PortType{Id, Des}
}

func (parser PortTypeParser) ParseQuery(context *gin.Context) entity.PortType {
	Id := context.Query(parser.Id)
	Des := context.Query(parser.Description)
	return entity.PortType{Id, Des}
}
package server_parse

import (
	"CURD/entity"
	"github.com/gin-gonic/gin"
)

type RackUnitParser struct{
	Id, Description string
}

func (parser *RackUnitParser) SetId(Id string) {
	parser.Id = Id
}

func (parser *RackUnitParser) SetDescription(Des string) {
	parser.Description = Des
}

func (parser RackUnitParser) ParsePostForm(context *gin.Context) entity.RackUnit {
	Id := context.PostForm(parser.Id)
	Des := context.PostForm(parser.Description)
	return entity.RackUnit{Id, Des}
}

func (parser RackUnitParser) ParseQuery(context *gin.Context) entity.RackUnit {
	Id := context.Query(parser.Id)
	Des:= context.Query(parser.Description)
	return entity.RackUnit{Id, Des}
}
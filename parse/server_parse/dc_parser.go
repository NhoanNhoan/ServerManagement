package server_parse

import (
	"CURD/entity"
	"github.com/gin-gonic/gin"
)

type DCParser struct {
	Id, Description string
}

func (parser *DCParser) SetId(Id string) {
	parser.Id = Id
}

func (parser *DCParser) SetDescription(Description string) {
	parser.Description = Description
}

func (parser *DCParser) ParseQuery(context *gin.Context) entity.Entity {
	Id := context.Query(parser.Id)
	Des := context.Query(parser.Description)
	return entity.DataCenter{Id, Des}
}

func (parser DCParser) ParsePostForm(context *gin.Context) entity.Entity {
	Id := context.PostForm(parser.Id)
	Des := context.PostForm(parser.Description)
	return entity.DataCenter{Id, Des}
}
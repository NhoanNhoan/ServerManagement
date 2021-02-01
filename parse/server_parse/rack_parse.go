package server_parse

import (
	"CURD/entity"
	"github.com/gin-gonic/gin"
)

type RackParser struct {
	Id, Description string
}

func (parser *RackParser) SetId(Id string) {
	parser.Id = Id
}

func (parser *RackParser) SetDescription(Description string) {
	parser.Description = Description
}

func (parser RackParser) ParsePostForm(context *gin.Context) entity.Entity {
	Id := context.PostForm(parser.Id)
	Des := context.PostForm(parser.Description)
	return entity.Rack{Id, Des}
}

func (parser RackParser) ParseQuery(context *gin.Context) entity.Entity {
	Id := context.Query(parser.Id)
	Des := context.Query(parser.Description)
	return entity.Rack{Id, Des}
}
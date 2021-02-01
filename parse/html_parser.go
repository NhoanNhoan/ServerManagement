package parse

import (
	"CURD/entity"
	"github.com/gin-gonic/gin"
)

type HtmlParser interface {
	ParsePostForm(context *gin.Context) entity.Entity
	ParseGetForm(context *gin.Context) entity.Entity
}
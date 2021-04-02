package entity

import (
	"github.com/gin-gonic/gin"
)

type Entity interface {
	MakeByContext(c *gin.Context)
}


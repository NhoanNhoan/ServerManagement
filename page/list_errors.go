package page

import (
	"CURD/model"
)

type ListError struct {
	ServerName string
	Errors []model.Error
}
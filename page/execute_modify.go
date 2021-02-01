package page

import (
	"CURD/entity"
)

type ExecuteModify struct {
	Msg string
}

func (obj *ExecuteModify) New(server entity.Server) {
	obj.Msg = entity.UpdateServer(server)
}

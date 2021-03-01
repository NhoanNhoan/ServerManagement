package page

import (
	"CURD/entity"
)

type Authen struct {
	Msg string
	entity.User
}

func (authen *Authen) IsExistsUser() bool {
	if !authen.User.IsExists() {
		authen.Msg = "Username or password is invalid"
		return false
	}

	return true
}
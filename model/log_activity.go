package model

import "time"

type LogActivity struct {
	Id string
	*User
	Operation string
	Affected string
	At time.Time
}
package model

import (
	"time"
)

type ServerEvent struct {
	Id string
	IdServer string
	Description string
	OccurAt time.Time
}
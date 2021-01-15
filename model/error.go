package model

import (
	"time"
)

type Error struct {
	Id string
	Summary string
	Describe string
	Solution string
	LastOccur time.Time
	IdServer string
	*ErrorStatus
}
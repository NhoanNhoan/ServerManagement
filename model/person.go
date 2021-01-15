package model

import (
	"time"
)

type Person struct {
	Id string
	Name string
	BirthDate time.Time
	Team string
	OtherInfo string
}
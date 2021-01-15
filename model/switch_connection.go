package model

type SwitchConnection struct {
	Id string
	*Server
	*Switch
	ServerPort, SwitchPort int
	*CableType
}

type CableType struct {
	Id string
	Name string
	SignPort string
}
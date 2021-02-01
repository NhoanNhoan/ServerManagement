package model

type Location struct {
	DC DataCenter
	Rack
	Ustart RackUnit
	Uend RackUnit
}

type DataCenter struct {
	Id string
	Description string
}

type Rack struct {
	Id string
	Description string
}

type RackUnit struct {
	Id string
	Description string
}
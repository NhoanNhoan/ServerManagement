package model

type Location struct {
	DataCenter
	Rack
	UStart RackUnit
	UEnd RackUnit
}

type DataCenter struct {
	Id string
	Name string
}

type Rack struct {
	Id string
	Name string
}

type RackUnit struct {
	Id string
	Name string
}

func (loc *Location) String() string {
	return loc.DataCenter.Name + " - " + loc.Rack.Name +
	" - [" + loc.UStart.Name + ", " + loc.UEnd.Name + "]"
}
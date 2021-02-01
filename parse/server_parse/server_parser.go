package server_parse

import (
	"CURD/entity"
	"github.com/gin-gonic/gin"
)

const (
	SERVER_ID = "txtServerId"
	SERVER_NUM_DISK = "txtNumDisks"
	SERVER_MAKER = "txtMaker"
)

type ServerParser struct {
	DCParser
	RackParser
	RackUnitParser
	PortTypeParser
}

func (parser *ServerParser) SetDCParser(Id, Des string) {
	parser.DCParser.Id = Id
	parser.DCParser.Des = Des
}

func (parser *ServerParser) SetRackParser(Id, Des string) {
	parser.RackParser.Id = Id
	parser.RackParser.Des = Des
}

func (parser *RackUnitParser) SetRackParser(Id, Des string) {
	parser.RackUnitParser.Id = Id
	parser.RackUnitParser.Des = Des
}

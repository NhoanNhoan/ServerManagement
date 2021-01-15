package page

type Page struct {
	queries []string
	MakeQuery() string
	Init()
}

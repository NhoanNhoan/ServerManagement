package server

type EntityRepo interface {
	FetchAll() ([]interface{}, error)
	FetchById(id string) (interface{}, error)
	IsExists(id string) bool
	GenerateId() string
	Insert(... interface{}) error
	Update(... interface{}) error
	Delete(... interface{}) error
}


package server

import (
	"CURD/database"
	"CURD/entity"
	"database/sql"
)

type qcomp = database.QueryComponent
type icomp = database.InsertComponent
type ucomp = database.UpdateComponent
type dcomp = database.DeleteComponent

type Repository interface {
	Insert(icomp) error
	Update(ucomp) error
	Delete(dcomp) error
	Query(qcomp,
		scan func(interface{}, *sql.Rows) (interface{}, error)) ([]entity.Entity, error)
}
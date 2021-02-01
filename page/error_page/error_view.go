package error_page

import (
	"CURD/entity/error_entity"
)

type ErrorView struct {
	error_entity.Error
	Curators []error_entity.Curator
	AllPersons []error_entity.Person
	AllErrorStates []error_entity.ErrorState
}

func (obj *ErrorView) New(errData error_entity.Error) error {
	obj.SetError(errData)
	obj.AllPersons = error_entity.FetchPersons()

	obj.AllErrorStates = error_entity.FetchErrorStates()
	err := obj.fetchCurators()

	return err
}

func (obj *ErrorView) SetError(errData error_entity.Error) {
	obj.Error = errData
}

func (view *ErrorView) fetchCurators() error {
	var err error
	view.Curators, err = error_entity.FetchCurators(view.Error.Id)
	return err
}
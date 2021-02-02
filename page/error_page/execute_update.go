package error_page

import (
	"CURD/entity"
	"CURD/entity/error_entity"
	"github.com/gin-gonic/gin"
)

type ExecuteUpdateError struct {
	AllPersons []error_entity.Person
	error_entity.Error
	Curators []error_entity.Curator
	Msg string
}

func (obj *ExecuteUpdateError) New(context *gin.Context) {
	obj.makeErrorByContext(context)
	obj.AllPersons = error_entity.FetchPersons()
	obj.makeCuratorsByContext(context)
}

func (obj *ExecuteUpdateError) makeErrorByContext(context *gin.Context) {
	id := context.PostForm("txtErrorId")
	idServer := context.PostForm("txtServerId")
	summary := context.PostForm("txtSummary")
	des := context.PostForm("txtDescription")
	solution := context.PostForm("txtSolution")
	stateId := context.PostForm("cbStateId")

	obj.Error = error_entity.Error {
		Id: id,
		Summary: summary,
		Description: des,
		Solution: solution,
		Occurs: "",
		Server: entity.Server {Id: idServer},
		ErrorState: error_entity.ErrorState {Id: stateId},
	}
}

func (obj *ExecuteUpdateError) makeCuratorsByContext(context *gin.Context) {
	for i := range obj.AllPersons {
		checked := context.PostForm(obj.AllPersons[i].Id)
		if "yes" == checked {
			obj.Curators = append(obj.Curators, error_entity.Curator {
				IdError: obj.Error.Id,
				IdPerson: obj.AllPersons[i].Id,
			})
		}
	}
}

func (obj *ExecuteUpdateError) updateCurators() error {
	var err error

	for i := range obj.Curators {
		err = obj.Curators[i].Execute()
		if nil != err {
			return err
		}
	}

	return err
}

func (obj *ExecuteUpdateError) Execute() (err error) {
	err = error_entity.DeleteCuratorsByErrorId(obj.Error.Id)
	if nil != err {
		panic (err)
	}

	err = obj.Error.Update()

	if nil != err {
		obj.Msg = "Can't update"
		panic (err)
	}

	err = obj.updateCurators()

	if nil != err {
		obj.Msg = "Can't update"
		panic (err)
	}

	obj.Msg = "Success"
	return nil
}

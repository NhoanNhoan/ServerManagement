package error_page

import (
	"CURD/entity"
	"CURD/entity/error_entity"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ExecuteUpdateError struct {
	error_entity.Error
	Curators []error_entity.Curator
	Msg string
}

func (obj *ExecuteUpdateError) New(context *gin.Context) {
	obj.makeErrorByContext(context)
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
	numCurators, _ := strconv.Atoi(context.PostForm("txtNumCurators"))
	obj.Curators = make([]error_entity.Curator, numCurators)

	for i := 0; i < numCurators; i++ {
		curatorId := context.PostForm("txtCuratorId" + strconv.Itoa(i))
		errorId := obj.Error.Id
		personId := context.PostForm("cbCuratorId" + strconv.Itoa(i))
		obj.Curators[i] = error_entity.Curator {
			Id: curatorId,
			IdError: errorId,
			IdPerson: personId,
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

func (obj *ExecuteUpdateError) Execute() error {
	err := obj.Error.Update()

	if nil != err {
		obj.Msg = "Can't update"
		panic (err)
	}

	obj.updateCurators()

	obj.Msg = "Success"
	return nil
}

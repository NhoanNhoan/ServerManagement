package error_entity

import (
	"testing"
)

func TestNewError(t *testing.T) {
	id := "ERR001"
	var obj Error

	err := obj.New(id)
	
	if nil != err {
		panic (err)
	}
}

func TestExistsComp(t *testing.T) {
	obj := Error {Id: "ERR001"}
	if !obj.isExists() {
		t.Error ("Fail")
	}
}

func TestUpdateError(t *testing.T) {
	errInstance := Error {
		Id: "ERR001",
		Summary: "Non Summary",
	}

	err := errInstance.Update()
	if nil != err {
		t.Error ("Fail")
	}
}
package error_entity

import (
	"testing"
)

func TestNewPerson(t *testing.T) {
	idPerson := "PS00001"
	var p Person
	err := p.New(idPerson)
	
	if nil != err {
		t.Error ("Fail")
	}
}

func TestFetchAllPersons(t *testing.T) {
	persons := FetchPersons()
	success := (0 != len(persons))

	if !success {
		t.Error ("Fail")
	}
}
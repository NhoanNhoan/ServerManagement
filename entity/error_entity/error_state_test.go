package error_entity

import (
	"testing"
)

func TestNewErrorState(t *testing.T) {
	var state ErrorState
	err := state.New("ES0001")
	success := (nil == err)

	if !success {
		t.Error("Fail")
	}
}

func TestFetchErrorStates(t *testing.T) {
	states := FetchErrorStates()
	success := (0 != len(states))

	if !success {
		t.Error ("Fail")
	}
}
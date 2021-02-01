package error_page

import (
	"testing"
)

func TestNewErrors(t *testing.T) {
	var instance Errors
	err := instance.New()

	if nil != err {
		t.Error ("Fail")
	}
}
package entity

import (
	"testing"
)

func TestFetchAllTags(t *testing.T) {
	tags := FetchAllTags()
	if 0 == len(tags) {
		t.Error("Fail")
	}
}
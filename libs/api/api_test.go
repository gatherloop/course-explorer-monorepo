package api

import (
	"testing"
)

func TestApi(t *testing.T) {
	result := Api("works")
	if result != "Api works" {
		t.Error("Expected Api to append 'works'")
	}
}

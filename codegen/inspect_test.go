package codegen

import (
	"testing"
)

func TestParsing(t *testing.T) {
	_, err := ParseTdproto()
	if err != nil {
		t.Error(err)
	}
}

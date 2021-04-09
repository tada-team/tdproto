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

func TestConverter(t *testing.T) {
	if SnakeCaseToLowerCamel("test_string") != "testString" {
		t.Fatal("Failed snake case converter")
	}
}

package codegen

import (
	"testing"
)

func TestParsing2(t *testing.T) {
	tdproto := make(TdProto2)

	err := ParseTdproto2(&tdproto)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConverter(t *testing.T) {
	if SnakeCaseToLowerCamel("test_string") != "testString" {
		t.Fatal("Failed snake case converter")
	}
}

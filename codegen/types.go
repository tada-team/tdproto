package codegen

var typesMap = map[string]struct {
	TypeScript string
	Dart       string
}{
	"string": {
		TypeScript: "string",
		Dart:       "String",
	},
	"int": {
		TypeScript: "number",
		Dart:       "int",
	},
	"int64": {
		TypeScript: "number",
		Dart:       "int",
	},
	"uint16": {
		TypeScript: "number",
		Dart:       "int",
	},
	"uint": {
		TypeScript: "number",
		Dart:       "int",
	},
	"bool": {
		TypeScript: "boolean",
		Dart:       "bool",
	},
	"interface{}": {
		TypeScript: "any",
		Dart:       "dynamic",
	},
	"ISODateTimeString": {
		TypeScript: "string",
		Dart:       "DateTime",
	},
}

var tsDefaultMap = map[string]string{
	"string":       `''`,
	"number":       `0`,
	"boolean":      `false`,
	"MessageLinks": `[]`,
	"Mediasubtype": `''`,
}

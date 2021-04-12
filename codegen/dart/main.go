package main

import (
	"os"
	"text/template"

	"github.com/tada-team/tdproto/codegen"
)

const dartHeader = `
import 'package:freezed_annotation/freezed_annotation.dart';`

const dartClassTemplate = `

`
const dartEnumHeader = `import 'package:freezed_annotation/freezed_annotation.dart';
`

var dartEnumTemplate = template.Must(template.New("dartEnum").Parse(`
enum {{.Name}} { {{ range $value :=  .Values}}
  @JsonValue('{{$value}}')
  {{$value}},
  {{end}}
}
`))

func generateDart(tdprotoInfo *codegen.TdInfo) {
	for _, tdEnum := range tdprotoInfo.GetEnums() {
		dartEnumTemplate.Execute(os.Stdout, tdEnum)
	}
}

func main() {
	tdprotoInfo, err := codegen.ParseTdproto()

	if err != nil {
		panic(err)
	}

	generateDart(tdprotoInfo.TdModels)
}

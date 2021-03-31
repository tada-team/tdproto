package main

import (
	"fmt"
	"github.com/tada-team/tdproto/codegen"
	"os"
	"text/template"
)

const markdownStructureTemplate = `
### {{.Name}}
{{.Help}}
{{range $field := .Fields}}
* **{{ $field.JsonName }}** (
    {{- if $field.IsPrimitive -}} {{- $field.TypeStr -}} {{ else }} [{{- $field.TypeStr -}}](#{{- $field.TypeStr -}}) {{ end }}
	{{- if $field.IsReadOnly }}, readonly for clients{{ end -}}
	{{- if $field.IsPointer }}, nullable{{ end -}}
	{{- if $field.IsList }}, list{{end -}}
	{{- if $field.IsOmitEmpty }}, omitempty{{ end -}}
  ) — {{ $field.Help }}.
{{end}}

`

func generateMarkdown(tdprotoInfo *codegen.TdInfo) {
	markdownTemplate, err := template.New("markdownStructure").Parse(markdownStructureTemplate)

	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stdout, "## Structures\n")

	for _, tdStructInfo := range tdprotoInfo.TdStructs {
		err := markdownTemplate.Execute(os.Stdout, tdStructInfo)
		if err != nil {
			panic(err)
		}
	}
}

func main() {

	tdprotoInfo, err := codegen.ParseTdproto()

	if err != nil {
		panic(err)
	}

	generateMarkdown(tdprotoInfo)
}

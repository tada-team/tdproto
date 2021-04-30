package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/tada-team/tdproto/codegen"
)

func main() {
	tdprotoInfo, err := codegen.ParseTdproto()
	if err != nil {
		panic(err)
	}

	if err := generateMarkdown(tdprotoInfo); err != nil {
		panic(err)
	}
}

type markdownEvent struct {
	Name           string
	Help           string
	Example        string
	EventStructStr string
}

type markdownStructField struct {
	codegen.TdStructField
}

type markdownStruct struct {
	codegen.TdStruct
	Fields []markdownStructField
}

func createMarkdownEvents(tdprotoInfo *codegen.TdInfo) (events []markdownEvent, err error) {
	for eventStructName, eventStr := range tdprotoInfo.TdEvents {
		eventExample, ok := eventExampleStr[eventStr]
		if !ok {
			eventExample = "EVENT MISSING EXAMPLE"
		}

		events = append(events, markdownEvent{
			Name:           eventStr,
			Help:           tdprotoInfo.TdStructs[eventStructName].Help,
			Example:        eventExample,
			EventStructStr: eventStructName,
		})
	}

	sort.Slice(events, func(i, j int) bool {
		return strings.ToLower(events[i].Name) < strings.ToLower(events[j].Name)
	})

	return events, nil
}

func generateMarkdown(tdprotoInfo *codegen.TdInfo) error {
	_, _ = fmt.Fprintln(os.Stdout, "## Structures")

	for _, tdStructInfo := range tdprotoInfo.TdStructs {
		if tdStructInfo.Help == "MISSING CLASS DOCUMENTATION" {
			continue
		}

		structureToPrint := markdownStruct{
			TdStruct: tdStructInfo,
			Fields:   make([]markdownStructField, 0),
		}

		allfields := make([]codegen.TdStructField, 0)
		allfields = append(allfields, tdStructInfo.Fields...)

		for _, anonStruct := range tdStructInfo.GetStructAnonymousStructs(tdprotoInfo) {
			allfields = append(allfields, anonStruct.Fields...)
		}

		for _, field := range allfields {
			if field.Help == "DOCUMENTATION MISSING" {
				continue
			}

			structureToPrint.Fields = append(structureToPrint.Fields, markdownStructField{
				TdStructField: field,
			})
		}

		if err := structureTemplate.Execute(os.Stdout, structureToPrint); err != nil {
			return err
		}
	}

	_, _ = fmt.Fprintln(os.Stdout, "## Events")

	markdownEvents, err := createMarkdownEvents(tdprotoInfo)
	if err != nil {
		return err
	}

	for _, event := range markdownEvents {
		if err := eventTemplate.Execute(os.Stdout, event); err != nil {
			return err
		}
	}
	return nil
}

var structureTemplate = template.Must(template.New("markdownStructure").Parse(`
### {{.TdStruct.Name}}
{{.TdStruct.Help}}
{{range $field := .Fields}}
* **{{ $field.TdStructField.JsonName }}** (
    {{- if $field.TdStructField.IsPrimitive -}} {{- $field.TdStructField.TypeStr -}} {{ else }} [{{- $field.TdStructField.TypeStr -}}](#{{- $field.TdStructField.TypeStr -}}) {{ end }}
	{{- if $field.TdStructField.IsReadOnly }}, readonly for clients{{ end -}}
	{{- if $field.TdStructField.IsPointer }}, nullable{{ end -}}
	{{- if $field.TdStructField.IsList }}, list{{end -}}
	{{- if $field.TdStructField.IsOmitEmpty }}, omitempty{{ end -}}
  ) — {{ $field.TdStructField.Help }}.{{end}}

`))

var eventTemplate = template.Must(template.New("markdownEvent").Parse(`
### "{{.Name}}"

Event structure: [{{.EventStructStr}}](#{{.EventStructStr}})

{{.Help}}

` +
	"```" +
	`
{{.Example}}
` +
	"```"))

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/tada-team/tdproto/codegen"
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

const markdownEventTemplate = `
### "{{.Name}}"

Event structure: [{{.EventStructStr}}](#{{.EventStructStr}})

{{.Help}}

` +
	"```" +
	`
{{.Example}}
` +
	"```"

type markdownEvent struct {
	Name           string
	Help           string
	Example        string
	EventStructStr string
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
	structureTemplate, err := template.New("markdownStructure").Parse(markdownStructureTemplate)
	if err != nil {
		return err
	}

	eventTemplate, err := template.New("markdownEvent").Parse(markdownEventTemplate)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, "## Structures")

	for _, tdStructInfo := range tdprotoInfo.TdStructs {
		for _, anonStruct := range tdStructInfo.GetStructAnonymousStructs(tdprotoInfo) {
			tdStructInfo.Fields = append(tdStructInfo.Fields, anonStruct.Fields...)
		}

		err := structureTemplate.Execute(os.Stdout, tdStructInfo)
		if err != nil {
			return err
		}
	}

	fmt.Fprintln(os.Stdout, "## Events")

	markdownEvents, err := createMarkdownEvents(tdprotoInfo)
	if err != nil {
		return err
	}

	for _, event := range markdownEvents {
		err := eventTemplate.Execute(os.Stdout, event)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {

	tdprotoInfo, err := codegen.ParseTdproto()
	if err != nil {
		panic(err)
	}

	generateMarkdown(tdprotoInfo)
}

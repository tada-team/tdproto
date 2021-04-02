package main

import (
	"fmt"
	"github.com/tada-team/tdproto/codegen"
	"os"
	"sort"
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

func createMarkdownEvents(tdprotoInfo *codegen.TdInfo) []markdownEvent {
	var events []markdownEvent

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
		return events[i].Name < events[j].Name
	})

	return events
}

func generateMarkdown(tdprotoInfo *codegen.TdInfo) {
	structureTemplate, err := template.New("markdownStructure").Parse(markdownStructureTemplate)

	if err != nil {
		panic(err)
	}

	eventTemplate, err := template.New("markdownEvent").Parse(markdownEventTemplate)

	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stdout, "## Structures\n")

	for _, tdStructInfo := range tdprotoInfo.TdStructs {
		err := structureTemplate.Execute(os.Stdout, tdStructInfo)
		if err != nil {
			panic(err)
		}
	}

	fmt.Fprint(os.Stdout, "## Events\n")

	markdownEvents := createMarkdownEvents(tdprotoInfo)

	for _, event := range markdownEvents {
		err := eventTemplate.Execute(os.Stdout, event)

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

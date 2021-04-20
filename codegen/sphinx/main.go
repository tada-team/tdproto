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

	rstEvents, err := generateRstServerEvents(tdprotoInfo)
	if err != nil {
		panic(err)
	}

	if err := printRstEvents(rstEvents); err != nil {
		panic(err)
	}
}

type rstEventField struct {
	Name string
	Help string
}

type rstEvent struct {
	Name           string
	Help           string
	Example        string
	EventStructStr string
	Fields         []rstEventField
}

func generateRstServerEvents(tdprotoInfo *codegen.TdInfo) (events []rstEvent, err error) {
	for eventStructName, eventStr := range tdprotoInfo.TdEvents {
		eventExample, ok := eventExampleStr[eventStr]
		if !ok {
			eventExample = "EVENT MISSING EXAMPLE"
		}

		originalStruct, ok := tdprotoInfo.TdStructs[eventStructName]
		if !ok {
			fmt.Fprintf(os.Stderr, "Failed to find struct %s of event %s.\n", eventStructName, eventStr)
			continue
		}

		eventExampleFormatted := strings.ReplaceAll(eventExample, "\n", "\n   ")

		event := rstEvent{
			Name:           eventStr,
			Help:           originalStruct.Help,
			Example:        eventExampleFormatted,
			EventStructStr: eventStructName,
		}

		var paramsStruct codegen.TdStruct

		for _, originalField := range originalStruct.Fields {
			if originalField.Name == "Params" {
				paramsStruct, ok = tdprotoInfo.TdStructs[originalField.TypeStr]
				if !ok {
					fmt.Fprintf(os.Stderr, "Failed to find parameter type %s of event %s.\n", originalField.TypeStr, eventStructName)
					continue
				}
			}
		}

		for _, paramField := range paramsStruct.Fields {
			event.Fields = append(event.Fields, rstEventField{
				Name: paramField.JsonName,
				Help: paramField.Help,
			})
		}

		for _, anonStruct := range paramsStruct.GetStructAnonymousStructs(tdprotoInfo) {
			for _, anonStructField := range anonStruct.Fields {
				event.Fields = append(event.Fields, rstEventField{
					Name: anonStructField.JsonName,
					Help: anonStructField.Help,
				})
			}

		}

		sort.Slice(event.Fields, func(i, j int) bool {
			return strings.ToLower(event.Fields[i].Name) < strings.ToLower(event.Fields[j].Name)
		})

		events = append(events, event)
	}

	sort.Slice(events, func(i, j int) bool {
		return strings.ToLower(events[i].Name) < strings.ToLower(events[j].Name)
	})

	return events, nil
}

var eventTemplate = template.Must(template.New("rstEvent").Parse(`
{{.Name}}
----------------------------------------------------------------------------

{{.Help}}

{{range $field := .Fields}}
* ` + "``" + "{{$field.Name}}" + "``" + ` - {{$field.Help}}
{{end}}


.. code-block:: json
   
   {{.Example}}

`))

func printRstEvents(events []rstEvent) error {
	_, err := fmt.Fprintln(os.Stdout, `Td Events
	======================================`)
	if err != nil {
		return err
	}

	for _, event := range events {
		err := eventTemplate.Execute(os.Stdout, event)
		if err != nil {
			return err
		}
	}

	return nil
}

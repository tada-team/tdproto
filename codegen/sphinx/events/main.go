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

	clientEvents, serverEvents, err := generateRstServerEvents(tdprotoInfo)
	if err != nil {
		panic(err)
	}

	if err := printRstEvents(clientEvents, serverEvents); err != nil {
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

func generateRstServerEvents(tdprotoInfo *codegen.TdProto) (clientEvents []rstEvent, serverEvents []rstEvent, err error) {
	for eventStructName, eventStr := range tdprotoInfo.TdEvents.TdEvents {
		eventExample, ok := eventExampleStr[eventStr]
		if !ok {
			eventExample = ""
		}

		originalStruct, ok := tdprotoInfo.TdEvents.TdStructs[eventStructName]
		if !ok {
			return clientEvents, serverEvents, fmt.Errorf("failed to find struct %s of event %s", eventStructName, eventStr)
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
			if originalField.Name != "Params" {
				continue
			}
			switch originalField.PackageName {
			case "tdevents":
				paramsStruct, ok = tdprotoInfo.TdEvents.TdStructs[originalField.TypeStr]
			case "tdproto":
				paramsStruct, ok = tdprotoInfo.TdModels.TdStructs[originalField.TypeStr]
			default:
				ok = false
			}

			if !ok {
				fmt.Fprintf(os.Stderr, "Failed to find struct %s of event %s.\n", eventStructName, eventStr)
				continue
			}

		}

		for _, paramField := range paramsStruct.GetAllJsonFields(tdprotoInfo.TdEvents) {
			fieldHelp := paramField.Help
			if fieldHelp == "" {
				fieldHelp = "DOCUMENTATION MISSING"
			}

			event.Fields = append(event.Fields, rstEventField{
				Name: paramField.JsonName,
				Help: fieldHelp,
			})
		}

		sort.Slice(event.Fields, func(i, j int) bool {
			return strings.ToLower(event.Fields[i].Name) < strings.ToLower(event.Fields[j].Name)
		})

		if strings.HasPrefix(event.Name, "client") {
			clientEvents = append(clientEvents, event)
		} else {
			serverEvents = append(serverEvents, event)
		}
	}

	sort.Slice(clientEvents, func(i, j int) bool {
		return strings.ToLower(clientEvents[i].Name) < strings.ToLower(clientEvents[j].Name)
	})

	sort.Slice(serverEvents, func(i, j int) bool {
		return strings.ToLower(serverEvents[i].Name) < strings.ToLower(serverEvents[j].Name)
	})

	return clientEvents, serverEvents, nil
}

var eventTemplate = template.Must(template.New("rstEvent").Parse(`
{{.Name}}
----------------------------------------------------------------------------

{{.Help}}
{{range $field := .Fields}}
* ` + "``" + "{{$field.Name}}" + "``" + ` - {{$field.Help}}{{end}}

{{if .Example -}}
.. code-block:: json
   
   {{.Example}}

{{else}}**MISSING EXAMPLE**
{{end}}`))

func printRstEvents(clientEvents []rstEvent, serverEvents []rstEvent) error {
	_, err := fmt.Fprintln(os.Stdout, `Client events
======================================`)
	if err != nil {
		return err
	}

	for _, event := range clientEvents {
		err := eventTemplate.Execute(os.Stdout, event)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintln(os.Stdout, `Server events
======================================`)
	if err != nil {
		return err
	}

	for _, event := range serverEvents {
		err := eventTemplate.Execute(os.Stdout, event)
		if err != nil {
			return err
		}
	}

	return nil
}

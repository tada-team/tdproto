package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"
	"unicode"

	"github.com/tada-team/tdproto/codegen"
)

func main() {
	tdprotoInfo, err := codegen.ParseTdproto()
	if err != nil {
		panic(err)
	}

	err = generateRstJson(tdprotoInfo)
	if err != nil {
		panic(err)
	}
}

var jsTypesMap = map[string]string{
	"string":            "string",
	"int":               "number",
	"int64":             "number",
	"uint16":            "number",
	"uint":              "number",
	"bool":              "boolean",
	"interface{}":       "any",
	"ISODateTimeString": "string",
	"time.Time":         "string",
}

type rstJsonField struct {
	codegen.TdStructField
	IsJsonPrimitive bool
	TypeStr         string
}

type rstJsonStruct struct {
	codegen.TdStruct
	Fields []rstJsonField
}

var jsonTemplate = template.Must(template.New("rstJson").Parse(`
.. _tdproto-{{- .TdStruct.Name}}:

{{.TdStruct.Name}}
-------------------------------------------------------------

{{.TdStruct.Help}}

**Fields**:
{{range $field := .Fields}}
* ` + "``" + "{{$field.TdStructField.JsonName}}" + "``" +
	` ({{- if $field.IsJsonPrimitive}}{{$field.TypeStr}}{{else}}:ref:` + "`" + `tdproto-{{$field.TdStructField.TypeStr}}` + "`" +
	`{{end}}) - {{$field.TdStructField.Help}}
{{- if $field.IsOmitEmpty}}\
  . Maybe omitted{{else -}}{{end}}
{{- if $field.IsPointer}}\
  . Might be null{{else -}}{{end}}{{end}}
`))

var enumTemplate = template.Must(template.New("rstEnum").Parse(`
.. _tdproto-{{- .Name}}:

{{.Name}} enum
-------------------------------------------------------------
**Possible values**:
{{range $value := .Values}}
* {{$value}}{{end}}

`))

var typeAliasTemplate = template.Must(template.New("rstType").Parse(`
.. _tdproto-{{- .Name}}:

{{.Name}} type alias of ` + "``{{.BaseType}}``" + `
-------------------------------------------------------------
{{.Help}}
**Base Type**: {{.BaseType}}

{{if .IsArray}}**Is array**{{end}}

`))

func generateRstJson(tdprotoInfo *codegen.TdInfo) error {

	enumedTypeAliases := make(map[string]string)

	fmt.Fprintln(os.Stdout, "Enums index\n============================")
	enumsList := tdprotoInfo.GetEnums()
	sort.Slice(enumsList, func(i, j int) bool {
		return strings.ToLower(enumsList[i].Name) < strings.ToLower(enumsList[j].Name)
	})
	for _, enum := range enumsList {
		err := enumTemplate.Execute(os.Stdout, enum)
		if err != nil {
			return err
		}
		enumedTypeAliases[enum.Name] = ""
	}

	fmt.Fprintln(os.Stdout, "Type aliases\n============================")
	var typesList []codegen.TdType
	for _, someType := range tdprotoInfo.TdTypes {
		typesList = append(typesList, someType)
	}
	sort.Slice(typesList, func(i, j int) bool {
		return strings.ToLower(typesList[i].Name) < strings.ToLower(typesList[j].Name)
	})
	for _, typeAlias := range typesList {
		_, isEnumed := enumedTypeAliases[typeAlias.Name]
		if isEnumed {
			continue
		}

		err := typeAliasTemplate.Execute(os.Stdout, typeAlias)
		if err != nil {
			return err
		}
	}

	var jsonObjects []rstJsonStruct

	for _, tdStruct := range tdprotoInfo.TdStructs {
		if tdStruct.Help == "MISSING CLASS DOCUMENTATION" {
			// Do not print structures without help
			continue
		}

		if unicode.IsLower([]rune(tdStruct.Name)[0]) {
			// Do not print private structs
			continue
		}

		isEvent := tdStruct.Name == "BaseEvent"

		newRstJson := rstJsonStruct{
			TdStruct: tdStruct,
		}

		for _, originalField := range tdStruct.Fields {
			newRstJson.Fields = append(newRstJson.Fields, rstJsonField{
				TdStructField: originalField,
			})
		}

		for _, anonStruct := range tdStruct.GetStructAnonymousStructs(tdprotoInfo) {
			if anonStruct.Name == "BaseEvent" {
				isEvent = true
			}

			for _, anonStructField := range anonStruct.Fields {
				newRstJson.Fields = append(newRstJson.Fields, rstJsonField{
					TdStructField: anonStructField,
				})
			}
		}

		if isEvent {
			continue
		}

		for i, field := range newRstJson.Fields {
			goFieldType := field.TdStructField.TypeStr

			primitiveType, ok := jsTypesMap[goFieldType]
			if ok {
				newRstJson.Fields[i].TypeStr = primitiveType
				newRstJson.Fields[i].IsJsonPrimitive = true
			} else {
				newRstJson.Fields[i].TypeStr = goFieldType
				newRstJson.Fields[i].IsJsonPrimitive = false
			}
		}

		jsonObjects = append(jsonObjects, newRstJson)
	}

	sort.Slice(jsonObjects, func(i, j int) bool {
		return strings.ToLower(jsonObjects[i].Name) < strings.ToLower(jsonObjects[j].Name)
	})

	fmt.Fprintln(os.Stdout, "JSON objects index\n============================")

	for _, object := range jsonObjects {
		err := jsonTemplate.Execute(os.Stdout, object)
		if err != nil {
			return err
		}
	}

	return nil
}

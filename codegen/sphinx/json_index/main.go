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

func (r rstJsonField) GetModifiers() string {
	if r.TdStructField.IsOmitEmpty {
		return " :abbr:`💥 (Maybe omitted)`"
	}

	if r.TdStructField.IsPointer {
		return " :abbr:`0️⃣ (Might be null)`"
	}

	return ""
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
	` ({{$field.TypeStr}}){{.GetModifiers}} - {{$field.TdStructField.Help}}{{end}}
`))

var enumTemplate = template.Must(template.New("rstEnum").Parse(`
.. _tdproto-{{- .Name}}:

{{.Name}}
-------------------------------------------------------------
**Possible values**:
{{range $value := .Values}}
* {{$value}}{{end}}

`))

var typeAliasTemplate = template.Must(template.New("rstType").Parse(`
.. _tdproto-{{- .Name}}:

{{.Name}}
-------------------------------------------------------------
{{if .Help}}
{{.Help}}
{{end}}
**Base Type**: {{.BaseType}}{{if .IsArray}}

**Is array**{{end}}
`))

var httpQueryTemplate = template.Must(template.New("rstQuery").Parse(`
.. _tdproto-{{- .Name}}Query:

{{.Name}}
-------------------------------------------------------------
{{if .Help}}
{{.Help}}{{end}}
{{range $paramName, $help := .ParamsNamesAndHelp}}
* ` + "``" + "{{$paramName}}" + "``" + ` - {{$help}}{{end}}
`))

func isEventStruct(structName string, tdprotoInfo *codegen.TdInfo) bool {
	if structName == "BaseEvent" {
		return true
	}

	_, isEvent := tdprotoInfo.TdEvents[structName]
	return isEvent
}

func generateRstJson(tdprotoInfo *codegen.TdInfo) error {

	enumedTypeAliases := make(map[string]string)

	fmt.Fprintln(os.Stdout, "\nEnums index\n============================")
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

	fmt.Fprintln(os.Stdout, "\nType aliases\n============================")
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
		if tdStruct.Help == "" {
			// Do not print structures without help
			continue
		}

		if unicode.IsLower([]rune(tdStruct.Name)[0]) {
			// Do not print private structs
			continue
		}

		if isEventStruct(tdStruct.Name, tdprotoInfo) {
			continue
		}

		if tdStruct.IsEventParams(tdprotoInfo) {
			continue
		}

		newRstJson := rstJsonStruct{
			TdStruct: tdStruct,
		}

		fieldMissingHelp := false

		for _, field := range tdStruct.GetAllJsonFields(tdprotoInfo) {
			if field.IsNotSerialized {
				continue
			}

			if field.Help == "" {
				fieldMissingHelp = true
				break
			}

			goFieldType := field.TypeStr

			var jsTypeStr string
			primitiveType, isJsonPrimitive := jsTypesMap[goFieldType]
			if isJsonPrimitive {
				jsTypeStr = primitiveType
			} else {
				jsTypeStr = fmt.Sprintf(":ref:`tdproto-%s`", goFieldType)
			}

			if field.IsList {
				jsTypeStr = fmt.Sprintf("array[%s]", jsTypeStr)
			}

			newRstJson.Fields = append(newRstJson.Fields, rstJsonField{
				TdStructField:   field,
				IsJsonPrimitive: isJsonPrimitive,
				TypeStr:         jsTypeStr,
			})
		}

		if fieldMissingHelp {
			continue
		}

		jsonObjects = append(jsonObjects, newRstJson)
	}

	sort.Slice(jsonObjects, func(i, j int) bool {
		return strings.ToLower(jsonObjects[i].Name) < strings.ToLower(jsonObjects[j].Name)
	})

	fmt.Fprintln(os.Stdout, "\nJSON objects index\n============================")

	for _, object := range jsonObjects {
		err := jsonTemplate.Execute(os.Stdout, object)
		if err != nil {
			return err
		}
	}

	fmt.Fprintln(os.Stdout, "\nHTTP Queries\n============================")

	for _, query := range tdprotoInfo.TdQueries {
		err := httpQueryTemplate.Execute(os.Stdout, query)
		if err != nil {
			return err
		}
	}

	return nil
}

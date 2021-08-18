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

	err = generateRstJson(tdprotoInfo.TdModels, "tdmodels")
	if err != nil {
		panic(err)
	}

	err = generateRstJson(tdprotoInfo.TdForms, "tdforms")
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
		return " omitempty"
	}

	if r.TdStructField.IsPointer {
		return " nullable"
	}

	return ""
}

type rstJsonStruct struct {
	codegen.TdStruct
	Fields    []rstJsonField
	TdPackage string
}

type rstEnum struct {
	codegen.TdEnum
	TdPackage string
}

type rstType struct {
	codegen.TdType
	TdPackage string
}

var jsonTemplate = template.Must(template.New("rstJson").Parse(`
.. tdproto:struct:: {{.TdStruct.Name}}
  :tdpackage: {{.TdPackage}}

  {{.TdStruct.Help}}
{{range $field := .Fields}}
  :field {{$field.TdStructField.JsonName}} {{$field.TypeStr}}{{.GetModifiers}}: {{$field.TdStructField.Help}}{{end}}
`))

var enumTemplate = template.Must(template.New("rstEnum").Parse(`
.. tdproto:enum:: {{.Name}}
  :tdpackage: {{.TdPackage}}

  **Possible values**:
{{range $value := .Values}}
  * {{$value}}{{end}}

`))

var typeAliasTemplate = template.Must(template.New("rstType").Parse(`
.. tdproto:type:: {{.Name}}
  :tdpackage: {{.TdPackage}}

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

func isEventStruct(structName string, tdprotoInfo *codegen.TdPackage) bool {
	if structName == "BaseEvent" {
		return true
	}

	_, isEvent := tdprotoInfo.TdEvents[structName]
	return isEvent
}

func generateEnumsRst(enumsList []codegen.TdEnum, tdPackageName string, enumedTypeAliases *map[string]string) error {

	fmt.Fprintf(os.Stdout, "\n%s enums index\n============================\n", tdPackageName)

	sort.Slice(enumsList, func(i, j int) bool {
		return strings.ToLower(enumsList[i].Name) < strings.ToLower(enumsList[j].Name)
	})
	for _, enum := range enumsList {
		err := enumTemplate.Execute(os.Stdout, rstEnum{
			TdEnum:    enum,
			TdPackage: tdPackageName,
		})
		if err != nil {
			return err
		}
		(*enumedTypeAliases)[enum.Name] = ""
	}

	return nil
}

func generateTypeAliasesRst(tdPackage *codegen.TdPackage, tdPackageName string, enumedTypeAliases *map[string]string) error {

	fmt.Fprintf(os.Stdout, "\n%s type aliases\n============================\n", tdPackageName)
	var typesList []codegen.TdType
	for _, someType := range tdPackage.TdTypes {
		typesList = append(typesList, someType)
	}
	sort.Slice(typesList, func(i, j int) bool {
		return strings.ToLower(typesList[i].Name) < strings.ToLower(typesList[j].Name)
	})
	for _, typeAlias := range typesList {
		_, isEnumed := (*enumedTypeAliases)[typeAlias.Name]
		if isEnumed {
			continue
		}

		err := typeAliasTemplate.Execute(os.Stdout, rstType{
			TdType:    typeAlias,
			TdPackage: tdPackageName,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func generateStructsRst(tdPackage *codegen.TdPackage, tdPackageName string, enumedTypeAliases *map[string]string) error {

	var jsonObjects []rstJsonStruct

	for _, tdStruct := range tdPackage.TdStructs {
		if tdStruct.Help == "" {
			// Do not print structures without help
			continue
		}

		if unicode.IsLower([]rune(tdStruct.Name)[0]) {
			// Do not print private structs
			continue
		}

		if isEventStruct(tdStruct.Name, tdPackage) {
			continue
		}

		if tdStruct.IsEventParams(tdPackage) {
			continue
		}

		newRstJson := rstJsonStruct{
			TdStruct:  tdStruct,
			TdPackage: tdPackageName,
		}

		fieldMissingHelp := false

		for _, field := range tdStruct.GetAllJsonFields(tdPackage) {
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
				jsTypeStr = fmt.Sprintf("`tdproto-%s`", goFieldType)
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

	fmt.Fprintf(os.Stdout, "\n%s JSON objects index\n============================\n", tdPackageName)

	for _, object := range jsonObjects {
		err := jsonTemplate.Execute(os.Stdout, object)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateQueryRst(tdPackage *codegen.TdPackage, tdPackageName string) error {
	fmt.Fprintf(os.Stdout, "\n%s HTTP Queries\n============================\n", tdPackageName)

	for _, query := range tdPackage.TdQueries {
		err := httpQueryTemplate.Execute(os.Stdout, query)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateRstJson(tdPackage *codegen.TdPackage, tdPackageName string) error {

	enumedTypeAliases := make(map[string]string)
	enumsList := tdPackage.GetEnums()

	if len(enumsList) > 0 {
		err := generateEnumsRst(enumsList, tdPackageName, &enumedTypeAliases)
		if err != nil {
			return err
		}
	}

	if len(tdPackage.TdTypes) > 0 {
		err := generateTypeAliasesRst(tdPackage, tdPackageName, &enumedTypeAliases)
		if err != nil {
			return err
		}
	}
	if len(tdPackage.TdStructs) > 0 {
		err := generateStructsRst(tdPackage, tdPackageName, &enumedTypeAliases)
		if err != nil {
			return err
		}
	}

	if len(tdPackage.TdQueries) > 0 {
		err := generateQueryRst(tdPackage, tdPackageName)
		if err != nil {
			return err
		}
	}

	return nil
}

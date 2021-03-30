package main

import (
	"fmt"
	"github.com/tada-team/tdproto/codegen"
	"os"
	"strings"
	"text/template"
)

const TypeScriptIndent = "  "
const TypeScriptClassHeaderTemplateStr = (`export class {{.Name}} implements TDProtoClass<{{.Name}}> {`)
const TypeScriptClassFieldTemplateStr = (`{{.Name}}{{if .IsOmitEmpty}} ? {{endif}}: {{ $data.TypeStr }}`)
const TypeScriptClassConstructorField = "constructor ("
const TypeScriptConstructorClose = ") {}"

const TypeScriptClosingBracket = "{"
const TypeScriptInterfaceHeaderTemplateStr = `export interface {{.Name}}JSON {`

var tsTypesMap = map[string]string{
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

var tsFieldNameSubstitutions = map[string]string{
	"Default": "isDefault",
	"New":     "isNew",
	"Public":  "isPublic",
	"Static":  "isStatic",
}

const TypeScriptHeaderStr = `interface TDProtoClass<T> {
  readonly mappableFields: ReadonlyArray<keyof T>;
}

`

const TypeScriptSumTypeTemplate = `type {{.Name}} =
  {{ range $value := .Values}} | '{{- $value -}}'
  {{end}}
`

const TypeScriptTypeTemplate = `
type {{.Name -}}JSON = {{.BaseType -}}JSON{{- if .IsArray }}[]{{end}}
type {{.Name}} = {{.BaseType}}{{- if .IsArray }}[]{{end}}

`

const TypeScriptInterfaceTemplate = `export interface {{.Name -}}JSON {
  {{- range $field :=  .Fields}}
  {{ $field.JsonName }}{{if $field.IsOmitEmpty}}?{{end}}: {{$field.TypeName -}}
    {{- if $field.IsNotPrimitive -}}JSON{{end}}
      {{- if $field.IsList -}}[]{{end}};{{end}}
}

export class {{.Name}} implements TDProtoClass<{{- .Name -}}> {
  constructor (
	{{- range $field :=  .Fields}}
    public {{if $field.IsReadOnly}}readonly {{end}}{{ $field.Name }}{{if $field.IsOmitEmpty}}?{{end}}: {{ $field.TypeName -}}
      {{- if $field.IsList -}}[]{{end}},{{end}}
  ) {}

  public static fromJSON (raw: {{.Name -}}JSON): {{.Name}} {
    return new {{.Name -}}(
      {{- range $field :=  .Fields}}
      {{- if $field.IsNotPrimitive}}
      {{if $field.IsOmitEmpty -}} raw.{{- $field.JsonName }} && {{end}}
        {{- if $field.IsList}}raw.{{- $field.JsonName }}.map({{ $field.TypeName }}.fromJSON)
        {{- else -}} {{- $field.TypeName -}}.fromJSON(raw.{{- $field.JsonName -}}){{end}}
      {{- else}}
      raw.{{- $field.JsonName -}}
      {{- end -}}
	  ,{{end}}
    )
  }

  public mappableFields = [
    {{- range $field :=  .Fields}}
    '{{- $field.Name -}}',{{end}}
  ] as const

  readonly #mapper = {
    /* eslint-disable @typescript-eslint/camelcase */
    {{- range $field :=  .Fields}}
    {{$field.Name -}}: () => ({ {{$field.JsonName -}}: this.{{$field.Name}}
      {{- if $field.IsNotPrimitive -}}{{- if $field.IsOmitEmpty}}?{{end}}
        {{- if $field.IsList}}.map(u => u.toJSON()){{else}}.toJSON(){{end}}{{end}} }),{{end}}
    /* eslint-enable @typescript-eslint/camelcase */
  }

  public toJSON (): {{.Name -}}JSON
  public toJSON (fields: Array<this['mappableFields'][number]>): Partial<{{- .Name -}}JSON>
  public toJSON (fields?: Array<this['mappableFields'][number]>) {
    if (fields && fields.length > 0) {
      return Object.assign({}, ...fields.map(f => this.#mapper[f]()))
    } else {
      return Object.assign({}, ...Object.values(this.#mapper).map(v => v()))
    }
  }
}

`

type TypeScriptTemplate struct {
	Class     *template.Template
	Interface *template.Template
}

type TypeScriptSumType struct {
	Name   string
	Values []string
}

type TypeScriptTypeInfo struct {
	Name     string
	IsArray  bool
	BaseType string
}

type TypeScriptFieldInfo struct {
	Name           string
	JsonName       string
	IsReadOnly     bool
	IsOmitEmpty    bool
	TypeName       string
	IsNotPrimitive bool
	IsList         bool
}

type TypeScriptClassInfo struct {
	Name   string
	Fields []TypeScriptFieldInfo
}

type TypeScriptInfo struct {
	Classes  []TypeScriptClassInfo
	Types    []TypeScriptTypeInfo
	SumTypes []TypeScriptSumType
}

func convertTadaInfoToTypeScript(tdprotoInfo *codegen.TdInfo) TypeScriptInfo {

	var tsInfo TypeScriptInfo
	var unwrapStructArrays = make(map[string]string)

	for _, tadaEnumInfo := range tdprotoInfo.GetEnums() {
		var tsEnumValues []string
		for _, enumValue := range tadaEnumInfo.Values {
			tsEnumValues = append(tsEnumValues, strings.Trim(enumValue, "\""))
		}

		tsInfo.SumTypes = append(tsInfo.SumTypes, TypeScriptSumType{
			Name:   tadaEnumInfo.Name,
			Values: tsEnumValues,
		})

		tsTypesMap[tadaEnumInfo.Name] = tadaEnumInfo.Name
	}

	for _, tadaTypeInfo := range tdprotoInfo.TdTypes {
		_, isPrimitive := tsTypesMap[tadaTypeInfo.BaseType]

		if isPrimitive {
			// Enums are handled elsewhere
			continue
		}

		// Unwrap arrays of structs
		if tadaTypeInfo.IsArray {
			unwrapStructArrays[tadaTypeInfo.Name] = tadaTypeInfo.BaseType
			continue
		}

		tsNewType := TypeScriptTypeInfo{
			Name:     tadaTypeInfo.Name,
			IsArray:  tadaTypeInfo.IsArray,
			BaseType: tadaTypeInfo.BaseType,
		}

		tsInfo.Types = append(tsInfo.Types, tsNewType)
	}

	for _, tadaStructInfo := range tdprotoInfo.TdStructs {

		tsNewClass := TypeScriptClassInfo{
			Name: tadaStructInfo.Name,
		}

		for _, tadaStructField := range tadaStructInfo.Fields {
			tsFieldName, isSubstituted := tsFieldNameSubstitutions[tadaStructField.Name]
			if !isSubstituted {
				tsFieldName = codegen.ToCamelCase(tadaStructField.Name)
			}

			tsTypeName, ok := tsTypesMap[tadaStructField.TypeStr]

			isNotPrimitive := false

			if !ok {
				tsTypeName = tadaStructField.TypeStr
				isNotPrimitive = true
			}

			isList := tadaStructField.IsList
			unwrappedTypeName, doUnwrap := unwrapStructArrays[tsTypeName]

			if doUnwrap {
				tsTypeName = unwrappedTypeName
				isList = true
			}

			tsNewClass.Fields = append(tsNewClass.Fields, TypeScriptFieldInfo{
				Name:           tsFieldName,
				JsonName:       tadaStructField.JsonName,
				IsReadOnly:     tadaStructField.IsReadOnly,
				IsOmitEmpty:    tadaStructField.IsOmitEmpty,
				TypeName:       tsTypeName,
				IsNotPrimitive: isNotPrimitive,
				IsList:         isList,
			})
		}

		tsInfo.Classes = append(tsInfo.Classes, tsNewClass)
	}

	return tsInfo
}

func generateTypeScript(tdprotoInfo *codegen.TdInfo) {

	tsInfo := convertTadaInfoToTypeScript(tdprotoInfo)

	classTemplate, err := template.New("tsInterface").Parse(TypeScriptInterfaceTemplate)

	if err != nil {
		panic(err)
	}

	typeTemplate, err := template.New("tsType").Parse(TypeScriptTypeTemplate)

	if err != nil {
		panic(err)
	}

	sumTemplate, err := template.New("tsSumTypes").Parse(TypeScriptSumTypeTemplate)

	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stdout, TypeScriptHeaderStr)

	for _, tsSumTypeInfo := range tsInfo.SumTypes {
		err := sumTemplate.Execute(os.Stdout, tsSumTypeInfo)
		if err != nil {
			panic(err)
		}
	}

	for _, tsTypeInfo := range tsInfo.Types {
		err := typeTemplate.Execute(os.Stdout, tsTypeInfo)
		if err != nil {
			panic(err)
		}
	}

	for _, tsClassInfo := range tsInfo.Classes {
		err := classTemplate.Execute(os.Stdout, tsClassInfo)
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

	generateTypeScript(tdprotoInfo)

}

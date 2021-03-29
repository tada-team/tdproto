package main

import (
	//"fmt"
	"github.com/tada-team/tdproto/codegen"
	"os"
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
}

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
        {{- if $field.IsList}}raw.{{- $field.JsonName }}.map({{ $field.Name }}.fromJSON)
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

func convertTadaInfoToTypeScript(tdprotoInfo *codegen.TadaInfo) []TypeScriptClassInfo {

	var tsClassesInfo []TypeScriptClassInfo

	for _, tadaStructInfo := range tdprotoInfo.TadaStructs {
		tsNewClass := TypeScriptClassInfo{
			Name: tadaStructInfo.Name,
		}

		for _, tadaStructField := range tadaStructInfo.Fields {
			tsTypeName, ok := tsTypesMap[tadaStructField.TypeStr]

			isNotPrimitive := false

			if !ok {
				tsTypeName = tadaStructField.TypeStr
				isNotPrimitive = true
			}

			tsNewClass.Fields = append(tsNewClass.Fields, TypeScriptFieldInfo{
				Name:           codegen.ToCamelCase(tadaStructField.Name),
				JsonName:       tadaStructField.JsonName,
				IsReadOnly:     tadaStructField.IsReadOnly,
				IsOmitEmpty:    tadaStructField.IsOmitEmpty,
				TypeName:       tsTypeName,
				IsNotPrimitive: isNotPrimitive,
				IsList:         tadaStructField.IsList,
			})
		}

		tsClassesInfo = append(tsClassesInfo, tsNewClass)
	}

	return tsClassesInfo
}

func generateTypeScript(tdprotoInfo *codegen.TadaInfo) {

	tsClassesInfo := convertTadaInfoToTypeScript(tdprotoInfo)

	tmpl, err := template.New("tsInterface").Parse(TypeScriptInterfaceTemplate)

	if err != nil {
		panic(err)
	}

	for _, tsClassInfo := range tsClassesInfo {
		err := tmpl.Execute(os.Stdout, tsClassInfo)
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

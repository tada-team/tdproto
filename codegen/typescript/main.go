package main

import (
	"fmt"
	"github.com/tada-team/tdproto/codegen"
	"os"
	"sort"
	"strings"
	"text/template"
)

var tsTypesMap = map[string]string{
	"JID":               "JID",
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

/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable @typescript-eslint/interface-name-prefix */

`

const TypeScriptSumTypeTemplate = `type {{.Name}} =
  {{range $value := .Values}} | '{{- $value -}}'
  {{end}}
`

const TypeScriptTypeTemplate = `
type {{.Name}} = {{.BaseType}}

`

const TypeScriptInterfaceTemplate = `export interface {{.Name -}}JSON {
  {{- range $field :=  .Fields}}
  {{$field.JsonName}}{{if $field.IsOmitEmpty}}?{{end}}: {{$field.TypeName -}}
    {{- if $field.IsNotPrimitive -}}JSON{{end}}
      {{- if $field.IsList -}}[]{{end}};{{end}}
}

export class {{.Name}} implements TDProtoClass<{{- .Name -}}> {
  /**
   * {{.Help}}
   {{range $field :=  .Fields}}* @param {{$field.Name}} {{$field.Help}}
   {{end}}*/
  constructor (
	{{- range $field :=  .Fields}}
    public {{if $field.IsReadOnly}}readonly {{end}}{{$field.Name}}{{if $field.IsOmitEmpty}}?{{end}}: {{$field.TypeName -}}
      {{- if $field.IsList -}}[]{{end}},{{end}}
  ) {}

  public static fromJSON (raw: {{.Name -}}JSON): {{.Name}} {
    return new {{.Name -}}(
      {{- range $field :=  .Fields}}
      {{- if $field.IsNotPrimitive}}
      {{if $field.IsOmitEmpty -}} raw.{{- $field.JsonName}} && {{end}}
        {{- if $field.IsList}}raw.{{- $field.JsonName}}.map({{$field.TypeName}}.fromJSON)
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

type TypeScriptSumType struct {
	Name   string
	Values []string
}

type TypeScriptTypeAliasInfo struct {
	Name     string
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
	Help           string
}

type TypeScriptClassInfo struct {
	Name   string
	Fields []TypeScriptFieldInfo
	Help   string
}

type TypeScriptInfo struct {
	Classes      []TypeScriptClassInfo
	TypesAliases []TypeScriptTypeAliasInfo
	SumTypes     []TypeScriptSumType
}

func convertTdprotoInfoToTypeScript(tdprotoInfo *codegen.TdInfo) TypeScriptInfo {

	var tsInfo TypeScriptInfo
	var unwrapStructArrays = make(map[string]string)
	var enumTypes = make(map[string]string)

	for _, tdprotoEnumInfo := range tdprotoInfo.GetEnums() {
		var tsEnumValues []string
		for _, enumValue := range tdprotoEnumInfo.Values {
			tsEnumValues = append(tsEnumValues, strings.Trim(enumValue, "\""))
		}

		tsInfo.SumTypes = append(tsInfo.SumTypes, TypeScriptSumType{
			Name:   tdprotoEnumInfo.Name,
			Values: tsEnumValues,
		})

		tsTypesMap[tdprotoEnumInfo.Name] = tdprotoEnumInfo.Name
		enumTypes[tdprotoEnumInfo.Name] = ""
	}

	for _, tdprotoTypeInfo := range tdprotoInfo.TdTypes {
		_, isEnum := enumTypes[tdprotoTypeInfo.Name]

		if isEnum {
			// Enums are handled elsewhere
			continue
		}

		// Unwrap arrays of structs
		if tdprotoTypeInfo.IsArray {
			unwrapStructArrays[tdprotoTypeInfo.Name] = tdprotoTypeInfo.BaseType
			continue
		}

		typeAliasName := tdprotoTypeInfo.Name
		typeAliasBaseType := tsTypesMap[tdprotoTypeInfo.BaseType]

		tsNewTypeAlias := TypeScriptTypeAliasInfo{
			Name:     typeAliasName,
			BaseType: typeAliasBaseType,
		}

		tsInfo.TypesAliases = append(tsInfo.TypesAliases, tsNewTypeAlias)
		tsTypesMap[typeAliasName] = typeAliasName
	}

	// Manually add JID as string
	tsInfo.TypesAliases = append(tsInfo.TypesAliases, TypeScriptTypeAliasInfo{
		Name:     "JID",
		BaseType: "string",
	})

	for _, tdprotoStructInfo := range tdprotoInfo.TdStructs {

		tsNewClass := TypeScriptClassInfo{
			Name: codegen.ToCamelCase(tdprotoStructInfo.Name),
			Help: tdprotoStructInfo.Help,
		}

		// Skip JID class
		if tsNewClass.Name == "JID" {
			continue
		}

		var tdprotoFields []codegen.TdStructField

		// Add fields defined in the struct body
		tdprotoFields = append(tdprotoFields, tdprotoStructInfo.Fields...)

		// Insert anonymous fields
		for _, anonymousFieldName := range tdprotoStructInfo.AnonnymousFields {
			anonymousStruct, ok := tdprotoInfo.TdStructs[anonymousFieldName]

			if !ok {
				panic(fmt.Errorf("anonymous struct missing %s", anonymousFieldName))
			}

			tdprotoFields = append(tdprotoFields, anonymousStruct.Fields...)
		}

		for _, tdprotoStructField := range tdprotoFields {
			tsFieldName, isSubstituted := tsFieldNameSubstitutions[tdprotoStructField.Name]
			if !isSubstituted {
				tsFieldName = codegen.ToLowerCamelCase(tdprotoStructField.Name)
			}

			tsTypeName, ok := tsTypesMap[tdprotoStructField.TypeStr]

			isNotPrimitive := false

			if !ok {
				tsTypeName = codegen.ToCamelCase(tdprotoStructField.TypeStr)
				isNotPrimitive = true
			}

			isList := tdprotoStructField.IsList
			unwrappedTypeName, doUnwrap := unwrapStructArrays[tsTypeName]

			if doUnwrap {
				tsTypeName = unwrappedTypeName
				isList = true
			}

			tsNewClass.Fields = append(tsNewClass.Fields, TypeScriptFieldInfo{
				Name:           tsFieldName,
				JsonName:       tdprotoStructField.JsonName,
				IsReadOnly:     tdprotoStructField.IsReadOnly,
				IsOmitEmpty:    tdprotoStructField.IsOmitEmpty,
				TypeName:       tsTypeName,
				IsNotPrimitive: isNotPrimitive,
				IsList:         isList,
				Help:           tdprotoStructField.Help,
			})
		}

		// Put non-optional arguments before optional
		sort.Slice(tsNewClass.Fields, func(i, j int) bool {
			if tsNewClass.Fields[i].IsOmitEmpty != tsNewClass.Fields[j].IsOmitEmpty {
				return !tsNewClass.Fields[i].IsOmitEmpty && tsNewClass.Fields[j].IsOmitEmpty
			} else {
				return (tsNewClass.Fields[i].Name < tsNewClass.Fields[j].Name)
			}

		})

		tsInfo.Classes = append(tsInfo.Classes, tsNewClass)
	}

	return tsInfo
}

func generateTypeScript(tdprotoInfo *codegen.TdInfo) {

	tsInfo := convertTdprotoInfoToTypeScript(tdprotoInfo)

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

	for _, tsTypeAliasInfo := range tsInfo.TypesAliases {
		err := typeTemplate.Execute(os.Stdout, tsTypeAliasInfo)
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

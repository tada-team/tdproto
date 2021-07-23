package main

import (
	"fmt"
	"os"
	"sort"
	"text/template"

	"github.com/tada-team/tdproto/codegen"
)

func main() {
	tdprotoInfo, err := codegen.ParseTdproto()
	if err != nil {
		panic(err)
	}

	if err := generateTypeScript(tdprotoInfo.TdModels); err != nil {
		panic(err)
	}
}

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
	"UiSettings":        "UiSettings",
}

var tsFieldNameSubstitutions = map[string]string{
	"Default": "isDefault",
	"New":     "isNew",
	"Public":  "isPublic",
	"Static":  "isStatic",
}

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

func getHelpClass(input string) string {
	if input != "" {
		return input
	} else {
		return "MISSING CLASS DOCUMENTATION"
	}
}

func getHelpField(input string) string {
	if input != "" {
		return input
	} else {
		return "DOCUMENTATION MISSING"
	}
}

func convertTdprotoInfoToTypeScript(tdprotoInfo *codegen.TdPackage) (tsInfo TypeScriptInfo, err error) {
	var unwrapStructArrays = make(map[string]string)
	var enumTypes = make(map[string]string)

	for _, tdprotoEnumInfo := range tdprotoInfo.GetEnums() {
		var tsEnumValues []string
		tsEnumValues = append(tsEnumValues, tdprotoEnumInfo.Values...)

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

	for _, tdprotoStructInfo := range tdprotoInfo.TdStructs {
		tsNewClass := TypeScriptClassInfo{
			Name: codegen.UppercaseFirstLetter(tdprotoStructInfo.Name),
			Help: getHelpClass(tdprotoStructInfo.Help),
		}

		for _, tdprotoStructField := range tdprotoStructInfo.GetAllJsonFields(tdprotoInfo) {
			if tdprotoStructField.IsNotSerialized {
				continue
			}

			tsFieldName, isSubstituted := tsFieldNameSubstitutions[tdprotoStructField.Name]
			if !isSubstituted {
				tsFieldName = codegen.SnakeCaseToLowerCamel(tdprotoStructField.JsonName)
			}

			isNotPrimitive := false

			tsTypeName, ok := tsTypesMap[tdprotoStructField.TypeStr]
			if !ok {
				tsTypeName = codegen.UppercaseFirstLetter(tdprotoStructField.TypeStr)
				isNotPrimitive = true
			}

			isList := tdprotoStructField.IsList
			unwrappedTypeName, doUnwrap := unwrapStructArrays[tsTypeName]
			if doUnwrap {
				tsTypeName = unwrappedTypeName
				isList = true
			}
			if tdprotoStructField.KeyTypeStr != "" {
				tsTypeName = fmt.Sprintf("Record<%s, %s>", tdprotoStructField.KeyTypeStr, tsTypeName)
			}

			tsNewClass.Fields = append(tsNewClass.Fields, TypeScriptFieldInfo{
				Name:           tsFieldName,
				JsonName:       tdprotoStructField.JsonName,
				IsReadOnly:     tdprotoStructField.IsReadOnly,
				IsOmitEmpty:    tdprotoStructField.IsOmitEmpty,
				TypeName:       tsTypeName,
				IsNotPrimitive: isNotPrimitive,
				IsList:         isList,
				Help:           getHelpField(tdprotoStructField.Help),
			})
		}

		// Put non-optional arguments before optional
		sort.Slice(tsNewClass.Fields, func(i, j int) bool {
			if tsNewClass.Fields[i].IsOmitEmpty != tsNewClass.Fields[j].IsOmitEmpty {
				return !tsNewClass.Fields[i].IsOmitEmpty && tsNewClass.Fields[j].IsOmitEmpty
			} else {
				return tsNewClass.Fields[i].Name < tsNewClass.Fields[j].Name
			}
		})

		tsInfo.Classes = append(tsInfo.Classes, tsNewClass)
	}

	// Sort everything by name
	sort.Slice(tsInfo.Classes, func(i, j int) bool {
		return tsInfo.Classes[i].Name < tsInfo.Classes[j].Name
	})

	sort.Slice(tsInfo.SumTypes, func(i, j int) bool {
		return tsInfo.SumTypes[i].Name < tsInfo.SumTypes[j].Name
	})

	sort.Slice(tsInfo.TypesAliases, func(i, j int) bool {
		return tsInfo.TypesAliases[i].Name < tsInfo.TypesAliases[j].Name
	})

	return tsInfo, nil
}

func generateTypeScript(tdprotoInfo *codegen.TdPackage) error {
	tsInfo, err := convertTdprotoInfoToTypeScript(tdprotoInfo)
	if err != nil {
		return err
	}

	_, _ = fmt.Fprint(os.Stdout, tsHeader)

	for _, tsSumTypeInfo := range tsInfo.SumTypes {
		if err := tsSumTypesTemplate.Execute(os.Stdout, tsSumTypeInfo); err != nil {
			return err
		}
	}

	for _, tsTypeAliasInfo := range tsInfo.TypesAliases {
		if err := tsTypeTemplate.Execute(os.Stdout, tsTypeAliasInfo); err != nil {
			return err
		}
	}

	_, err = fmt.Fprint(os.Stdout, tsManualClasses)
	if err != nil {
		return nil
	}

	for _, tsClassInfo := range tsInfo.Classes {
		if err := tsInterfaceTemplate.Execute(os.Stdout, tsClassInfo); err != nil {
			return err
		}
	}

	return nil
}

const tsHeader = `/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable @typescript-eslint/explicit-module-boundary-types */
/* eslint-disable @typescript-eslint/no-unused-vars */

interface TDProtoClass<T> {
  readonly mappableFields: ReadonlyArray<keyof T>;
}


// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type UiSettings = Record<string, any>

`

const tsManualClasses = `
export interface TeamUnreadJSON {
   /* eslint-disable camelcase */
   direct: UnreadJSON;
   group: UnreadJSON;
   task: UnreadJSON;
   /* eslint-enable camelcase */
}
 
export class TeamUnread implements TDProtoClass<TeamUnread> {
  constructor (
    public direct: Unread,
    public group: Unread,
    public task: Unread
  ) {}
 
  public static fromJSON (raw: TeamUnreadJSON): TeamUnread {
    return new TeamUnread(
      Unread.fromJSON(raw.direct),
      Unread.fromJSON(raw.group),
      Unread.fromJSON(raw.task),
    )
  }
 
  public mappableFields = [
   'direct',
   'group',
   'task',
  ] as const
 
  readonly #mapper = {
   /* eslint-disable camelcase */
   direct: () => ({ direct: this.direct.toJSON() }),
   group: () => ({ group: this.group.toJSON() }),
   task: () => ({ task: this.task.toJSON() }),
   /* eslint-enable camelcase */
  }
 
  public toJSON (): TeamUnreadJSON
  public toJSON (fields: Array<this['mappableFields'][number]>): Partial<TeamUnreadJSON>
  public toJSON (fields?: Array<this['mappableFields'][number]>) {
    if (fields && fields.length > 0) {
      return Object.assign({}, ...fields.map(f => this.#mapper[f]()))
    } else {
      return Object.assign({}, ...Object.values(this.#mapper).map(v => v()))
    }
  }
}

`

var tsInterfaceTemplate = template.Must(template.New("tsInterface").Parse(`export interface {{.Name -}}JSON {
  /* eslint-disable camelcase */
  {{- range $field :=  .Fields}}
  {{- if eq $field.TypeName "any"}}
  // eslint-disable-next-line @typescript-eslint/no-explicit-any{{- end}}
  {{$field.JsonName}}{{if $field.IsOmitEmpty}}?{{end}}: {{$field.TypeName -}}
    {{- if $field.IsNotPrimitive -}}JSON{{end}}
      {{- if $field.IsList -}}[]{{end}};{{end}}
  /* eslint-enable camelcase */
}

export class {{.Name}} implements TDProtoClass<{{- .Name -}}> {
  /**
   * {{.Help}}
   {{range $field :=  .Fields}}* @param {{$field.Name}} {{$field.Help}}
   {{end}}*/
  constructor (
	{{- range $field :=  .Fields}}
    {{- if eq $field.TypeName "any"}}
    // eslint-disable-next-line @typescript-eslint/no-explicit-any{{- end}}
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
    /* eslint-disable camelcase */
    {{- range $field :=  .Fields}}
    {{$field.Name -}}: () => ({ {{$field.JsonName -}}: this.{{$field.Name}}
      {{- if $field.IsNotPrimitive -}}{{- if $field.IsOmitEmpty}}?{{end}}
        {{- if $field.IsList}}.map(u => u.toJSON()){{else}}.toJSON(){{end}}{{end}} }),{{end}}
    /* eslint-enable camelcase */
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

`))

var tsTypeTemplate = template.Must(template.New("tsType").Parse(`
export type {{.Name}} = {{.BaseType}}

`))

var tsSumTypesTemplate = template.Must(template.New("tsSumTypes").Parse(`export type {{.Name}} =
  {{range $value := .Values}} | '{{- $value -}}'
  {{end}}
`))

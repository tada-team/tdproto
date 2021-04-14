package main

import (
	"os"
	"text/template"

	"github.com/tada-team/tdproto/codegen"
)

var dartTypeMap = map[string]string{
	"string":            "String",
	"int":               "int",
	"int64":             "int",
	"uint16":            "int",
	"uint":              "int",
	"ISODateTimeString": "DateTime",
}

var dartClassTemplate = template.Must(template.New("dartClass").Parse(`import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:tdproto_dart/tdproto_dart.dart';

part '{{.Parent.Name}}.freezed.dart';
part '{{.Parent.Name}}.g.dart';

/// {{.Parent.Help}}
@freezed
abstract class {{.Parent.Name}} with _${{.Parent.Name}} {
  const factory {{.Parent.Name}}({
    {{range $field := .Fields -}}
    /// {{$field.Parent.Help}}
    @JsonKey(name: '{{$field.Parent.JsonName}}') 
	{{- if eq $field.DartType "DateTime"}} @DateTimeConverter(){{end -}}
    {{- if $field.Parent.IsOmitEmpty}} {{else}} @required {{end -}}
    {{if $field.Parent.IsList}}List<{{$field.DartType}}> {{else}}{{$field.DartType}} {{end -}}
	{{$field.Name}},
    
    {{end}}
  }) = _{{.Parent.Name}};

  factory {{.Parent.Name}}.fromJson(Map<String, dynamic> json) => _${{.Parent.Name}}FromJson(json);
}
`))

var dartEnumTemplate = template.Must(template.New("dartEnum").Parse(`import 'package:freezed_annotation/freezed_annotation.dart';

enum {{.Name}} { {{range $value :=  .Values}}
  @JsonValue('{{$value}}')
  {{$value}},
  {{end}}
}
`))

var dartLibTemplate = template.Must(template.New("dartLib").Parse(`library tdproto_dart;

// Generated enums:
{{range $value := .GeneratedEnums}}export './src/enums/{{$value}}.dart';
{{end}}
// Generated models:
{{range $value := .GeneratedModels}}export './src/models/{{$value}}/{{$value}}.dart';
{{end}}
`))

type DartLibInfo struct {
	GeneratedEnums  []string
	GeneratedModels []string
}

type DartClassField struct {
	Name     string
	DartType string
	Parent   codegen.TdStructField
}

type DartClass struct {
	Fields []DartClassField
	Parent codegen.TdStruct
}

func generateDart(tdprotoInfo *codegen.TdPackage) error {
	for _, tdEnum := range tdprotoInfo.GetEnums() {
		err := dartEnumTemplate.Execute(os.Stdout, tdEnum)
		if err != nil {
			return err
		}
	}

	dartClasses := generateDartClasses(tdprotoInfo)

	for _, dartClass := range dartClasses {
		err := dartClassTemplate.Execute(os.Stdout, dartClass)
		if err != nil {
			return err
		}
	}

	return nil
}

func getDartTypeFromGoType(goType string) string {
	primitiveType, ok := dartTypeMap[goType]
	if !ok {
		return goType
	} else {
		return primitiveType
	}
}

func generateDartClasses(tdprotoInfo *codegen.TdPackage) (dartClasses []DartClass) {
	for structName, structInfo := range tdprotoInfo.TdStructs {

		newDartClass := DartClass{
			Parent: structInfo,
		}

		for _, tdField := range structInfo.Fields {
			newDartClass.Fields = append(newDartClass.Fields, DartClassField{
				Name:     codegen.LowercaseFirstLetter(structName),
				DartType: getDartTypeFromGoType(tdField.TypeStr),
				Parent:   tdField,
			})
		}

		dartClasses = append(dartClasses, newDartClass)
	}

	return
}

func main() {
	tdprotoInfo, err := codegen.ParseTdproto()

	if err != nil {
		panic(err)
	}

	err = generateDart(tdprotoInfo.TdModels)
	if err != nil {
		panic(err)
	}

}

package main

import (
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/tada-team/tdproto/codegen"
)

const enumsPathPrefix = "src/enums"
const modelsPathPrefix = "src/models"

var dartTypeMap = map[string]string{
	"string":            "String",
	"int":               "int",
	"int64":             "int",
	"uint16":            "int",
	"uint":              "int",
	"ISODateTimeString": "DateTime",
	"interface{}":       "dynamic",
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

func generateDartLibInfo(tdprotoInfo *codegen.TdPackage) (libInfo DartLibInfo) {
	for structName := range tdprotoInfo.TdStructs {
		libInfo.GeneratedModels = append(libInfo.GeneratedModels, structName)
	}

	for _, tdEnum := range tdprotoInfo.GetEnums() {
		libInfo.GeneratedEnums = append(libInfo.GeneratedEnums, tdEnum.Name)
	}

	return
}

func generateDart(tdprotoInfo *codegen.TdPackage, basePath string) error {
	libFile, err := os.Create(path.Join(basePath, "tdproto_dart.dart"))
	if err != nil {
		return err
	}

	err = dartLibTemplate.Execute(libFile, generateDartLibInfo(tdprotoInfo))
	if err != nil {
		return err
	}

	err = libFile.Close()
	if err != nil {
		return err
	}

	for _, tdEnum := range tdprotoInfo.GetEnums() {
		enumFile, err := os.Create(path.Join(basePath, enumsPathPrefix, fmt.Sprintf("%s.dart", tdEnum.Name)))
		if err != nil {
			return err
		}

		err = dartEnumTemplate.Execute(enumFile, tdEnum)
		if err != nil {
			return err
		}

		err = enumFile.Close()
		if err != nil {
			return err
		}
	}

	dartClasses := generateDartClasses(tdprotoInfo)

	for _, dartClass := range dartClasses {
		classFolderPath := path.Join(basePath, modelsPathPrefix, dartClass.Parent.Name)

		err := os.Mkdir(classFolderPath, 0o750)
		if err != nil {
			return nil
		}

		classFile, err := os.Create(path.Join(classFolderPath, fmt.Sprintf("%s.dart", dartClass.Parent.Name)))
		if err != nil {
			return err
		}

		err = dartClassTemplate.Execute(classFile, dartClass)
		if err != nil {
			return err
		}

		err = classFile.Close()
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

func createDirectoryStructure(basePath string) error {
	enumsPath := path.Join(basePath, enumsPathPrefix)
	modelsPath := path.Join(basePath, modelsPathPrefix)

	err := os.MkdirAll(enumsPath, 0o750)
	if err != nil {
		return err
	}

	err = os.MkdirAll(modelsPath, 0o750)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	tdprotoInfo, err := codegen.ParseTdproto()

	if err != nil {
		panic(err)
	}

	newTempDir, err := os.MkdirTemp("", "tdproto_dart")
	if err != nil {
		panic(err)
	}
	println(newTempDir)
	createDirectoryStructure(newTempDir)

	err = generateDart(tdprotoInfo.TdModels, newTempDir)
	if err != nil {
		panic(err)
	}

}

package main

import (
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/tada-team/tdproto/codegen"
)

const enumsPathPrefix = "./src/enums"
const modelsPathPrefix = "./src/models"

var dartTypeMap = map[string]string{
	"string":            "String",
	"int":               "int",
	"int64":             "int",
	"uint16":            "int",
	"uint":              "int",
	"ISODateTimeString": "DateTime",
	"interface{}":       "dynamic",
	"time.Time":         "String",
}

var dartClassTemplate = template.Must(template.New("dartClass").Parse(`import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:tdproto_dart/tdproto_dart.dart';

part '{{.SnakeCase}}.freezed.dart';
part '{{.SnakeCase}}.g.dart';

/// {{.Parent.Help}}
@freezed
abstract class {{.Name}} with _${{.Name}} {
  const factory {{.Name}}({
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
{{range $value := .GeneratedEnums}}export '{{$value}}';
{{end}}
// Generated models:
{{range $value := .GeneratedModels}}export '{{$value}}';
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
	Fields    []DartClassField
	Parent    codegen.TdStruct
	Name      string
	SnakeCase string
}

func generateDart(tdprotoInfo *codegen.TdPackage, basePath string) error {
	var libInfo DartLibInfo

	for _, tdEnum := range tdprotoInfo.GetEnums() {
		enumFileName := codegen.ToSnakeCase(tdEnum.Name)
		enumFilePath := path.Join(enumsPathPrefix, fmt.Sprintf("%s.dart", enumFileName))
		libInfo.GeneratedEnums = append(libInfo.GeneratedEnums, enumFilePath)

		enumFile, err := os.Create(path.Join(basePath, enumFilePath))
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
		dartClassFilename := dartClass.SnakeCase
		dartClassFolderPath := path.Join(modelsPathPrefix, dartClassFilename)
		dartClassFilePath := path.Join(dartClassFolderPath, fmt.Sprintf("%s.dart", dartClassFilename))
		libInfo.GeneratedModels = append(libInfo.GeneratedModels, dartClassFilePath)

		err := os.Mkdir(path.Join(basePath, dartClassFolderPath), 0o750)
		if err != nil {
			return nil
		}

		classFile, err := os.Create(path.Join(basePath, dartClassFilePath))
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

	libFile, err := os.Create(path.Join(basePath, "tdproto_dart.dart"))
	if err != nil {
		return err
	}

	err = dartLibTemplate.Execute(libFile, libInfo)
	if err != nil {
		return err
	}

	err = libFile.Close()
	if err != nil {
		return err
	}

	return nil
}

func getDartTypeFromGoType(goType string, tdprotoInfo *codegen.TdPackage) string {
	primitiveType, ok := dartTypeMap[goType]
	if ok {
		return primitiveType
	}

	for _, tdType := range tdprotoInfo.TdTypes {
		if tdType.Name == goType {
			return getDartTypeFromGoType(tdType.BaseType, tdprotoInfo)
		}
	}

	return goType
}

func generateDartClasses(tdprotoInfo *codegen.TdPackage) (dartClasses []DartClass) {
	for structName, structInfo := range tdprotoInfo.TdStructs {

		newDartClass := DartClass{
			Parent:    structInfo,
			Name:      codegen.UppercaseFirstLetter(structInfo.Name),
			SnakeCase: codegen.ToSnakeCase(structInfo.Name),
		}

		var allFields []codegen.TdStructField
		allFields = append(allFields, structInfo.Fields...)
		for _, anonStruct := range structInfo.GetStructAnonymousStructs(tdprotoInfo) {
			allFields = append(allFields, anonStruct.Fields...)
		}

		for _, tdField := range allFields {
			newDartClass.Fields = append(newDartClass.Fields, DartClassField{
				Name:     codegen.LowercaseFirstLetter(structName),
				DartType: getDartTypeFromGoType(tdField.TypeStr, tdprotoInfo),
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

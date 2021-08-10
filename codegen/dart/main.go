package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
	"unicode"

	"github.com/tada-team/tdproto/codegen"
)

const libPathPrefix = "./lib/"
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
	"bool":              "bool",
	"UiSettings":        "Map<String, dynamic>",
}

var dartFieldNameSubstitutions = map[string]string{
	"Default": "isDefault",
	"New":     "isNew",
	"Public":  "isPublic",
	"Static":  "isStatic",
}

var dartClassTemplate = template.Must(template.New("dartClass").Parse(`import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:tdproto_dart/tdproto_dart.dart';

part '{{.SnakeCase}}.freezed.dart';
part '{{.SnakeCase}}.g.dart';

{{if eq .Parent.Help "MISSING CLASS DOCUMENTATION"}}// {{else}}/// {{end}}{{.Parent.Help}}.
@freezed
abstract class {{.Name}} with _${{.Name}} {
  const factory {{.Name}}({
    {{range $field := .Fields -}}
    {{if eq $field.Parent.Help "DOCUMENTATION MISSING"}}// {{else}}/// {{end}}{{$field.Parent.Help}}.
    {{if $field.IsDeprecated}}@Deprecated('{{$field.Parent.Help}}.') {{end}}@JsonKey(name: '{{$field.Parent.JsonName}}') 
	{{- if eq $field.DartType "DateTime"}} @DateTimeConverter(){{end -}}
    {{- if $field.Parent.IsOmitEmpty}} {{else}} @required {{end -}}
    {{if $field.Parent.IsList}}List<{{$field.DartType}}> {{else}}{{$field.DartType}} {{end -}}
	{{$field.Name}},
    
    {{end}}
  }) = _{{.Name}};

  factory {{.Name}}.fromJson(Map<String, dynamic> json) => _${{.Name}}FromJson(json);
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

// Converters:
export 'src/converters/date_time_converter.dart';

// Interfaces:
export 'src/interfaces/i_response.dart';
export 'src/interfaces/i_websocket_event.dart';

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
	Name         string
	DartType     string
	IsList       bool
	IsDeprecated bool
	Parent       codegen.TdStructField
}

type DartClass struct {
	Fields    []DartClassField
	Parent    codegen.TdStruct
	Name      string
	SnakeCase string
}

func snakeCaseOrLower(input string) string {
	for _, char := range input {
		if unicode.IsLower(char) {
			return codegen.ToSnakeCase(input)
		}
	}

	return strings.ToLower(input)
}

func lowercaseFirstOrAll(input string) string {
	for _, char := range input {
		if unicode.IsLower(char) {
			return codegen.LowercaseFirstLetter(input)
		}
	}

	return strings.ToLower(input)
}

func writeFileFromTemplate(fileName string, template *template.Template, data interface{}, useExclusive bool) error {
	fileFlags := os.O_WRONLY | os.O_CREATE
	if useExclusive {
		fileFlags |= os.O_EXCL
	}

	file, err := os.OpenFile(fileName, fileFlags, 0o640)
	if err != nil {
		return err
	}

	err = template.Execute(file, data)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	b, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	log.Printf("writeFileFromTemplate:\n%s\n---", string(b))

	return nil
}

func generateDart(tdprotoInfo *codegen.TdPackage, baseLibPath string) error {
	var libInfo DartLibInfo

	for _, tdEnum := range tdprotoInfo.GetEnums() {
		enumFileName := codegen.ToSnakeCase(tdEnum.Name)
		enumFilePath := path.Join(enumsPathPrefix, fmt.Sprintf("%s.dart", enumFileName))
		libInfo.GeneratedEnums = append(libInfo.GeneratedEnums, enumFilePath)

		err := writeFileFromTemplate(path.Join(baseLibPath, enumFilePath), dartEnumTemplate, tdEnum, true)
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

		err := os.Mkdir(path.Join(baseLibPath, dartClassFolderPath), 0o750)
		if err != nil {
			return nil
		}

		err = writeFileFromTemplate(path.Join(baseLibPath, dartClassFilePath), dartClassTemplate, dartClass, true)
		if err != nil {
			return err
		}
	}

	err := writeFileFromTemplate(path.Join(baseLibPath, "./tdproto_dart.dart"), dartLibTemplate, libInfo, false)
	if err != nil {
		return err
	}

	return nil
}

func getDartTypeFromGoType(goType string, tdprotoInfo *codegen.TdPackage) (string, bool) {
	primitiveType, ok := dartTypeMap[goType]
	if ok {
		return primitiveType, false
	}

	for _, tdType := range tdprotoInfo.TdTypes {
		if tdType.Name == goType {
			unwrappedTypeName, isUnwrappedArray := getDartTypeFromGoType(tdType.BaseType, tdprotoInfo)
			return unwrappedTypeName, isUnwrappedArray || tdType.IsArray
		}
	}

	return codegen.UppercaseFirstLetter(goType), false
}

func generateDartClasses(tdprotoInfo *codegen.TdPackage) (dartClasses []DartClass) {
	for _, structInfo := range tdprotoInfo.TdStructs {

		newDartClass := DartClass{
			Parent:    structInfo,
			Name:      codegen.UppercaseFirstLetter(structInfo.Name),
			SnakeCase: snakeCaseOrLower(structInfo.Name),
		}

		var allFields []codegen.TdStructField
		allFields = append(allFields, structInfo.Fields...)
		for _, anonStruct := range structInfo.GetStructAnonymousStructs(tdprotoInfo) {
			allFields = append(allFields, anonStruct.Fields...)
		}

		for _, tdField := range allFields {
			dartTypeStr, isList := getDartTypeFromGoType(tdField.TypeStr, tdprotoInfo)

			dartFieldName, isSubstituted := dartFieldNameSubstitutions[tdField.Name]
			if !isSubstituted {
				dartFieldName = lowercaseFirstOrAll(tdField.Name)
			}

			newDartClass.Fields = append(newDartClass.Fields, DartClassField{
				Name:         dartFieldName,
				DartType:     dartTypeStr,
				IsList:       isList || tdField.IsList,
				Parent:       tdField,
				IsDeprecated: strings.Contains(tdField.Help, "Deprecated"),
			})
		}

		dartClasses = append(dartClasses, newDartClass)
	}

	// Manually add TeamUnread
	dartClasses = append(dartClasses, DartClass{
		Name:      "TeamUnread",
		SnakeCase: "team_unread",
		Parent: codegen.TdStruct{
			Help: "Manually added",
		},
		Fields: []DartClassField{
			{Name: "Direct", DartType: "Unread", Parent: codegen.TdStructField{
				Help:     "Manually added",
				JsonName: "direct",
			}},
			{Name: "Group", DartType: "Unread", Parent: codegen.TdStructField{
				Help:     "Manually added",
				JsonName: "group",
			}},
			{Name: "Task", DartType: "Unread", Parent: codegen.TdStructField{
				Help:     "Manually added",
				JsonName: "task",
			}},
		},
	})

	return
}

func createDirectoryStructure(basePath string) error {
	enumsPath := path.Join(basePath, libPathPrefix, enumsPathPrefix)
	modelsPath := path.Join(basePath, libPathPrefix, modelsPathPrefix)

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

	flag.Parse()

	var dartLibPath string
	switch len_args := len(flag.Args()); len_args {
	case 0:
		newTempDir, err := os.MkdirTemp("", "tdproto_dart")
		if err != nil {
			panic(err)
		}
		dartLibPath = newTempDir
		println("Temp directory: ", newTempDir)
	case 1:
		dartLibPath = flag.Arg(0)
	default:
		panic(fmt.Errorf("expected zero or one argument got %d", len_args))
	}

	err = createDirectoryStructure(dartLibPath)
	if err != nil {
		panic(err)
	}

	err = generateDart(tdprotoInfo.TdModels, path.Join(dartLibPath, libPathPrefix))
	if err != nil {
		panic(err)
	}

}

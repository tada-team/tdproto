package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/tada-team/tdproto/codegen"
)

func main() {
	if err := do(); err != nil {
		log.Println(err)
	}
}

var dartFile = template.Must(template.New("").Parse(`import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:tdproto_dart/tdproto_dart.dart';

part '{{.Struct.SnakeName}}.freezed.dart';
part '{{.Struct.SnakeName}}.g.dart';

{{ if $.Struct.IsEnum -}}

/// {{.Struct.Help}}
enum {{.Struct.Name}} {
  {{ range $v := $.Struct.EnumValues }}
  // {{$v.Help}}
  @JsonValue({{ $v.DartValue }})
  {{ $v.DartName }},
  {{ end }}
}

{{- else -}}

/// {{.Struct.Help}}
@freezed
abstract class {{.Struct.Name}} with _${{.Struct.Name}} {
  const factory {{.Struct.Name}}({
{{ range $f := .Struct.Fields }}
    /// {{$f.Help}}.{{if $f.Readonly}} Readonly.{{end}}
    @JsonKey(name: '{{$f.Json}}')
	{{- if eq $f.DartType "DateTime" }} @DateTimeConverter(){{ end -}}
	{{- if $f.DartRequired }} @required{{ end -}}
    {{- if $f.List }} List<{{ $f.DartType }}>{{ else }} {{ $f.DartType }}{{ end }} {{ $f.JSName }},
{{ end }}
  }) = _{{.Struct.Name}};

  factory {{.Struct.Name}}.fromJson(Map<String, dynamic> json) => _${{.Struct.Name}}FromJson(json);
}

{{- end -}}
`))

type tplContext struct {
	Struct *codegen.Struct
}

func do() error {
	if len(os.Args) != 2 {
		return errors.New("path required")
	}

	path := os.Args[1]
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	structs, err := codegen.Parse()
	if err != nil {
		return err
	}

	for _, s := range structs {
		switch s.Name {
		case "UploadPreview", "PdfVersion", "Upload", "MarkupEntity", "MarkupType":
			log.Println("export:", s.Name)
			if err := save(path, s); err != nil {
				return err
			}
		}
	}

	return nil
}

func save(path string, s *codegen.Struct) error {
	dist := filepath.Join(path, s.SnakeName())
	if _, err := os.Stat(dist); os.IsNotExist(err) {
		log.Println("mkdir:", dist)
		if err := os.MkdirAll(dist, os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	fname := filepath.Join(dist, s.SnakeName()+".dart")
	log.Println("save:", fname)

	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := dartFile.Execute(f, tplContext{Struct: s}); err != nil {
		return err
	}

	//cmd := exec.Command("dart", "fmt", fname)
	//if err := cmd.Run(); err != nil {
	//	return err
	//}

	return nil
}

package main

import (
	"os"
	"reflect"
	"text/template"

	"github.com/tada-team/tdproto/codegen/api_paths"
)

type pathDoc struct {
	MethodName  string
	Path        string
	Description string
	IsArray     bool
	ResultText  string
}

var pathsTemplate = template.Must(template.New("rstPath").Parse(`
.. http:{{- .MethodName -}}:: {{.Path}}

  {{.Description}}

  :resjson boolean ok: True if no error occured.
  {{if .IsArray}}:resjson array result:{{else}}:resjson object result:{{end}}{{.ResultText}}
  :status 200: No error.
`))

func generateSpecRst(path string, spec api_paths.HttpSpec, method string) error {

	pathsTemplate.Execute(os.Stdout, pathDoc{
		Path:        path,
		MethodName:  method,
		Description: spec.Description,
		ResultText:  spec.ResponceDescription,
		IsArray:     reflect.TypeOf(spec.Responce).Kind() == reflect.Slice,
	})

	return nil
}

func generatePathsRst() error {

	for path, spec := range api_paths.TeamPaths {
		if spec.Get != nil {
			err := generateSpecRst(path, *spec.Get, "get")
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func main() {
	err := generatePathsRst()
	if err != nil {
		panic(err)
	}
}

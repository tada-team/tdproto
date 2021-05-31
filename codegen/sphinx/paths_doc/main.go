package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/tada-team/tdproto/codegen/api_paths"
)

type pathDoc struct {
	MethodName         string
	Path               string
	Description        string
	RequestDescription string
	IsArray            bool
	ResultText         string
}

func (p pathDoc) ToSwaggerUrl() string {

	suffix := fmt.Sprintf(
		"%s%s", p.MethodName,
		strings.ReplaceAll(strings.ReplaceAll(
			strings.ReplaceAll(p.Path, "/", "_"),
			"{", "_"), "}", "_"))

	return fmt.Sprintf("`🔍 Try it! <https://tada-team.github.io/td-swagger-ui/#/default/%s>`__", suffix)
}

var pathsTemplate = template.Must(template.New("rstPath").Parse(`
.. http:{{- .MethodName -}}:: {{.Path}}

  {{.Description}}

  {{.ToSwaggerUrl}}

  :resjson boolean ok: True if no error occured.{{if .RequestDescription}}
  :reqjson object: {{.RequestDescription}}{{end}}
  {{if .IsArray}}:resjson array result:{{else}}:resjson object result:{{end}} {{.ResultText}}
  :status 200: No error.
`))

func generateSpecRst(path string, spec api_paths.HttpSpec, method string) error {

	err := pathsTemplate.Execute(os.Stdout, pathDoc{
		Path:               path,
		MethodName:         method,
		Description:        spec.Description,
		ResultText:         spec.ResponceDescription,
		RequestDescription: spec.RequestDescription,
		IsArray:            reflect.TypeOf(spec.Responce).Kind() == reflect.Slice,
	})
	if err != nil {
		return err
	}

	return nil
}

func generatePathsRst() error {

	for _, spec := range api_paths.TeamPaths {
		if spec.Get != nil {
			err := generateSpecRst(spec.Path, *spec.Get, "get")
			if err != nil {
				return err
			}
		}
		if spec.Post != nil {
			err := generateSpecRst(spec.Path, *spec.Post, "post")
			if err != nil {
				return err
			}
		}
		if spec.Put != nil {
			err := generateSpecRst(spec.Path, *spec.Put, "put")
			if err != nil {
				return err
			}
		}
		if spec.Delete != nil {
			err := generateSpecRst(spec.Path, *spec.Delete, "delete")
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

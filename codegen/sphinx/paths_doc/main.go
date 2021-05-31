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
	RequestObjectName  string
	ResultIsArray      bool
	ResultDescription  string
	ResultObjectName   string
}

func (p pathDoc) ToSwaggerUrl() string {

	suffix := fmt.Sprintf(
		"%s%s", p.MethodName,
		strings.ReplaceAll(strings.ReplaceAll(
			strings.ReplaceAll(p.Path, "/", "_"),
			"{", "_"), "}", "_"))

	return fmt.Sprintf("`🔍 Try it! <https://tada-team.github.io/td-swagger-ui/#/default/%s>`__", suffix)
}

func (p pathDoc) ToParams() string {

	var builder strings.Builder

	possibleParameters := []string{
		"team_id", "contact_id",
		"chat_id", "message_id",
		"group_id"}

	for _, paramString := range possibleParameters {
		if strings.Contains(p.Path, paramString) {
			builder.WriteString(fmt.Sprintf("\n  :param %s: ID of the %s.", paramString, strings.Split(paramString, "_")[0]))
		}
	}

	return builder.String()
}

func (p pathDoc) ToRequestText() string {
	if p.RequestDescription != "" {
		return p.RequestDescription
	}

	if p.RequestObjectName == "" {
		return ""
	}

	return fmt.Sprintf("The :ref:`tdproto-%s` object.", p.RequestObjectName)

}

func (p pathDoc) ToResultText() string {
	if p.ResultDescription != "" {
		return p.ResultDescription
	}

	if p.ResultIsArray {
		return fmt.Sprintf("Array of :ref:`tdproto-%s` objects.", p.ResultObjectName)
	} else {
		return fmt.Sprintf("The :ref:`tdproto-%s` object.", p.ResultObjectName)
	}
}

var pathsTemplate = template.Must(template.New("rstPath").Parse(`
.. http:{{- .MethodName -}}:: {{.Path}}

  {{.Description}}

  {{.ToSwaggerUrl}}
  {{.ToParams}}{{if .ToRequestText}}
  :reqjson object: {{.ToRequestText}}{{end}}
  :resjson boolean ok: True if no error occured.
  {{if .ResultIsArray}}:resjson array result:{{else}}:resjson object result:{{end}} {{.ToResultText}}
  :status 200: No error.
`))

func generateSpecRst(path string, spec api_paths.HttpSpec, method string) error {

	isArray := reflect.TypeOf(spec.Responce).Kind() == reflect.Slice

	var resultObjectName string
	if isArray {
		resultObjectName = reflect.TypeOf(spec.Responce).Elem().Name()
	} else {
		resultObjectName = reflect.TypeOf(spec.Responce).Name()
	}

	var requestObjectName string
	if spec.Request != nil {
		requestObjectName = reflect.TypeOf(spec.Request).Name()
		if requestObjectName == "" && spec.Request != nil {
			return fmt.Errorf("failed to get request type name %v", spec.Request)
		}
	}

	err := pathsTemplate.Execute(os.Stdout, pathDoc{
		Path:               path,
		MethodName:         method,
		Description:        spec.Description,
		RequestDescription: spec.RequestDescription,
		RequestObjectName:  requestObjectName,
		ResultDescription:  spec.ResponceDescription,
		ResultIsArray:      isArray,
		ResultObjectName:   resultObjectName,
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

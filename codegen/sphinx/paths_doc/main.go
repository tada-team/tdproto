package main

import (
	"flag"
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
		return fmt.Sprintf("List of :ref:`tdproto-%s` objects.", p.ResultObjectName)
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
  :resjson boolean ok: True if no error occured.{{if .ResultObjectName}}
  {{if .ResultIsArray}}:resjson array result:{{else}}:resjson object result:{{end}} {{.ToResultText}}{{end}}
  :status 200: No error.
`))

func generateSpecRst(path string, spec api_paths.HttpSpec, method string) error {

	if spec.Description == nil {
		return fmt.Errorf("path %s %s missing description", method, path)
	}

	var resultObjectName string
	var isArray bool
	if spec.Responce != nil {
		isArray = reflect.TypeOf(spec.Responce).Kind() == reflect.Slice
		if isArray {
			resultObjectName = reflect.TypeOf(spec.Responce).Elem().Name()
		} else {
			resultObjectName = reflect.TypeOf(spec.Responce).Name()
		}
	}

	var description string
	switch d := reflect.TypeOf(spec.Description).Kind(); d {
	case reflect.String:
		description = spec.Description.(string)
	case reflect.Slice:
		if reflect.TypeOf(spec.Description).Elem().Kind() != reflect.String {
			return fmt.Errorf("path %s %s description is not slice of strings", method, path)
		}
		description = strings.Join(spec.Description.([]string), "\n\n  ")
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
		Description:        description,
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

func generatePathsRst(pathCollectionName string) error {

	specCollection, found := api_paths.AllPaths[pathCollectionName]
	if !found {
		return fmt.Errorf("path collection not found %v", pathCollectionName)
	}

	collectionTitle, found := api_paths.PathTitles[pathCollectionName]
	if !found {
		return fmt.Errorf("no title found for collection %s", pathCollectionName)
	}

	fmt.Fprintf(os.Stdout, `%s
----------------------------------------------
`,
		collectionTitle)

	for _, spec := range specCollection {
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
	flag.Parse()
	arg := flag.Arg(0)

	err := generatePathsRst(arg)
	if err != nil {
		panic(err)
	}
}

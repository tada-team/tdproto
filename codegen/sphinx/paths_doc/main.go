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
	RequestQueryName   string
	ResultDescription  string
	ResultObjectName   string
	ResultKind         reflect.Kind
}

func (p pathDoc) ToSwaggerUrl() string {

	suffix := fmt.Sprintf(
		"%s%s", p.MethodName,
		strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(
			strings.ReplaceAll(p.Path, "/", "_"),
			"{", "_"), "}", "_"), ".", "_"))

	return fmt.Sprintf("`🔍 Try it! <https://tada-team.github.io/td-swagger-ui/#/default/%s>`__", suffix)
}

func (p pathDoc) ToParams() string {

	var builder strings.Builder

	possibleParameters := []string{
		"team_id", "contact_id",
		"chat_id", "message_id",
		"group_id", "task_id"}

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

	return fmt.Sprintf("The :tdproto:tdmodels:`%s` object.", p.RequestObjectName)

}

func (p pathDoc) ToResultText() string {
	if p.ResultDescription != "" {
		return p.ResultDescription
	}

	if p.ResultKind == reflect.Slice {
		return fmt.Sprintf("List of :tdproto:ref:`%s` objects.", p.ResultObjectName)
	} else {
		return fmt.Sprintf("The :tdproto:ref:`%s` object.", p.ResultObjectName)
	}
}

func (p pathDoc) ToResultType() string {
	switch p.ResultKind {
	case reflect.Slice:
		return "array"
	case reflect.Struct:
		return "object"
	case reflect.String:
		return "string"
	case reflect.Int:
		return "int"
	}

	panic(fmt.Errorf("unknown result kind %v", p.ResultKind))
}

func (p pathDoc) ToQueryLink() string {
	if p.RequestQueryName == "" {
		return ""
	}

	return fmt.Sprintf("\n\n  Query parameters: :ref:`tdproto-%sQuery`", p.RequestQueryName)
}

var pathsTemplate = template.Must(template.New("rstPath").Parse(`
.. http:{{- .MethodName -}}:: {{.Path}}

  {{.Description}}{{.ToQueryLink}}

  {{.ToSwaggerUrl}}
{{.ToParams}}{{if .ToRequestText}}
  :reqjson object: {{.ToRequestText}}{{end}}
  :resjson boolean ok: True if no error occured.{{if .ResultObjectName}}
  :resjson {{.ToResultType}} result: {{.ToResultText}}{{end}}
  :status 200: No error.
`))

func generateSpecRst(path string, spec api_paths.OperationSpec, method string) error {

	if spec.Description == nil {
		return fmt.Errorf("path %s %s missing description", method, path)
	}

	var resultObjectName string
	var resultKind reflect.Kind
	if spec.Response != nil {
		resultKind = reflect.TypeOf(spec.Response).Kind()
		if resultKind == reflect.Slice {
			resultObjectName = reflect.TypeOf(spec.Response).Elem().Name()
		} else {
			resultObjectName = reflect.TypeOf(spec.Response).Name()
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
		requestTypeOf := reflect.TypeOf(spec.Request)
		requestObjectName = requestTypeOf.Name()
		if requestObjectName == "" && spec.Request != nil {
			return fmt.Errorf("failed to get request type name %v", spec.Request)
		}
	}

	var requestQueryName string
	if spec.QueryStruct != nil {
		queryTypeOf := reflect.TypeOf(spec.QueryStruct)
		requestQueryName = queryTypeOf.Name()
		if requestQueryName == "" {
			return fmt.Errorf("failed to get query type name %v", spec.QueryStruct)
		}
	}

	err := pathsTemplate.Execute(os.Stdout, pathDoc{
		Path:               path,
		MethodName:         method,
		Description:        description,
		RequestDescription: spec.RequestDescription,
		RequestObjectName:  requestObjectName,
		RequestQueryName:   requestQueryName,
		ResultDescription:  spec.ResponseDescription,
		ResultObjectName:   resultObjectName,
		ResultKind:         resultKind,
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

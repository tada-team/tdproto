package main

import (
	"encoding/json"
	"fmt"

	"github.com/tada-team/setof"
	"github.com/tada-team/tdproto/codegen"
	"github.com/tada-team/tdproto/tdapi"
	"github.com/tada-team/tdproto/tdapi/openapi"
)

func main() {
	tdInfo, err := codegen.ParseTdproto()
	if err != nil {
		panic(err)
	}

	openapiInfo, err := generateOpenApiStruct(tdInfo)
	if err != nil {
		panic(err)
	}

	jsonBytes, err := json.MarshalIndent(openapiInfo, "", "  ")
	if err != nil {
		panic(err)
	}

	if _, err = fmt.Println(string(jsonBytes)); err != nil {
		panic(err)
	}
}

func generateOpenApiStruct(tdInfo *codegen.TdInfo) (openapi.Root, error) {
	root := openapi.Root{
		Openapi: "3.0.3",
		Info: openapi.Info{
			Title:   "tdproto",
			Version: "0.0.1", // TODO: take from git tag
		},
		Servers: []openapi.Server{
			{Url: "https://web.tada.team"},
		},
		Components: openapi.Components{
			Schemas: make(map[string]openapi.Schema),
		},
		Paths: tdapi.GetPaths(),
	}

	usedRefs := make(setof.Strings)
	for _, path := range root.Paths {
		for _, operation := range path.Iter() {
			for _, response := range operation.Responses.Iter() {
				for _, content := range response.Content.Iter() {
					usedRefs.Update(content.Schema.Refs()...)
				}
			}
		}
	}

	for name := range tdInfo.TdStructs {
		if usedRefs.Contains(openapi.SchemaRef(name)) {
			if err := addSchema(root.Components.Schemas, name, tdInfo); err != nil {
				return root, err
			}
		}
	}

	return root, nil
}

var golangTypeToOpenApi = map[string]openapi.Type{
	"string": openapi.String,
	"bool":   openapi.Boolean,
}

func addTypeSchema(components map[string]openapi.Schema, name string, tdInfo *codegen.TdInfo) error {
	tdTypeInfo, found := tdInfo.TdTypes[name]
	if !found {
		return fmt.Errorf("type alias not found: %s", name)
	}

	schema := openapi.Schema{
		Type:        golangTypeToOpenApi[tdTypeInfo.BaseType],
		Description: tdTypeInfo.Help,
	}

	components[name] = schema
	return nil
}

func addStructSchema(components map[string]openapi.Schema, name string, tdInfo *codegen.TdInfo) error {
	tdStructInfo, found := tdInfo.TdStructs[name]
	if !found {
		return fmt.Errorf("struct type not found: %s", name)
	}

	schema := openapi.Schema{
		Type:        openapi.Object,
		Properties:  make(map[string]openapi.Schema),
		Description: tdStructInfo.Help,
	}

	var allStructFields []codegen.TdStructField
	allStructFields = append(allStructFields, tdStructInfo.Fields...)
	for _, anonStruct := range tdStructInfo.GetStructAnonymousStructs(tdInfo) {
		allStructFields = append(allStructFields, anonStruct.Fields...)
	}

	for _, tdField := range allStructFields {
		prop := openapi.SchemaFromTdField(tdField)
		if prop.Ref != "" {
			if err := addSchema(components, tdField.TypeStr, tdInfo); err != nil {
				return err
			}
		}
		schema.Properties[tdField.JsonName] = prop
		if !tdField.IsOmitEmpty {
			schema.Required = append(schema.Required, tdField.JsonName)
		}
	}

	components[name] = schema
	return nil
}

func addSchema(components map[string]openapi.Schema, name string, tdInfo *codegen.TdInfo) error {
	_, found := tdInfo.TdStructs[name]
	if found {
		return addStructSchema(components, name, tdInfo)
	}

	_, found = tdInfo.TdTypes[name]
	if found {
		return addTypeSchema(components, name, tdInfo)
	}

	// HACK: add maps to codegen
	if name == "TeamUnread" {
		err := addStructSchema(components, "Unread", tdInfo)
		if err != nil {
			return err
		}

		components[name] = openapi.Schema{
			Type: openapi.Object,
			AdditionalProperties: &openapi.Schema{
				Ref: openapi.SchemaRef("Unread"),
			},
		}
		return nil
	}

	return fmt.Errorf("could not resolve schema %s", name)
}

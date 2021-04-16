package main

import (
	"encoding/json"
	"fmt"

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

var golangTypeToOpenApiType = map[string]openapi.Type{
	"string":            openapi.String,
	"int":               openapi.Number,
	"int64":             openapi.Number,
	"uint16":            openapi.Number,
	"uint":              openapi.Number,
	"bool":              openapi.Boolean,
	"ISODateTimeString": openapi.String,
	"time.Time":         openapi.String,
}

//func typeStrToSchema(typeStr string) interface{} {
//	openapiType, isPrimitive := golangTypeToOpenApiType[typeStr]
//
//	if isPrimitive {
//		return openapi.Schema{
//			Type:       openapiType,
//			Properties: nil,
//		}
//	} else {
//		return openapi.Reference{
//			ReferencePath: "#/components/schemas/" + typeStr,
//		}
//	}
//}

//func createSchemaArrayFromField(tdField codegen.TdStructField) openapi.Schema {
//	return openapi.Schema{
//		Type:  "array",
//		Items: typeStrToSchema(tdField.TypeStr),
//	}
//}

//func createSchemaFromField(tdField codegen.TdStructField) (newObject interface{}, err error) {
//	if tdField.IsList {
//		return createSchemaArrayFromField(tdField), nil
//	}
//
//	typeSchema := typeStrToSchema(tdField.TypeStr)
//
//	schema, ok := typeSchema.(openapi.Schema)
//	if ok && tdField.IsPointer {
//		schema.Nullable = true
//	}
//
//	return typeSchema, nil
//}

//func createSchemaFromType(tdType codegen.TdType) (openapi.Schema, error) {
//	if tdType.IsArray {
//		return openapi.Schema{
//			Type:  "array",
//			Items: typeStrToSchema(tdType.BaseType),
//		}, nil
//	} else {
//		return openapi.Schema{
//			Type:       golangTypeToOpenApiType[tdType.BaseType],
//			Properties: nil,
//		}, nil
//	}
//}

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

	//usedRefs := make(setof.Strings)
	//for _, path := range root.Paths {
	//	for _, operation := range path.Operations() {
	//		for _, response := range operation.Responses {
	//			for _, mediaType := range response.Content {
	//				usedRefs.Update(mediaType.Schema.ReferencePath)
	//			}
	//		}
	//	}
	//}

	//for _, tdStructInfo := range tdInfo.TdStructs {
	//	structName := tdStructInfo.Name
	//	if !usedRefs.Contains(structName) {
	//		continue
	//	}
	//
	//	openapiSchema := openapi.Schema{
	//		Type:       "object",
	//		Properties: make(map[string]interface{}),
	//		Required:   make([]string, 0),
	//	}
	//
	//	allStructFields := make([]codegen.TdStructField, 0)
	//	allStructFields = append(allStructFields, tdStructInfo.Fields...)
	//	for _, anonStruct := range tdStructInfo.GetStructAnonymousStructs(tdInfo) {
	//		allStructFields = append(allStructFields, anonStruct.Fields...)
	//	}
	//
	//	for _, tdField := range allStructFields {
	//		propertyObject, err := createSchemaFromField(tdField)
	//		if err != nil {
	//			return root, err
	//		}
	//
	//		openapiSchema.Properties[tdField.JsonName] = propertyObject
	//
	//		if !tdField.IsOmitEmpty {
	//			openapiSchema.Required = append(openapiSchema.Required, tdField.JsonName)
	//		}
	//	}
	//
	//	root.Components.Schemas[structName] = openapiSchema
	//}

	//for _, tdType := range tdInfo.TdTypes {
	//	if !usedRefs.Contains(tdType.Name) {
	//		continue
	//	}
	//
	//	newTypeSchema, err := createSchemaFromType(tdType)
	//	if err != nil {
	//		return root, err
	//	}
	//	root.Components.Schemas[tdType.Name] = newTypeSchema
	//}

	//root.Components.Schemas["interface{}"] = openapi.Schema{
	//	Type: "object",
	//}
	//
	//root.Components.Schemas["TeamUnread"] = openapi.Schema{
	//	Type: "object",
	//}
	//
	//root.Components.Schemas["UiSettings"] = openapi.Schema{
	//	Type: "object",
	//}

	return root, nil
}

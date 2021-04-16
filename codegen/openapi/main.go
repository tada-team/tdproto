package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/tada-team/tdproto/codegen"
)

var golangTypeToOpenApiType = map[string]string{
	"string":            "string",
	"int":               "number",
	"int64":             "number",
	"uint16":            "number",
	"uint":              "number",
	"bool":              "boolean",
	"ISODateTimeString": "string",
	"time.Time":         "string",
}

type OpenApiMediaType struct {
	Schema OpenApiRef `json:"schema"`
}

type OpenApiRequestBody struct {
	Content map[string]OpenApiMediaType `json:"content,omitempty"`
}

type OpenApiParameter struct {
	Name     string      `json:"name"`
	In       string      `json:"in"`
	Required bool        `json:"required"`
	Schema   interface{} `json:"schema"`
}

type OpenApiResponse struct {
	Description string                      `json:"description"`
	Content     map[string]OpenApiMediaType `json:"content"`
}

type OpenApiOperation struct {
	Summary     string                     `json:"summary"`
	RequestBody *OpenApiRequestBody        `json:"requestBody,omitempty"`
	Parameters  []OpenApiParameter         `json:"parameters,omitempty"`
	Responses   map[string]OpenApiResponse `json:"responses"`
}

type OpenApiPath struct {
	Get *OpenApiOperation `json:"get,omitempty"`
}

type OpenApiRef struct {
	ReferencePath string `json:"$ref"`
}

type OpenApiSchema struct {
	Type       string                 `json:"type"`
	IsNullable bool                   `json:"nullable,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Required   []string               `json:"required,omitempty"`
	Items      interface{}            `json:"items,omitempty"`
}

type OpenApiComponents struct {
	Schemas map[string]OpenApiSchema `json:"schemas"`
}

type OpenApiInfo struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

type OpenApiServer struct {
	Url string `json:"url"`
}

type OpenAPiRoot struct {
	OpenApiVersion string                 `json:"openapi"`
	Components     OpenApiComponents      `json:"components"`
	Info           OpenApiInfo            `json:"info"`
	Paths          map[string]OpenApiPath `json:"paths"`
	Servers        []OpenApiServer        `json:"servers,omitempty"`
}

func createResponceRefJson(someType interface{}) OpenApiResponse {
	typeToRef := reflect.TypeOf(someType)

	return OpenApiResponse{
		Content: map[string]OpenApiMediaType{
			"application/json": {
				Schema: createRefFromTypeStr(typeToRef.Name()),
			},
		},
	}
}

func createRefFromTypeStr(typeStr string) (newRef OpenApiRef) {
	newRef.ReferencePath = referenceSchema(typeStr)
	return
}

func referenceSchema(typeStr string) string {
	return fmt.Sprintf("#/components/schemas/%s", typeStr)
}

func typeStrToSchema(typeStr string) interface{} {
	openapiType, isPrimitive := golangTypeToOpenApiType[typeStr]

	if isPrimitive {
		return OpenApiSchema{
			Type:       openapiType,
			Properties: nil,
		}
	} else {
		return OpenApiRef{
			ReferencePath: fmt.Sprintf("#/components/schemas/%s", typeStr),
		}
	}
}

func createSchemaArrayFromField(tdField codegen.TdStructField) (newSchema OpenApiSchema) {

	newSchema.Type = "array"
	newSchema.Items = typeStrToSchema(tdField.TypeStr)

	return
}

func createSchemaFromField(tdField codegen.TdStructField) (newObject interface{}, err error) {

	if tdField.IsList {
		return createSchemaArrayFromField(tdField), nil
	}

	typeSchema := typeStrToSchema(tdField.TypeStr)

	schema, ok := typeSchema.(OpenApiSchema)
	if ok {
		if tdField.IsPointer {
			schema.IsNullable = true
		}

	}

	return typeSchema, nil
}

func createSchemaFromType(tdType codegen.TdType) (newSchema OpenApiSchema, err error) {
	if tdType.IsArray {
		return OpenApiSchema{
			Type:  "array",
			Items: typeStrToSchema(tdType.BaseType),
		}, nil
	} else {
		return OpenApiSchema{
			Type:       golangTypeToOpenApiType[tdType.BaseType],
			Properties: nil,
		}, nil
	}
}

func generateOpenApiStruct(tdInfo *codegen.TdInfo) (openapiInfo OpenAPiRoot, err error) {
	openapiInfo.OpenApiVersion = "3.0.3"

	openapiInfo.Info.Title = "Tdproto"
	openapiInfo.Info.Version = "0.0.1"

	openapiInfo.Paths = TdPaths

	openapiInfo.Servers = []OpenApiServer{
		{Url: "https://web.tada.team/api/v4"},
	}

	openapiInfo.Components.Schemas = make(map[string]OpenApiSchema)

	for _, tdStructInfo := range tdInfo.TdStructs {
		var openapiSchema OpenApiSchema = OpenApiSchema{
			Type:       "object",
			Properties: make(map[string]interface{}),
			Required:   make([]string, 0),
		}

		structName := tdStructInfo.Name

		allStructFields := make([]codegen.TdStructField, 0)
		allStructFields = append(allStructFields, tdStructInfo.Fields...)
		for _, anonStruct := range tdStructInfo.GetStructAnonymousStructs(tdInfo) {
			allStructFields = append(allStructFields, anonStruct.Fields...)
		}

		for _, tdField := range allStructFields {
			propertyObject, err := createSchemaFromField(tdField)
			if err != nil {
				return openapiInfo, err
			}

			openapiSchema.Properties[tdField.JsonName] = propertyObject

			if !tdField.IsOmitEmpty {
				openapiSchema.Required = append(openapiSchema.Required, tdField.JsonName)
			}
		}

		openapiInfo.Components.Schemas[structName] = openapiSchema
	}

	for _, tdType := range tdInfo.TdTypes {
		newTypeSchema, err := createSchemaFromType(tdType)
		if err != nil {
			return openapiInfo, err
		}

		openapiInfo.Components.Schemas[tdType.Name] = newTypeSchema
	}

	openapiInfo.Components.Schemas["interface{}"] = OpenApiSchema{
		Type: "object",
	}
	openapiInfo.Components.Schemas["TeamUnread"] = OpenApiSchema{
		Type: "object",
	}
	openapiInfo.Components.Schemas["UiSettings"] = OpenApiSchema{
		Type: "object",
	}

	return
}

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

	_, err = fmt.Printf("%s", jsonBytes)
	if err != nil {
		panic(err)
	}
}

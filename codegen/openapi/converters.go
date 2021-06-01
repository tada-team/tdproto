package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/tada-team/tdproto/codegen"
	"github.com/tada-team/tdproto/codegen/api_paths"
)

var golangTypeToOpenApiType = map[string]openApiType{
	"string":            openApiString,
	"int":               openApiInteger,
	"int64":             openApiInteger,
	"uint16":            openApiInteger,
	"uint":              openApiInteger,
	"bool":              openApiBoolean,
	"ISODateTimeString": openApiString,
	"time.Time":         openApiString,
	"Struct":            openApiObject,
	"Slice":             openApiArray,
	"interface{}":       openApiObject,
}

var golangTypeToOpenApiFormat = map[string]openApiFormat{
	"int32":             openApiInt32,
	"int64":             openApiInt64,
	"ISODateTimeString": openApiDateTime,
}

func schemaRef(name string) openApiRef {
	return openApiRef{
		Ref: "#/components/schemas/" + name,
	}
}

func schemaCreate(typeStr string, help string, isList bool) (newSchema openApiSchema) {
	if help != "" {
		newSchema.Description = help
	}

	primtiveType, isPrimitive := golangTypeToOpenApiType[typeStr]
	if isPrimitive {
		if isList {
			newSchema.Type = openApiArray
			newSchema.Items = &openApiSchema{Type: primtiveType}
			return
		}

		newSchema.Type = primtiveType
		return
	}

	newSchema.openApiRef = schemaRef(typeStr)

	return
}

func schemaFromTdField(tdField codegen.TdStructField) (res openApiSchema) {
	res.Description = tdField.Help

	primtiveType, isPrimitive := golangTypeToOpenApiType[tdField.TypeStr]
	if isPrimitive {
		format, hasFormat := golangTypeToOpenApiFormat[tdField.TypeStr]

		if tdField.IsList {
			res.Type = openApiArray
			res.Items = &openApiSchema{Type: primtiveType}
			if hasFormat {
				res.Items.Format = format
			}

			return
		}

		res.Type = primtiveType
		if hasFormat {
			res.Format = format
		}
		return
	}

	res.openApiRef = schemaRef(tdField.TypeStr)

	return
}

func addTypeSchema(components map[string]openApiSchema, name string, tdInfo *codegen.TdInfo) error {
	tdTypeInfo, found := tdInfo.TdTypes[name]
	if !found {
		return fmt.Errorf("type alias not found: %s", name)
	}

	schema := openApiSchema{
		Type:        golangTypeToOpenApiType[tdTypeInfo.BaseType],
		Description: tdTypeInfo.Help,
	}

	format, hasFormat := golangTypeToOpenApiFormat[name]
	if hasFormat {
		schema.Format = format
	}

	components[name] = schema
	return nil
}

func addStructSchema(components map[string]openApiSchema, name string, tdInfo *codegen.TdInfo) error {
	tdStructInfo, found := tdInfo.TdStructs[name]
	if !found {
		return fmt.Errorf("struct type not found: %s", name)
	}

	schema := openApiSchema{
		Type:        openApiObject,
		Properties:  make(map[string]openApiSchema),
		Description: tdStructInfo.Help,
	}

	var allStructFields []codegen.TdStructField
	allStructFields = append(allStructFields, tdStructInfo.Fields...)
	for _, anonStruct := range tdStructInfo.GetStructAnonymousStructs(tdInfo) {
		allStructFields = append(allStructFields, anonStruct.Fields...)
	}

	for _, tdField := range allStructFields {
		prop := schemaFromTdField(tdField)

		schema.Properties[tdField.JsonName] = prop
		if !tdField.IsOmitEmpty {
			schema.Required = append(schema.Required, tdField.JsonName)
		}
	}

	components[name] = schema
	return nil
}

func addMapType(components map[string]openApiSchema, name string, tdInfo *codegen.TdInfo) error {
	mapInfo, found := tdInfo.TdMapTypes[name]
	if !found {
		return fmt.Errorf("map type not found: %s", name)
	}

	valueSchema := schemaCreate(
		mapInfo.ValueTypeStr,
		mapInfo.Help,
		false,
	)

	mapSchema := openApiSchema{
		Type:                 openApiObject,
		AdditionalProperties: &valueSchema,
	}

	components[name] = mapSchema

	return nil
}

func interfaceToOaContents(someData interface{}, newContents *openApiContents, wrapInResult bool) error {

	resultSchema := openApiSchema{}

	if someData != nil {
		dataType := reflect.TypeOf(someData)
		switch dataType.Kind() {
		case reflect.Struct:
			resultSchema.openApiRef = schemaRef(dataType.Name())
		case reflect.Slice:
			resultSchema.Type = openApiArray
			resultSchema.Items = &openApiSchema{
				openApiRef: schemaRef(dataType.Elem().Name()),
			}
		default:
			return fmt.Errorf("cannot convert data to OpenApi %#v", someData)
		}
	}

	if wrapInResult {
		newContents.ApplicationJSON = &openApiContent{
			Schema: openApiSchema{
				Type: openApiObject,
				Properties: map[string]openApiSchema{
					"ok": {
						Type: openApiBoolean,
					},
					"result": resultSchema,
				},
			},
		}
	} else {
		newContents.ApplicationJSON = &openApiContent{
			Schema: resultSchema,
		}
	}
	return nil
}

func getDescription(method api_paths.HttpSpec) string {
	if method.Description == nil {
		return ""
	}

	switch d := reflect.TypeOf(method.Description).Kind(); d {
	case reflect.String:
		return method.Description.(string)
	case reflect.Slice:
		if reflect.TypeOf(method.Description).Elem().Kind() != reflect.String {
			return ""
		}
		return strings.Join(method.Description.([]string), "\n")
	}

	return ""
}

func convertPathSpecMethod(method api_paths.HttpSpec, operation **openApiOperation) error {
	getRepsonce := openApiResponse{}
	err := interfaceToOaContents(method.Responce, &getRepsonce.Content, true)
	if err != nil {
		return err
	}

	*operation = &openApiOperation{
		Responses: map[string]openApiResponse{
			"200": getRepsonce,
		},
		Description: getDescription(method),
	}

	if method.Request != nil {
		requestContents := openApiContents{}
		interfaceToOaContents(method.Request, &requestContents, false)

		requestBody := openApiRequestBody{
			Content: requestContents,
		}

		(*operation).RequestBody = &requestBody
	}

	return nil
}

func pathSpecToOpenApiPath(path api_paths.PathSpec, newPath *openApiPath) error {

	if path.Get != nil {
		err := convertPathSpecMethod(*path.Get, &newPath.Get)
		if err != nil {
			return err
		}
	}

	if path.Put != nil {
		err := convertPathSpecMethod(*path.Put, &newPath.Put)
		if err != nil {
			return err
		}
	}

	if path.Delete != nil {
		err := convertPathSpecMethod(*path.Delete, &newPath.Delete)
		if err != nil {
			return err
		}
	}

	if path.Post != nil {
		err := convertPathSpecMethod(*path.Post, &newPath.Post)
		if err != nil {
			return err
		}
	}

	return nil
}

func addPathParameters(path string, newPath *openApiPath) error {

	newParameters := make([]openApiParameter, 0)

	possibleParameters := []string{
		"team_id", "contact_id",
		"chat_id", "message_id",
		"group_id"}

	for _, paramString := range possibleParameters {
		if strings.Contains(path, fmt.Sprintf("{%s}", paramString)) {
			newParameters = append(newParameters, openApiParameter{
				Name:     paramString,
				In:       InPath,
				Required: true,
				Schema: openApiSchema{
					Type: openApiString,
				},
			})
		}
	}

	if len(newParameters) > 0 {
		newPath.Parameters = newParameters
	}

	return nil
}

package openapi

import "github.com/tada-team/tdproto/codegen"

type Schema struct {
	Type        Type              `json:"type,omitempty"`
	Format      Format            `json:"format,omitempty"`
	Properties  map[string]Schema `json:"properties,omitempty"`
	Items       *Schema           `json:"items,omitempty"`
	Required    []string          `json:"required,omitempty"`
	Description string            `json:"description,omitempty"`
	Example     interface{}       `json:"example,omitempty"`
	Ref         string            `json:"$ref,omitempty"`
}

func (s Schema) Refs() []string {
	var res []string

	if s.Ref != "" {
		res = append(res, s.Ref)
	}

	if s.Items != nil {
		res = append(res, s.Items.Refs()...)
	}

	for _, prop := range s.Properties {
		res = append(res, prop.Refs()...)
	}

	return res
}

func SchemaRef(name string) string {
	return "#/components/schemas/" + name
}

func SchemaFromTypeName(name string) Schema {
	openapiType, isPrimitive := golangTypeToOpenApiType[name]
	if isPrimitive {
		return Schema{Type: openapiType}
	}
	return Schema{Ref: SchemaRef(name)}
}

func SchemaArrayFromTdField(tdField codegen.TdStructField) Schema {
	schema := SchemaFromTypeName(tdField.TypeStr)
	return Schema{Type: Array, Items: &schema}
}

func SchemaFromTdField(tdField codegen.TdStructField) (res Schema) {
	if tdField.IsList {
		res = SchemaArrayFromTdField(tdField)
	} else {
		res = SchemaFromTypeName(tdField.TypeStr)
	}
	res.Description = tdField.Help
	return
}

var golangTypeToOpenApiType = map[string]Type{
	"string":            String,
	"int":               Number,
	"int64":             Number,
	"uint16":            Number,
	"uint":              Number,
	"bool":              Boolean,
	"ISODateTimeString": String,
	"time.Time":         String,
}

package openapi

type Schema struct {
	Type       Type                `json:"type,omitempty"`
	Format     Format              `json:"format,omitempty"`
	Properties map[string]Property `json:"properties,omitempty"`
	Required   []string            `json:"required,omitempty"`
	Ref        string              `json:"$ref,omitempty"`
}


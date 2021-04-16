package openapi

type Property struct {
	Type        Type                `json:"type,omitempty"`
	Format      Format              `json:"format,omitempty"`
	Description string              `json:"description,omitempty"`
	Example     interface{}         `json:"example,omitempty"`
	Ref         string              `json:"$ref,omitempty"`
	Properties  map[string]Property `json:"properties,omitempty"`
}

package openapi

func StringProperty(description string) Property {
	return Property{Type: String, Description: description}
}

func ObjectProperty(description string, properties map[string]Property) Property {
	return Property{Type: Object, Description: description, Properties: properties}
}

type Property struct {
	Type        Type                `json:"type,omitempty"`
	Format      Format              `json:"format,omitempty"`
	Description string              `json:"description,omitempty"`
	Example     interface{}         `json:"example,omitempty"`
	Ref         string              `json:"$ref,omitempty"`
	Properties  map[string]Property `json:"properties,omitempty"`
}

package openapi
//
//type Property struct {
//	Type        Type                `json:"type,omitempty"`
//	Format      Format              `json:"format,omitempty"`
//	Description string              `json:"description,omitempty"`
//	Example     interface{}         `json:"example,omitempty"`
//	Ref         string              `json:"$ref,omitempty"`
//	Properties  map[string]Property `json:"properties,omitempty"`
//}
//
//func (p Property) Refs() []string {
//	var res []string
//	if p.Ref != "" {
//		res = append(res, p.Ref)
//	}
//	for _, prop := range p.Properties {
//		res = append(res, prop.Refs()...)
//	}
//	return res
//}
//
//func StringProperty(description string, example interface{}) Property {
//	return Property{
//		Type:        String,
//		Description: description,
//		Example:     example,
//	}
//}
//
//func ObjectProperty(description string, properties map[string]Property, example interface{}) Property {
//	return Property{
//		Type:        Object,
//		Description: description,
//		Properties:  properties,
//		Example:     example,
//	}
//}

package openapi

type Contents struct {
	ApplicationJSON               *Content `json:"application/json,omitempty"`
	ApplicationXWWWFormUrlencoded *Content `json:"application/x-www-form-urlencoded,omitempty"`
}

type Content struct {
	Schema Schema `json:"schema"`
}


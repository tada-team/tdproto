package openapi

type Property struct {
	Type        Type        `json:"type"`
	Format      Format      `json:"format"`
	Description string      `json:"description,omitempty"`
	Example     interface{} `json:"example"`
}


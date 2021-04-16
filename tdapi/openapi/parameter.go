package openapi

type Parameter struct {
	Name        string `json:"name"`
	In          string `json:"in,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Description string `json:"description,omitempty"`
	Schema      Schema `json:"schema"`
}

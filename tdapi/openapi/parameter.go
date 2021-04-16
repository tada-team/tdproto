package openapi

func PathParameter(name, description string) Parameter {
	return Parameter{
		Name:        name,
		In:          "path",
		Required:    true,
		Description: description,
		Schema: Schema{
			Type: String,
		},
	}
}

type Parameter struct {
	Name        string `json:"name"`
	In          string `json:"in,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Description string `json:"description,omitempty"`
	Schema      Schema `json:"schema"`
}

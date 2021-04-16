package openapi

type Root struct {
	Openapi    string          `json:"openapi"`
	Info       Info            `json:"info"`
	Servers    []Server        `json:"servers,omitempty"`
	Paths      map[string]Path `json:"paths"`
	Components Components      `json:"components"`
}


package openapi

type Operation struct {
	Summary     string       `json:"summary"`
	Description string       `json:"description,omitempty"`
	Responses   Responses    `json:"responses"`
	RequestBody *RequestBody `json:"requestBody,omitempty"`
	//Parameters  []Parameter  `json:"parameters,omitempty"`
}

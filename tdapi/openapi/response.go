package openapi

type Responses struct {
	Status200 *Response `json:"200,omitempty"`
	Status404 *Response `json:"404,omitempty"`
	Status403 *Response `json:"403,omitempty"`
	Status422 *Response `json:"422,omitempty"`
}

type Response struct {
	Description string   `json:"description"`
	Content     Contents `json:"content"`
}


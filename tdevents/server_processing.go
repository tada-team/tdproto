package tdevents

func NewServerProcessing(num, total int, action, message string, hasError bool) (r ServerProcessing) {
	r.Name = r.GetName()
	r.Params.Action = action
	r.Params.Message = message
	r.Params.HasError = hasError
	r.Params.Num = num
	r.Params.Total = total
	return r
}

// Status of background operation
type ServerProcessing struct {
	BaseEvent
	Params serverProcessingParams `json:"params"`
}

func (p ServerProcessing) GetName() string { return "server.processing" }

// Params of the server.processing event
type serverProcessingParams struct {
	// Action name
	Action string `json:"action"`

	// Message
	Message string `json:"message"`

	// Has error
	HasError bool `json:"has_error"`

	// Current processing item
	Num int `json:"num"`

	// Total processing items
	Total int `json:"total"`
}

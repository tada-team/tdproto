package tdproto

func NewServerProcessing(num, total int, action, message string, hasError bool, actionType ActionType) (r ServerProcessing) {
	r.Name = r.GetName()
	r.Params.Action = action
	r.Params.ActionType = actionType
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

	// ActionType. Ex: [contact_import || task_import || archive_unpacking || generate_chats]
	ActionType ActionType `json:"action_type,omitempty"`

	// Message
	Message string `json:"message"`

	// Body
	Body string `json:"body,omitempty"`

	// Has error
	HasError bool `json:"has_error"`

	// Current processing item
	Num int `json:"num"`

	// Total processing items
	Total int `json:"total"`
}

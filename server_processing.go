package tdproto

func NewServerProcessing(num, total int, action, message string, hasError bool) (r ServerProcessing) {
	r.Name = r.GetName()
	r.Params.Action = action
	r.Params.Message = message
	r.Params.HasError = hasError
	r.Params.Num = num
	r.Params.Total = total
	return r
}

type ServerProcessing struct {
	BaseEvent
	Params serverProcessingParams `json:"params"`
}

func (p ServerProcessing) GetName() string { return "server.processing" }

type serverProcessingParams struct {
	Action   string `json:"action"`
	Message  string `json:"message"`
	HasError bool   `json:"has_error"`
	Num      int    `json:"num"`
	Total    int    `json:"total"`
}

package tdproto

func NewServerWarning(message, orig string) (r ServerWarning) {
	r.Name = r.GetName()
	r.Params.Message = message
	r.Params.Orig = orig
	return r
}

// Something went wrong with client message
type ServerWarning struct {
	Params serverWarningParams `json:"params"`
	BaseEvent
}

func (p ServerWarning) GetName() string { return "server.warning" }

// Params of the server.warning event
type serverWarningParams struct {
	// Debug information
	Orig interface{} `json:"orig"`

	// Message
	Message string `json:"message"`
}

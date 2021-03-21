package tdproto

func NewServerWarning(message, orig string) (r ServerWarning) {
	r.Name = r.GetName()
	r.Params.Message = message
	r.Params.Orig = orig
	return r
}

// Something went wrong with client message
type ServerWarning struct {
	BaseEvent
	Params serverWarningParams `json:"params"`
}

func (p ServerWarning) GetName() string { return "server.warning" }

type serverWarningParams struct {
	// Message
	Message string `json:"message"`

	// Debug information
	Orig interface{} `json:"orig"`
}

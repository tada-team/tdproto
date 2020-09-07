package tdproto

func NewServerWarning(message, orig string) (r ServerWarning) {
	r.BaseEvent.Name = "server.warning"
	r.Params.Message = message
	r.Params.Orig = orig
	return r
}

type ServerWarning struct {
	BaseEvent
	Params ServerWarningParams `json:"params"`
}

type ServerWarningParams struct {
	Message string      `json:"message"`
	Orig    interface{} `json:"orig"`
}

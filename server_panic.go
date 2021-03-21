package tdproto

func NewServerPanic(code string) (r ServerPanic) {
	r.Name = r.GetName()
	r.Params.Code = code
	return r
}

// Critical server error
type ServerPanic struct {
	BaseEvent
	Params serverPanicParams `json:"params"`
}

func (p ServerPanic) GetName() string { return "server.panic" }

type serverPanicParams struct {
	// Error code
	Code string `json:"code"`

	// Debug message
	Debug string `json:"debug,omitempty"`
}

package tdproto

func NewServerPanic(code string) (r ServerPanic) {
	r.Name = r.GetName()
	r.Params.Code = code
	return r
}

type ServerPanic struct {
	BaseEvent
	Params serverPanicParams `json:"params"`
}

func (p ServerPanic) GetName() string { return "server.panic" }

type serverPanicParams struct {
	Code  string `json:"code"`
	Debug string `json:"debug,omitempty"`
}

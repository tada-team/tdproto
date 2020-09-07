package tdproto

func NewServerPanic(code string) (r ServerPanic) {
	r.BaseEvent.Name = "server.panic"
	r.Params.Code = code
	return r
}

type ServerPanic struct {
	BaseEvent
	Params ServerPanicParams `json:"params"`
}

type ServerPanicParams struct {
	Code  string `json:"code"`
	Debug string `json:"debug,omitempty"`
}

package tdproto

func NewServerDebug(text string) (r ServerDebug) {
	r.BaseEvent.Name = "server.debug"
	r.Params.Text = text
	return r
}

type ServerDebug struct {
	BaseEvent
	Params ServerDebugParams `json:"params"`
}

type ServerDebugParams struct {
	Text string `json:"text"`
}

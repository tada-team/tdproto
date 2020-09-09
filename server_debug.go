package tdproto

func NewServerDebug(text string) (r ServerDebug) {
	r.BaseEvent.Name = "server.debug"
	r.Params.Text = text
	return r
}

type ServerDebug struct {
	BaseEvent
	Params serverDebugParams `json:"params"`
}

type serverDebugParams struct {
	Text string `json:"text"`
}

package tdproto

func NewServerDebug(text string) (r ServerDebug) {
	r.Name = r.GetName()
	r.Params.Text = text
	return r
}

type ServerDebug struct {
	BaseEvent
	Params serverDebugParams `json:"params"`
}

func (p ServerDebug) GetName() string { return "server.debug" }

type serverDebugParams struct {
	Text string `json:"text"`
}

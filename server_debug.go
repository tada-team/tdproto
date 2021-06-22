package tdproto

func NewServerDebug(text string) (r ServerDebug) {
	r.Name = r.GetName()
	r.Params.Text = text
	return r
}

// Debug message
type ServerDebug struct {
	BaseEvent
	Params serverDebugParams `json:"params"`
}

func (p ServerDebug) GetName() string { return "server.debug" }

// Params of the server.debug event
type serverDebugParams struct {
	// Debug message
	Text string `json:"text"`
}

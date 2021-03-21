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

type serverDebugParams struct {
	// Debug message
	Text string `json:"text"`
}

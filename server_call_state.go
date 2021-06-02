package tdproto

func NewServerCallState(callEvent CallEvent) (r ServerCallState) {
	r.Name = r.GetName()
	r.Params = callEvent
	return
}

// Call information
type ServerCallState struct {
	BaseEvent
	Params CallEvent `json:"params"`
}

func (p ServerCallState) GetName() string { return "server.call.state" }

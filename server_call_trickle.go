package tdproto

// Send trickle candidate for webrtc connection from server
type ServerCallTrickle struct {
	BaseEvent
	Params callTrickleParams `json:"params"`
}

func (p ServerCallTrickle) GetName() string { return "server.call.trickle" }

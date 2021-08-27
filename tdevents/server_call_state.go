package tdevents

import "github.com/tada-team/tdproto"

func NewServerCallState(callEvent tdproto.CallEvent) (r ServerCallState) {
	r.Name = r.GetName()
	r.Params = callEvent
	return
}

// Call information
type ServerCallState struct {
	BaseEvent
	Params tdproto.CallEvent `json:"params"`
}

func (p ServerCallState) GetName() string { return "server.call.state" }

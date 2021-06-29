package tdevents

import "github.com/tada-team/tdproto"

// Leave call
type ClientCallLeave struct {
	BaseEvent
	Params clientCallLeaveParams `json:"params"`
}

func (p ClientCallLeave) GetName() string { return "client.call.leave" }

// Params of the client.call.leave event
type clientCallLeaveParams struct {
	// Chat or contact id
	Jid tdproto.JID `json:"jid"`

	// Reason, if any
	Reason string `json:"reason,omitempty"`
}

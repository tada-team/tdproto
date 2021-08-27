package tdevents

import "github.com/tada-team/tdproto"

// Call buzzing cancelled
type ClientCallBuzzCancel struct {
	BaseEvent
	Params clientCallBuzzCancelParams `json:"params"`
}

func (p ClientCallBuzzCancel) GetName() string { return "client.call.buzzcancel" }

// Params of the client.call.buzzcancel event
type clientCallBuzzCancelParams struct {
	// Chat or contact id
	Jid tdproto.JID `json:"jid"`
}

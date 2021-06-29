package tdws

import "github.com/tada-team/tdproto"

type ClientCallBuzz struct {
	BaseEvent
	Params clientCallBuzzParams `json:"params"`
}

func (p ClientCallBuzz) GetName() string { return "client.call.buzz" }

// Call buzzing
type clientCallBuzzParams struct {
	// Chat or contact id
	Jid tdproto.JID `json:"jid"`

	// List of call participants. Empty value means all participants in call
	Members []tdproto.JID `json:"members,omitempty"`
}

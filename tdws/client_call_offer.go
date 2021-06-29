package tdws

import "github.com/tada-team/tdproto"

// Start a call
type ClientCallOffer struct {
	BaseEvent
	Params clientCallOfferParams `json:"params"`
}

func (p ClientCallOffer) GetName() string { return "client.call.offer" }

// Params of the client.call.offer event
type clientCallOfferParams struct {
	// Chat or contact id
	Jid tdproto.JID `json:"jid"`

	// Mute state
	Muted bool `json:"muted"`

	// Is trickle mode enabled
	Trickle bool `json:"trickle"`

	// SDP (session description protocol) data
	Sdp string `json:"sdp"`
}

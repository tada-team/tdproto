package tdws

import "github.com/tada-team/tdproto"

// For exchange Session Description with server when client's Local Session Description is changed
type ClientCallSdp struct {
	BaseEvent
	Params clientCallSdpParams `json:"params"`
}

func (p ClientCallSdp) GetName() string { return "client.call.sdp" }

// Params of the client.call.sdp event
type clientCallSdpParams struct {
	// Chat or contact id
	Jid tdproto.JID `json:"jid"`

	// Call id
	Uid string `json:"uid"`

	// SDP data
	JSEP tdproto.JSEP `json:"jsep"`
}

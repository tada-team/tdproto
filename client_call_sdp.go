package tdproto

// For exchange Session Description with server when client's Local Session Description is changed
type ClientCallSdp struct {
	BaseEvent
	Params clientCallSdpParams `json:"params"`
}

func (p ClientCallSdp) GetName() string { return "client.call.sdp" }

type clientCallSdpParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Call id
	Uid string `json:"uid"`

	// SDP data
	JSEP JSEP `json:"jsep"`
}

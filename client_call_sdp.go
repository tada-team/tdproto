package tdproto

// Call parameters
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

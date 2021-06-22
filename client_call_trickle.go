package tdproto

// Send trickle candidate for webrtc connection from client
type ClientCallTrickle struct {
	BaseEvent
	Params callTrickleParams `json:"params"`
}

func (p ClientCallTrickle) GetName() string { return "client.call.trickle" }

type callTrickleParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Trickle candidate
	Candidate string `json:"candidate"`

	// SDP mid
	SdpMid string `json:"sdp_mid,omitempty"`

	// SDP index
	SdpMlineIndex int `json:"sdp_mline_index,omitempty"`
}

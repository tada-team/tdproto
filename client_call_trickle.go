package tdproto

// Send trickle candidate for webrtc connection
type ClientCallTrickle struct {
	BaseEvent
	Params clientCallTrickleParams `json:"params"`
}

func (p ClientCallTrickle) GetName() string { return "client.call.trickle" }

type clientCallTrickleParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Trickle candidate
	Candidate string `json:"candidate"`

	// SDP mid
	SdpMid *string `json:"sdp_mid,omitempty"`

	// SDP index
	SdpMlineIndex *int `json:"sdp_mline_index,omitempty"`
}

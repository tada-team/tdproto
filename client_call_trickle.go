package tdproto

// Send trickle candidate for webrtc connection
type ClientCallTrickle struct {
	BaseEvent
	Params clientCallTrickleParams `json:"params"`
}

func (p ClientCallTrickle) GetName() string { return "client.call.trickle" }

type clientCallTrickleParams struct {
	Jid           JID    `json:"jid"`
	Candidate     string `json:"candidate"`
	SdpMid        string `json:"sdp_mid"`
	SdpMlineIndex int    `json:"sdp_mline_index"`
}

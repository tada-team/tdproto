package tdproto

// deprecated: use http api
type ClientCallTrickle struct {
	BaseEvent
	Params clientCallTrickleParams `json:"params"`
}

type clientCallTrickleParams struct {
	Jid           JID    `json:"jid"`
	Candidate     string `json:"candidate"`
	SdpMid        string `json:"sdp_mid"`
	SdpMlineIndex int    `json:"sdp_mline_index"`
}

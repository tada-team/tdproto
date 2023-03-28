package tdproto

func NewServerlCallTrickle(jid JID, candidate, sdpMid string, sdpMlineIndex int) (r ServerCallTrickle) {
	r.Name = r.GetName()
	r.Params.Jid = jid
	r.Params.Candidate = candidate
	r.Params.SdpMid = sdpMid
	r.Params.SdpMlineIndex = sdpMlineIndex
	return
}

// Send trickle candidate for webrtc connection
type ServerCallTrickle struct {
	BaseEvent
	Params serverCallTrickleParams `json:"params"`
}

func (p ServerCallTrickle) GetName() string { return "server.call.trickle" }

// Params of server.call.trickle event
type serverCallTrickleParams struct {
	// Chat or contact id
	Jid JID `json:"jid,omitempty"`

	// Trickle candidate
	Candidate string `json:"candidate,omitempty"`

	// SDP mid
	SdpMid string `json:"sdp_mid,omitempty"`

	// SDP index
	SdpMlineIndex int `json:"sdp_mline_index,omitempty"`
}

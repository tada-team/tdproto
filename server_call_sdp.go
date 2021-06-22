package tdproto

func NewServerCallSdp(jid JID, uid, sdpType, sdp string) (r ServerCallSdp) {
	r.Name = r.GetName()
	r.Params.Jid = jid
	r.Params.JSEP.Type = sdpType
	r.Params.JSEP.SDP = sdp
	r.Params.Uid = uid
	return r
}

// For exchange Session Description with client when server's Local Session Description is changed
type ServerCallSdp struct {
	BaseEvent
	Params serverCallSdpParams `json:"params"`
}

func (p ServerCallSdp) GetName() string { return "server.call.sdp" }

// Params of the server.call.sdp event
type serverCallSdpParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Call id
	Uid string `json:"uid"`

	// SDP data
	JSEP JSEP `json:"jsep"`
}

package tdproto

func NewServerCallSdp(jid JID, uid, sdpType, sdp string) (r ServerCallSdp) {
	r.Name = r.GetName()
	r.Params.Jid = jid
	r.Params.JSEP.Type = sdpType
	r.Params.JSEP.SDP = sdp
	r.Params.Uid = uid
	return r
}

// Call parameters
type ServerCallSdp struct {
	BaseEvent
	Params serverCallSdpParams `json:"params"`
}

func (p ServerCallSdp) GetName() string { return "server.call.sdp" }

type serverCallSdpParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Call id
	Uid string `json:"uid"`

	// SDP data
	JSEP JSEP `json:"jsep"`
}

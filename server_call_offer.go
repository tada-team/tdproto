package tdproto

// Signal client about offer
type ServerCallOffer struct {
	BaseEvent
	Params serverCallOfferParams `json:"params"`
}

func (p ServerCallOffer) GetName() string { return "server.call.offer" }

type serverCallOfferParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`
	// SDP (session description protocol) data
	Sdp string `json:"sdp"`
}

func NewServerCallOffer(jid JID, sdp string) (r ServerCallOffer) {
	r.Name = r.GetName()
	r.Params.Jid = jid
	r.Params.Sdp = sdp
	return r
}

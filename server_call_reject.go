package tdproto

func NewServerCallReject(jid JID, reason string, uid string) (r ServerCallReject) {
	r.Name = r.GetName()
	r.Params.Jid = jid
	r.Params.Reason = reason
	r.Params.Uid = uid
	return r
}

// Call rejected
type ServerCallReject struct {
	BaseEvent
	Params serverCallRejectParams `json:"params"`
}

func (p ServerCallReject) GetName() string { return "server.call.reject" }

type serverCallRejectParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Call id
	Uid string `json:"uid"`

	// Reason, if any
	Reason string `json:"reason"`
}

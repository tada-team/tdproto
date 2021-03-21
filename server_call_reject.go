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
	Jid    JID    `json:"jid"`
	Reason string `json:"reason"`
	Uid    string `json:"uid"`
}

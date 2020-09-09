package tdproto

func NewServerCallReject(jid JID, reason string, uid string) (r ServerCallReject) {
	r.BaseEvent.Name = "server.call.reject"
	r.Params.Jid = jid
	r.Params.Reason = reason
	r.Params.Uid = uid
	return r
}

type ServerCallReject struct {
	BaseEvent
	Params serverCallRejectParams `json:"params"`
}

type serverCallRejectParams struct {
	Jid    JID    `json:"jid"`
	Reason string `json:"reason"`
	Uid    string `json:"uid"`
}

package tdproto

func NewServerCallLeave(jid JID, uid string) (r ServerCallLeave) {
	r.BaseEvent.Name = "server.call.leave"
	r.Params.Jid = jid
	r.Params.Uid = uid
	return r
}

type ServerCallLeave struct {
	BaseEvent
	Params ServerCallLeaveParams `json:"params"`
}

type ServerCallLeaveParams struct {
	Jid JID    `json:"jid"`
	Uid string `json:"uid"`
}

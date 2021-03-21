package tdproto

func NewServerCallLeave(jid JID, uid string) (r ServerCallLeave) {
	r.Name = r.GetName()
	r.Params.Jid = jid
	r.Params.Uid = uid
	return r
}

// Participant leave a call
type ServerCallLeave struct {
	BaseEvent
	Params serverCallLeaveParams `json:"params"`
}

func (p ServerCallLeave) GetName() string { return "server.call.leave" }

type serverCallLeaveParams struct {
	Jid JID    `json:"jid"`
	Uid string `json:"uid"`
}

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

// Params of the server.call.leave event
type serverCallLeaveParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Call uid
	Uid string `json:"uid"`
}

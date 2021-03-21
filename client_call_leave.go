package tdproto

// Leave call
type ClientCallLeave struct {
	BaseEvent
	Params clientCallLeaveParams `json:"params"`
}

func (p ClientCallLeave) GetName() string { return "client.call.leave" }

type clientCallLeaveParams struct {
	Jid    JID    `json:"jid"`
	Reason string `json:"reason"`
}

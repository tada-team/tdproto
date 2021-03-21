package tdproto

// Leave call
type ClientCallLeave struct {
	BaseEvent
	Params clientCallLeaveParams `json:"params"`
}

func (p ClientCallLeave) GetName() string { return "client.call.leave" }

type clientCallLeaveParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Reason, if any
	Reason string `json:"reason"`
}

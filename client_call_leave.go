package tdproto

// deprecated: use http api
type ClientCallLeave struct {
	BaseEvent
	Params clientCallLeaveParams `json:"params"`
}

func (p ClientCallLeave) GetName() string { return "client.call.leave" }

type clientCallLeaveParams struct {
	Jid                 JID    `json:"jid"`
	Reason              string `json:"reason"`
	LeaveWithoutClosing bool   `json:"test_not_in_room,omitempty"`
}

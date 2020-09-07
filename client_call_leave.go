package tdproto

type ClientCallLeave struct {
	BaseEvent
	Params ClientCallLeaveParams `json:"params"`
}

type ClientCallLeaveParams struct {
	Jid                 JID    `json:"jid"`
	Reason              string `json:"reason"`
	LeaveWithoutClosing bool   `json:"test_not_in_room,omitempty"`
}

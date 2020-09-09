package tdproto

type ClientCallLeave struct {
	BaseEvent
	Params clientCallLeaveParams `json:"params"`
}

type clientCallLeaveParams struct {
	Jid                 JID    `json:"jid"`
	Reason              string `json:"reason"`
	LeaveWithoutClosing bool   `json:"test_not_in_room,omitempty"`
}

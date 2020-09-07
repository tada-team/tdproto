package tdproto

type ClientCallReject struct {
	BaseEvent
	Params ClientCallRejectParams `json:"params"`
}

type ClientCallRejectParams struct {
	Jid    JID    `json:"jid"`
	Reason string `json:"reason"`
}

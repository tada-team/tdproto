package tdproto

type ClientCallReject struct {
	BaseEvent
	Params clientCallRejectParams `json:"params"`
}

type clientCallRejectParams struct {
	Jid    JID    `json:"jid"`
	Reason string `json:"reason"`
}

package tdproto

// Reject the call
type ClientCallReject struct {
	BaseEvent
	Params clientCallRejectParams `json:"params"`
}

func (p ClientCallReject) GetName() string { return "client.call.reject" }

type clientCallRejectParams struct {
	Jid    JID    `json:"jid"`
	Reason string `json:"reason"`
}

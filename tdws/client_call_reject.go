package tdws

// Reject the call
type ClientCallReject struct {
	BaseEvent
	Params clientCallRejectParams `json:"params"`
}

func (p ClientCallReject) GetName() string { return "client.call.reject" }

// Params of the client.call.reject event
type clientCallRejectParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Reason, if any
	Reason string `json:"reason,omitempty"`
}

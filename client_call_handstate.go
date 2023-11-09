package tdproto

// Change hand state in call
type ClientCallHandState struct {
	BaseEvent
	Params clientCallHandStateParams `json:"params"`
}

func (p ClientCallHandState) GetName() string { return "client.call.handstate" }

// Params of the client.call.handstate event
type clientCallHandStateParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// HandState - state of hand
	HandState bool `json:"hand_state"`
}

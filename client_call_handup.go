package tdproto

// Change handup state in call
type ClientCallHandUp struct {
	BaseEvent
	Params clientCallHandUpParams `json:"params"`
}

func (p ClientCallHandUp) GetName() string { return "client.call.handup" }

// Params of the client.call.handup event
type clientCallHandUpParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Enabled handup state
	Enabled bool `json:"enabled"`
}

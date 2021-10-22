package tdproto

// Change video state in call
type ClientCallVideo struct {
	BaseEvent
	Params clientCallVideoParams `json:"params"`
}

func (p ClientCallVideo) GetName() string { return "client.call.video" }

// Params of the client.call.video event
type clientCallVideoParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Enable video state
	Enabled bool `json:"enabled"`
}

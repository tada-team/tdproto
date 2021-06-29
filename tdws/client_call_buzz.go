package tdws

// deprecated: use http api
type ClientCallBuzz struct {
	BaseEvent
	Params clientCallBuzzParams `json:"params"`
}

func (p ClientCallBuzz) GetName() string { return "client.call.buzz" }

// Call buzzing
type clientCallBuzzParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// List of call participants. Empty value means all participants in call
	Members []JID `json:"members,omitempty"`
}

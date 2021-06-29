package tdws

// Change mute state in call
type ClientCallSound struct {
	BaseEvent
	Params clientCallSoundParams `json:"params"`
}

func (p ClientCallSound) GetName() string { return "client.call.sound" }

// Params of the client.call.sound event
type clientCallSoundParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Mute state
	Muted bool `json:"muted"`
}

package tdproto

// deprecated: use http api
type ClientCallSound struct {
	BaseEvent
	Params clientCallSoundParams `json:"params"`
}

func (p ClientCallSound) GetName() string { return "client.call.sound" }

type clientCallSoundParams struct {
	Jid   JID  `json:"jid"`
	Muted bool `json:"muted"`
}

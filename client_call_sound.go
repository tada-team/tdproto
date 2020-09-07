package tdproto

type ClientCallSound struct {
	BaseEvent
	Params ClientCallSoundParams `json:"params"`
}

type ClientCallSoundParams struct {
	Jid   JID  `json:"jid"`
	Muted bool `json:"muted"`
}

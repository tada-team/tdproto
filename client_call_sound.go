package tdproto

type ClientCallSound struct {
	BaseEvent
	Params clientCallSoundParams `json:"params"`
}

type clientCallSoundParams struct {
	Jid   JID  `json:"jid"`
	Muted bool `json:"muted"`
}

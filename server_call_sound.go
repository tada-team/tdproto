package tdproto

func NewServerCallSound(jid JID, muted bool) (r ServerCallSound) {
	r.BaseEvent.Name = "server.call.sound"
	r.Params.Jid = jid
	r.Params.Muted = muted
	return r
}

type ServerCallSound struct {
	BaseEvent
	Params ServerCallSoundParams `json:"params"`
}

type ServerCallSoundParams struct {
	Jid   JID  `json:"jid"`
	Muted bool `json:"muted"`
}

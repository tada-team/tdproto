package tdproto

func NewServerCallSound(jid JID, muted bool) (r ServerCallSound) {
	r.Name = r.GetName()
	r.Params.Jid = jid
	r.Params.Muted = muted
	return r
}

type ServerCallSound struct {
	BaseEvent
	Params serverCallSoundParams `json:"params"`
}

func (p ServerCallSound) GetName() string { return "server.call.sound" }

type serverCallSoundParams struct {
	Jid   JID  `json:"jid"`
	Muted bool `json:"muted"`
}

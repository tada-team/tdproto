package tdproto

func NewServerCallSound(jid JID, muted bool) (r ServerCallSound) {
	r.Name = r.GetName()
	r.Params.Jid = jid
	r.Params.Muted = muted
	return r
}

// Mute/unmute call participant
type ServerCallSound struct {
	BaseEvent
	Params ServerCallSoundParams `json:"params"`
}

func (p ServerCallSound) GetName() string { return "server.call.sound" }

type ServerCallSoundParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Mute state
	Muted bool `json:"muted"`
}

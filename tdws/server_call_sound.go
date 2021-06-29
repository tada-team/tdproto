package tdws

import "github.com/tada-team/tdproto"

func NewServerCallSound(jid tdproto.JID, muted bool) (r ServerCallSound) {
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

// Params of the server.call.sound event
type ServerCallSoundParams struct {
	// Chat or contact id
	Jid tdproto.JID `json:"jid"`

	// Mute state
	Muted bool `json:"muted"`
}

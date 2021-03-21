package tdproto

func NewServerCallMuteall(jid JID, muted bool) (r ServerCallMuteall) {
	r.Name = r.GetName()
	r.Params.Jid = jid
	r.Params.Muted = muted
	return r
}

// All participants in call muted
type ServerCallMuteall struct {
	BaseEvent
	Params serverCallMuteallParams `json:"params"`
}

func (p ServerCallMuteall) GetName() string { return "server.call.muteall" }

type serverCallMuteallParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Mute state
	Muted bool `json:"muted"`
}

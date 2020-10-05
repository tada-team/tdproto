package tdproto

func NewServerCallMuteall(jid JID, muted bool, error string) (r ServerCallMuteall) {
	r.Name = r.GetName()
	r.Params.Jid = jid
	r.Params.Muted = muted
	r.Params.Error = error
	return r
}

type ServerCallMuteall struct {
	BaseEvent
	Params serverCallMuteallParams `json:"params"`
}

func (p ServerCallMuteall) GetName() string { return "server.call.muteall" }

type serverCallMuteallParams struct {
	Jid   JID    `json:"jid"`
	Muted bool   `json:"muted"`
	Error string `json:"reason"`
}

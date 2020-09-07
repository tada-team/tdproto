package tdproto

func NewServerCallMuteall(jid JID, muted bool, error string) (r ServerCallMuteall) {
	r.BaseEvent.Name = "server.call.muteall"
	r.Params.Jid = jid
	r.Params.Muted = muted
	r.Params.Error = error
	return r
}

type ServerCallMuteall struct {
	BaseEvent
	Params ServerCallMuteallParams `json:"params"`
}

type ServerCallMuteallParams struct {
	Jid   JID    `json:"jid"`
	Muted bool   `json:"muted"`
	Error string `json:"reason"`
}

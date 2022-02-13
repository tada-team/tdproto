package tdproto

func NewServerCallScreenShare(screenShareEnabled bool, callJid, actorJid JID) (r ServerCallScreenShare) {
	r.Name = r.GetName()
	r.Params.ScreenShareEnabled = screenShareEnabled
	r.Params.CallJid = callJid
	r.Params.ActorJid = actorJid
	return r
}

// ServerCallScreenShare screen share event
type ServerCallScreenShare struct {
	BaseEvent
	Params serverCallScreenShareParams `json:"params"`
}

func (p ServerCallScreenShare) GetName() string { return "server.call.screenshare" }

// Params of the server.call.screenshare event
type serverCallScreenShareParams struct {
	// Chat or contact id
	CallJid JID `json:"call_jid"`

	// Contact id which enable/disable screen sharing
	ActorJid JID `json:"actor_jid"`

	// Enabled or disabled screen share
	ScreenShareEnabled bool `json:"screenshare_enabled"`
}

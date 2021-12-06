package tdproto

// ClientCallScreenShare event for enable/disable screen share
type ClientCallScreenShare struct {
	BaseEvent
	Params clientCallScreenShareParams `json:"params"`
}

func (p ClientCallScreenShare) GetName() string { return "client.call.screenshare" }

// Params of the client.call.screenshare event
type clientCallScreenShareParams struct {
	// ScreenShareEnabled enabled or disabled screen share
	ScreenShareEnabled bool `json:"screenshare_enabled"`

	// CallJid Chat or contact id
	CallJid JID `json:"call_jid"`
}

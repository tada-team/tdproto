package tdproto

// Mute all other call participants
type ClientCallMuteAll struct {
	BaseEvent
	Params clientCallMuteAllParams `json:"params"`
}

func (p ClientCallMuteAll) GetName() string { return "client.call.muteall" }

type clientCallMuteAllParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`
}

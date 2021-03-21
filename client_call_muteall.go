package tdproto

// Mute all other call participants
type ClientCallMuteAll struct {
	BaseEvent
	Params clientCallMuteAllParams `json:"params"`
}

func (p ClientCallMuteAll) GetName() string { return "client.call.muteall" }

type clientCallMuteAllParams struct {
	Jid JID `json:"jid"`
}

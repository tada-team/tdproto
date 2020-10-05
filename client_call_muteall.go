package tdproto

// deprecated: use http api
type ClientCallMuteAll struct {
	BaseEvent
	Params clientCallMuteAllParams `json:"params"`
}

func (p ClientCallMuteAll) GetName() string { return "client.call.muteall" }

type clientCallMuteAllParams struct {
	Jid JID `json:"jid"`
}

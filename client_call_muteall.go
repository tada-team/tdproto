package tdproto

type ClientCallMuteAll struct {
	BaseEvent
	Params clientCallMuteAllParams `json:"params"`
}

type clientCallMuteAllParams struct {
	Jid JID `json:"jid"`
}

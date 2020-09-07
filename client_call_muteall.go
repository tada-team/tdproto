package tdproto

type ClientCallMuteAll struct {
	BaseEvent
	Params ClientCallMuteAllParams `json:"params"`
}

type ClientCallMuteAllParams struct {
	Jid JID `json:"jid"`
}

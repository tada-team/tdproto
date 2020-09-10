package tdproto

// deprecated: use http api
type ClientCallMuteAll struct {
	BaseEvent
	Params clientCallMuteAllParams `json:"params"`
}

type clientCallMuteAllParams struct {
	Jid JID `json:"jid"`
}

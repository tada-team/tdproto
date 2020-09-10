package tdproto

// deprecated: use http api
type ClientCallBuzzCancel struct {
	BaseEvent
	Params clientCallBuzzCancelParams `json:"params"`
}

type clientCallBuzzCancelParams struct {
	Jid JID `json:"jid"`
}

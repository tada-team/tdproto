package tdproto

type ClientCallBuzzCancel struct {
	BaseEvent
	Params clientCallBuzzCancelParams `json:"params"`
}

type clientCallBuzzCancelParams struct {
	Jid JID `json:"jid"`
}

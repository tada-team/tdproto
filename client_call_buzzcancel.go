package tdproto

type ClientCallBuzzCancel struct {
	BaseEvent
	Params ClientCallBuzzCancelParams `json:"params"`
}

type ClientCallBuzzCancelParams struct {
	Jid JID `json:"jid"`
}

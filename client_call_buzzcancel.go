package tdproto

// Call buzzing cancelled
type ClientCallBuzzCancel struct {
	BaseEvent
	Params clientCallBuzzCancelParams `json:"params"`
}

func (p ClientCallBuzzCancel) GetName() string { return "client.call.buzzcancel" }

type clientCallBuzzCancelParams struct {
	Jid JID `json:"jid"`
}

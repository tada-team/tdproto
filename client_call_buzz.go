package tdproto

// deprecated: use http api
type ClientCallBuzz struct {
	BaseEvent
	Params clientCallBuzzParams `json:"params"`
}

func (p ClientCallBuzz) GetName() string { return "client.call.buzz" }

type clientCallBuzzParams struct {
	Jid     JID   `json:"jid"`
	Members []JID `json:"members"`
}

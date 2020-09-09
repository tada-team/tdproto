package tdproto

type ClientCallBuzz struct {
	BaseEvent
	Params clientCallBuzzParams `json:"params"`
}

type clientCallBuzzParams struct {
	Jid     JID   `json:"jid"`
	Members []JID `json:"members"`
}

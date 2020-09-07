package tdproto

type ClientCallBuzz struct {
	BaseEvent
	Params ClientCallBuzzParams `json:"params"`
}

type ClientCallBuzzParams struct {
	Jid     JID   `json:"jid"`
	Members []JID `json:"members"`
}

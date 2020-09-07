package tdproto

type ClientCallOffer struct {
	BaseEvent
	Params ClientCallOfferParams `json:"params"`
}

type ClientCallOfferParams struct {
	Jid     JID    `json:"jid"`
	Muted   bool   `json:"muted"`
	Trickle bool   `json:"trickle"`
	Sdp     string `json:"sdp"`
}

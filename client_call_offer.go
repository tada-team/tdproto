package tdproto

type ClientCallOffer struct {
	BaseEvent
	Params clientCallOfferParams `json:"params"`
}

type clientCallOfferParams struct {
	Jid     JID    `json:"jid"`
	Muted   bool   `json:"muted"`
	Trickle bool   `json:"trickle"`
	Sdp     string `json:"sdp"`
}

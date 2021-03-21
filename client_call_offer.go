package tdproto

// Start a call
type ClientCallOffer struct {
	BaseEvent
	Params clientCallOfferParams `json:"params"`
}

func (p ClientCallOffer) GetName() string { return "client.call.offer" }

type clientCallOfferParams struct {
	Jid     JID    `json:"jid"`
	Muted   bool   `json:"muted"`
	Trickle bool   `json:"trickle"`
	Sdp     string `json:"sdp"`
}

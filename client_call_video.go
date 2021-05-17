package tdproto

// Signal server about adding video to calls
type ClientCallVideo struct {
	BaseEvent
	Params clientCallVideoParams `json:"params"`
}

func (p ClientCallVideo) GetName() string { return "client.call.video" }

type clientCallVideoParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// SDP (session description protocol) data
	Sdp string `json:"sdp"`
}

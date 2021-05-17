package tdproto

// Signal server about answer
type ClientCallAnswer struct {
	BaseEvent
	Params clientCallAnswerParams `json:"params"`
}

func (p ClientCallAnswer) GetName() string { return "client.call.answer" }

type clientCallAnswerParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// SDP (session description protocol) data
	Sdp string `json:"sdp"`
}

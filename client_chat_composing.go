package tdproto

func NewClientChatComposing(jid JID, composing bool, draft *string) (r ClientChatComposing) {
	r.Name = r.GetName()
	r.Params.Jid = jid
	r.Params.Composing = composing
	r.Params.Draft = draft
	return r
}

// Typing or recording audiomessage
type ClientChatComposing struct {
	BaseEvent
	Params clientChatComposingParams `json:"params"`
}

func (p ClientChatComposing) GetName() string { return "client.chat.composing" }

// Params of the client.chat.composing event
type clientChatComposingParams struct {
	// Message draft data
	Draft *string `json:"draft,omitempty"`

	// Chat or contact id
	Jid JID `json:"jid"`

	// true = audiomessage, false = text typing
	IsAudio bool `json:"is_audio,omitempty"`

	// true = start typing / audio recording, false = stop
	Composing bool `json:"composing,omitempty"`
}

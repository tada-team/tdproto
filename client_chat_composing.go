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

type clientChatComposingParams struct {
	Jid       JID     `json:"jid"`
	IsAudio   bool    `json:"is_audio,omitempty"`
	Composing bool    `json:"composing,omitempty"`
	Draft     *string `json:"draft,omitempty"`
}

package tdproto

type ClientChatComposing struct {
	BaseEvent
	Params clientChatComposingParams `json:"params"`
}

type clientChatComposingParams struct {
	Jid       JID     `json:"jid"`
	IsAudio   bool    `json:"is_audio"`
	Composing bool    `json:"composing"`
	Draft     *string `json:"draft"`
}

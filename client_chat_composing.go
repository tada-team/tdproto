package tdproto

type ClientChatComposing struct {
	BaseEvent
	Params ClientChatComposingParams `json:"params"`
}

type ClientChatComposingParams struct {
	Jid       JID     `json:"jid"`
	IsAudio   bool    `json:"is_audio"`
	Composing bool    `json:"composing"`
	Draft     *string `json:"draft"`
}

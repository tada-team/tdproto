package tdproto

func NewServerChatComposing(composing, isAudio bool, chat, actor *JID) (r ServerChatComposing) {
	r.BaseEvent.Name = "server.chat.composing"
	r.Params.Jid = *chat
	r.Params.Actor = *actor
	r.Params.Composing = composing
	r.Params.IsAudio = isAudio
	return r
}

type ServerChatComposing struct {
	BaseEvent
	Params serverChatComposingParams `json:"params"`
}

type serverChatComposingParams struct {
	Jid       JID  `json:"jid"`
	Actor     JID  `json:"actor"`
	Composing bool `json:"composing"`
	IsAudio   bool `json:"is_audio,omitempty"`
}

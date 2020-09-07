package tdproto

func NewServerChatComposing(composing, isAudio bool, chat, actor *JID) (r ServerChatComposing) {
	r.BaseEvent.Unimportant = true
	r.BaseEvent.Name = "server.chat.composing"
	r.Params.Jid = *chat
	r.Params.Actor = *actor
	r.Params.Composing = composing
	r.Params.IsAudio = isAudio
	return r
}

type ServerChatComposing struct {
	BaseEvent
	Params ServerChatComposingParams `json:"params"`
}

type ServerChatComposingParams struct {
	Jid       JID  `json:"jid"`
	Actor     JID  `json:"actor"`
	Composing bool `json:"composing"`
	IsAudio   bool `json:"is_audio,omitempty"`
}

package tdproto

import "time"

func NewServerChatComposing(composing, isAudio bool, chat, actor *JID) (r ServerChatComposing) {
	r.Name = r.GetName()
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

func (p ServerChatComposing) GetName() string { return "server.chat.composing" }

type serverChatComposingParams struct {
	Jid        JID       `json:"jid"`
	Actor      JID       `json:"actor"`
	Composing  bool      `json:"composing"`
	IsAudio    bool      `json:"is_audio,omitempty"`
	ValidUntil time.Time `json:"valid_until,omitempty"`
}

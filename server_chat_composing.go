package tdproto

func NewServerChatComposing(composing, isAudio bool, chat, actor *JID) (r ServerChatComposing) {
	r.Name = r.GetName()
	r.Params.Jid = *chat
	r.Params.Actor = *actor
	r.Params.Composing = composing
	r.Params.IsAudio = isAudio
	return r
}

// Someone typing or recording audiomessage in chat
type ServerChatComposing struct {
	BaseEvent
	Params serverChatComposingParams `json:"params"`
}

func (p ServerChatComposing) GetName() string { return "server.chat.composing" }

type serverChatComposingParams struct {
	// Chat or contact id
	Jid JID `json:"jid"`

	// Actor id
	Actor JID `json:"actor"`

	// true = start typing / audio recording, false = stop
	Composing bool `json:"composing"`

	// true = audiomessage, false = text typing
	IsAudio bool `json:"is_audio,omitempty"`

	// Composing event max lifetime
	ValidUntil ISODateTimeString `json:"valid_until,omitempty"`
}

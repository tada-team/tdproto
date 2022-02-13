package tdproto

func NewClientChatLastread(jid JID, messageId *string) (r ClientChatLastread) {
	r.Name = r.GetName()
	r.ConfirmId = ConfirmId()
	r.Params.Jid = jid
	r.Params.LastReadMessageId = messageId
	return r
}

// Last read message in chat changed
type ClientChatLastread struct {
	Params clientChatLastreadParams `json:"params"`
	BaseEvent
}

func (p ClientChatLastread) GetName() string { return "client.chat.lastread" }

// Params of the client.chat.lastread event
type clientChatLastreadParams struct {
	// Last read message id. Omitted = last message in chat
	LastReadMessageId *string `json:"last_read_message_id,omitempty"`

	// Chat or contact id
	Jid JID `json:"jid"`
}

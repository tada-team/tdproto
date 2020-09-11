package tdproto

func NewClientChatLastread(jid JID, messageId *string) (r ClientChatLastread) {
	r.Name = "client.chat.lastread"
	r.ConfirmId = ConfirmId()
	r.Params.Jid = jid
	r.Params.LastReadMessageId = messageId
	return r
}

type ClientChatLastread struct {
	BaseEvent
	Params clientChatLastreadParams `json:"params"`
}

type clientChatLastreadParams struct {
	Jid               JID     `json:"jid"`
	LastReadMessageId *string `json:"last_read_message_id,omitempty"`
}

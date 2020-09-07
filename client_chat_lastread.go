package tdproto

type ClientChatLastread struct {
	BaseEvent
	Params ClientChatLastreadParams `json:"params"`
}

type ClientChatLastreadParams struct {
	Jid               JID     `json:"jid"`
	LastReadMessageId *string `json:"last_read_message_id"`
}

package tdproto

type ClientChatLastread struct {
	BaseEvent
	Params clientChatLastreadParams `json:"params"`
}

type clientChatLastreadParams struct {
	Jid               JID     `json:"jid"`
	LastReadMessageId *string `json:"last_read_message_id"`
}

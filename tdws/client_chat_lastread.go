package tdws

import "github.com/tada-team/tdproto"

func NewClientChatLastread(jid tdproto.JID, messageId *string) (r ClientChatLastread) {
	r.Name = r.GetName()
	r.ConfirmId = ConfirmId()
	r.Params.Jid = jid
	r.Params.LastReadMessageId = messageId
	return r
}

// Last read message in chat changed
type ClientChatLastread struct {
	BaseEvent
	Params clientChatLastreadParams `json:"params"`
}

func (p ClientChatLastread) GetName() string { return "client.chat.lastread" }

// Params of the client.chat.lastread event
type clientChatLastreadParams struct {
	// Chat or contact id
	Jid tdproto.JID `json:"jid"`

	// Last read message id. Omitted = last message in chat
	LastReadMessageId *string `json:"last_read_message_id,omitempty"`
}

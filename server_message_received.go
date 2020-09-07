package tdproto

type ReceivedMessage struct {
	Chat        JID    `json:"chat"`
	MessageId   string `json:"message_id"`
	Received    bool   `json:"received"`
	NumReceived int    `json:"num_received,omitempty"`
	Debug       string `json:"_debug,omitempty"`
}

func NewServerMessageReceived(messages []ReceivedMessage) (r ServerMessageReceived) {
	r.ConfirmId = ConfirmId()
	r.BaseEvent.Name = "server.message.received"
	r.Params.Messages = messages
	return r
}

type ServerMessageReceived struct {
	BaseEvent
	Params ServerMessageReceivedParams `json:"params"`
}

type ServerMessageReceivedParams struct {
	Messages []ReceivedMessage `json:"messages"`
}

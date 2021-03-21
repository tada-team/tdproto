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
	r.Name = r.GetName()
	r.Params.Messages = messages
	return r
}

// Message received by someone in chat
type ServerMessageReceived struct {
	BaseEvent
	Params serverMessageReceivedParams `json:"params"`
}

func (p ServerMessageReceived) GetName() string { return "server.message.received" }

type serverMessageReceivedParams struct {
	Messages []ReceivedMessage `json:"messages"`
}

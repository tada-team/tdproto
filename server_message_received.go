package tdproto

// Message receiving status
type ReceivedMessage struct {
	// Chat or contact id
	Chat JID `json:"chat"`

	// Message id
	MessageId string `json:"message_id"`

	// Is received
	Received bool `json:"received"`

	// Number of contacts received this message. Experimental.
	NumReceived int `json:"num_received,omitempty"`

	// Debug message, if any
	Debug string `json:"_debug,omitempty"`
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

// Params of the server.message.received event
type serverMessageReceivedParams struct {
	// received message data
	Messages []ReceivedMessage `json:"messages"`
}

package tdevents

func NewClientMessageDeleted(messageId string) (r ClientMessageDeleted) {
	r.Name = r.GetName()
	r.ConfirmId = ConfirmId()
	r.Params.MessageId = messageId
	return r
}

// Message deleted
type ClientMessageDeleted struct {
	BaseEvent
	Params clientMessageDeletedParams `json:"params"`
}

func (p ClientMessageDeleted) GetName() string { return "client.message.deleted" }

// Params of the client.message.deleted event
type clientMessageDeletedParams struct {
	// Message id
	MessageId string `json:"message_id,omitempty"`
}

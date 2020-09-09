package tdproto

func NewClientMessageDeleted(messageId string) (r ClientMessageDeleted) {
	r.Name = "client.message.deleted"
	r.ConfirmId = ConfirmId()
	r.Params.MessageId = messageId
	return r
}

type ClientMessageDeleted struct {
	BaseEvent
	Params clientMessageDeletedParams `json:"params"`
}

type clientMessageDeletedParams struct {
	MessageId string `json:"message_id,omitempty"`
}

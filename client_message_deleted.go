package tdproto

func NewClientMessageDeleted(messageId string) (r ClientMessageDeleted) {
	r.Name = r.GetName()
	r.ConfirmId = ConfirmId()
	r.Params.MessageId = messageId
	return r
}

type ClientMessageDeleted struct {
	BaseEvent
	Params clientMessageDeletedParams `json:"params"`
}

func (p ClientMessageDeleted) GetName() string { return "client.message.deleted" }

type clientMessageDeletedParams struct {
	MessageId string `json:"message_id,omitempty"`
}

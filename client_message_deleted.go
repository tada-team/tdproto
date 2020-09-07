package tdproto

type ClientMessageDeleted struct {
	BaseEvent
	Params ClientMessageDeletedParams `json:"params"`
}

type ClientMessageDeletedParams struct {
	MessageId string `json:"message_id,omitempty"`
}

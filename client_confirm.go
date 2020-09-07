package tdproto

type ClientConfirm struct {
	BaseEvent
	Params ClientConfirmParams `json:"params"`
}

type ClientConfirmParams struct {
	MessageId string `json:"message_id,omitempty"`
	ConfirmId string `json:"confirm_id"`
}

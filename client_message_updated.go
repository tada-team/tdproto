package tdproto

func NewClientMessageUpdated(p ClientMessageUpdatedParams) (r ClientMessageUpdated) {
	r.Name = "client.message.updated"
	r.ConfirmId = ConfirmId()
	r.Params = p
	return r
}

type ClientMessageUpdated struct {
	BaseEvent
	Params ClientMessageUpdatedParams `json:"params"`
}

type ClientMessageUpdatedParams struct {
	// chat, task or contact jid. Required.
	To JID `json:"to"`

	// message content. Required.
	Content MessageContent `json:"content"`

	// uid created by client. Recommended.
	MessageId string `json:"message_id,omitempty"`

	// Replied to message id. Not required.
	ReplyTo string `json:"reply_to,omitempty"`

	// forwarded messages (previously was for reply too). Not required.
	LinkedMessages []string `json:"linked_messages,omitempty"`

	// important flag. Not required. Default: false
	Important bool `json:"important,omitempty"`

	// disable links preview generation. Not required. Default: false
	Nopreview bool `json:"nopreview,omitempty"`

	// deprecated
	Comment string `json:"comment,omitempty"`
}

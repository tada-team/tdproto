package tdproto

func NewClientMessageUpdated(p ClientMessageUpdatedParams) (r ClientMessageUpdated) {
	r.Name = r.GetName()
	r.ConfirmId = ConfirmId()
	r.Params = p
	return r
}

type ClientMessageUpdated struct {
	BaseEvent
	Params ClientMessageUpdatedParams `json:"params"`
}

func (p ClientMessageUpdated) GetName() string { return "client.message.updated" }

type ClientMessageUpdatedParams struct {
	// Chat, task or contact jid. Required.
	To JID `json:"to"`

	// Message content. Required.
	Content MessageContent `json:"content"`

	// Uid created by client. Recommended.
	MessageId string `json:"message_id,omitempty"`

	// Replied to message id. Not required.
	ReplyTo string `json:"reply_to,omitempty"`

	// Forwarded messages (previously was for reply too). Not required.
	LinkedMessages []string `json:"linked_messages,omitempty"`

	// Important flag. Not required. Default: false
	Important bool `json:"important,omitempty"`

	// Disable links preview generation. Not required. Default: false
	Nopreview bool `json:"nopreview,omitempty"`

	// Message attachments
	Uploads []string `json:"uploads,omitempty"`

	// Backward compatibility mode
	OldStyleAttachment bool `json:"old_style_attachment,omitempty"`

	// Deprecated
	Comment string `json:"comment,omitempty"`
}

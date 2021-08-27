package tdforms

import "github.com/tada-team/tdproto"

// Message creation form
type MessageCreation struct {
	// Message type
	Type tdproto.Mediatype `json:"type"`

	// Message text
	Text string `json:"text"`

	// Message id
	MessageUid string `json:"message_id,omitempty"`

	// Message attachments
	Uploads []string `json:"uploads,omitempty"`

	// Backward compatibility mode
	OldStyleAttachment bool `json:"old_style_attachment,omitempty"`

	MessageUpdate
}

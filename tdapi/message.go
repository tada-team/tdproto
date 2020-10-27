package tdapi

import "github.com/tada-team/tdproto"

type MessageUpdate struct {
	// Important flag. Not required. Default: false
	Important bool `json:"important,omitempty"`

	// Disable links preview generation. Not required. Default: false
	Nopreview bool `json:"nopreview,omitempty"`
}

type Message struct {
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

// file upload:
// <form action="...?message_id=..." method="post" enctype="multipart/form-data">
//   <input type="file" name="file">
// </form>

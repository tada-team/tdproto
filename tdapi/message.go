package tdapi

import "github.com/tada-team/tdproto"

type MessageUpdate struct {
	Important bool `json:"important,omitempty"`
	Nopreview bool `json:"nopreview,omitempty"`
}

type Message struct {
	Mediatype  tdproto.Mediatype `json:"type"`
	Text       string            `json:"text"`
	MessageUid string            `json:"message_id,omitempty"`
	MessageUpdate
}

// file upload:
// <form action="...?message_id=..." method="post" enctype="multipart/form-data">
//   <input type="file" name="file">
// </form>

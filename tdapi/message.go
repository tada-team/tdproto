package tdapi

import "github.com/tada-team/tdproto"

type MessageFilter struct {
	UserParams
	Paginator

	// ?chat=jid,jid
	Chat string `schema:"chat"`

	// ?chat_type=task,group,direct|any (default: any)
	ChatType string `schema:"chat_type"`

	// ?sender=jid,jid
	Sender string `schema:"sender"`

	// ?has_upload=true|false|any (default: any)
	HasUpload string `schema:"has_upload"`

	// ?text=substr
	Text string `schema:"text"`

	// ?type=image,video,plain,file
	Type tdproto.Mediatype `schema:"type"`

	// ?important=true|any
	Important string `schema:"important"`

	// ?date_from=dt (include)
	DateFrom string `schema:"date_from"`

	// ?date_to=dt
	DateTo string `schema:"date_to"`

	// ?include_deleted=true|false (default: false)
	IncludeDeleted string `schema:"include_deleted"`

	// ?gentime_from=
	GentimeFrom string `schema:"gentime_from"`

	// ?exact=msgId
	Exact string `schema:"exact"`

	// ?unread=(true|false)
	Unread string `schema:"unread"`

	// ?old_from=msgId (exclude msgId)
	OldFrom string `schema:"old_from"`

	// ?new_from=msgId (exclude msgId)
	NewFrom string `schema:"new_from"`

	// ?old_from_inc=msgId (include msgId)
	OldFromInc string `schema:"old_from_inc"`

	// ?new_from_inc=msgId (include msgId)
	NewFromInc string `schema:"new_from_inc"`

	// ?around=msgId (include msgId)
	Around string `schema:"around"`
}

type MessageUpdate struct {
	// Important flag. Not required. Default: false
	Important bool `json:"important,omitempty"`

	// Disable links preview generation. Not required. Default: false
	Nopreview bool `json:"nopreview,omitempty"`

	// Draft message, send later
	SendAt string `json:"send_at,omitempty"`
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

// Message markup with checked links
type MessageMarkup struct {
	// Message markup
	Markup []tdproto.MarkupEntity `json:"markup"`

	// Deprecated: use markup instead
	Links tdproto.MessageLinks `json:"links"`
}

// file upload:
// <form action="...?message_id=..." method="post" enctype="multipart/form-data">
//   <input type="file" name="file">
// </form>

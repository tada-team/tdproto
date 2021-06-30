package tdapi

import "github.com/tada-team/tdproto"

// Uploads filter
type UploadFilter struct {
	Paginator

	// ?chat=jid,jid
	Chat string `schema:"chat"`

	// ?sender=jid,jid
	Sender string `schema:"sender"`

	// ?content_type=image/jpeg
	ContentType string `schema:"content_type"`

	// ?type=file,image,audio,video
	Type tdproto.Mediatype `schema:"type"`

	// ?text=substr
	Text string `schema:"text"`
}

// Upload + ShortMessage objects
type UploadShortMessages struct {
	// Upload + ShortMessage objects list
	Objects []tdproto.UploadShortMessage `json:"objects"`

	// Count
	Count int `json:"count"`

	// Limit
	Limit int `json:"limit"`

	// Offset
	Offset int `json:"offset"`
}

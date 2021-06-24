package tdapi

import "github.com/tada-team/tdproto"

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

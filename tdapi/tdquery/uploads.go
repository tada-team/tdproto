package tdquery

import "github.com/tada-team/tdproto"

// Uploads filter
type Uploads struct {
	Paginator

	// Comma separated chat jids
	Chat string `schema:"chat"`

	// Comma separated sender jids
	Sender string `schema:"sender"`

	// Content type
	ContentType string `schema:"content_type"`

	// Comma separated Mediatypes
	Type tdproto.Mediatype `schema:"type"`

	// Search string
	Text string `schema:"text"`
}

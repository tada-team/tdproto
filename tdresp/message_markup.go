package tdresp

import "github.com/tada-team/tdproto"

// Message markup with checked links
type MessageMarkup struct {
	// Message markup
	Markup []tdproto.MarkupEntity `json:"markup"`

	// Deprecated: use markup instead
	Links tdproto.MessageLinks `json:"links"`
}

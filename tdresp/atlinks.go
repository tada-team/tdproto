package tdresp

import "github.com/tada-team/tdproto"

// @-links autocomplete response
type AtLinks []AtLink

// @-link autocomplete information
type AtLink struct {
	// What should be inserted to the chat
	Key string `json:"key"`

	// What should be visible by user
	Title string `json:"title"`

	// Hint for user, if any
	Help string `json:"help,omitempty"`

	// Internal details, if any
	Meta *AtLinkMeta `json:"meta,omitempty"`
}

// @-link autocomplete details
type AtLinkMeta struct {
	// Contact or section jid
	Jid tdproto.JID `json:"jid"`

	// Section uid (for contact sections only)
	Section string `json:"section,omitempty"`
}

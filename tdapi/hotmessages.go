package tdapi

import "github.com/tada-team/tdproto"

// All-in-one response with all messages info for fast chat rendering
type Hotmessages struct {
	// chat information
	Chat tdproto.Chat `json:"chat"`

	// Messages list
	Messages []tdproto.Message `json:"messages"`

	// Contacts list
	Contacts []tdproto.Contact `json:"contacts"`
}

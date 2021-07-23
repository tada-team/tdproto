package tdresp

import "github.com/tada-team/tdproto"

// Chats response
type Chats struct {
	// Chat list
	Objects []tdproto.Chat `json:"objects"`

	// Contacts used in chats
	Contacts []tdproto.Contact `json:"contacts,omitempty"`

	// Count
	Count int `json:"count"`

	// Limit
	Limit int `json:"limit"`

	// Offset
	Offset int `json:"offset"`
}

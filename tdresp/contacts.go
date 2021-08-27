package tdresp

import "github.com/tada-team/tdproto"

// Paginated contacts
type Contacts struct {
	// Contacts list
	Objects []tdproto.Contact `json:"objects"`

	// Count
	Count int `json:"count"`

	// Limit
	Limit int `json:"limit"`

	// Offset
	Offset int `json:"offset"`
}

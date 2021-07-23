package tdresp

import "github.com/tada-team/tdproto"

// Messages response
type Messages struct {
	// Message list
	Objects []tdproto.Message `json:"objects"`

	// Count
	Count int `json:"count"`

	// Limit
	Limit int `json:"limit"`

	// Offset
	Offset int `json:"offset"`
}


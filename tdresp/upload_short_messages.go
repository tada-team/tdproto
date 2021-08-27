package tdresp

import "github.com/tada-team/tdproto"

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

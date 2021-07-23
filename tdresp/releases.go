package tdresp

import "github.com/tada-team/tdproto"

// Releases response
type Releases struct {
	// Android dist
	Android []tdproto.Dist `json:"android,omitempty"`

	// Linux dist
	Linux []tdproto.Dist `json:"linux,omitempty"`

	// macOS dist
	Mac []tdproto.Dist `json:"mac,omitempty"`

	// Windows dist
	Win []tdproto.Dist `json:"win,omitempty"`
}

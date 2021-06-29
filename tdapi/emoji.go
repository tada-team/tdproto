package tdapi

import "github.com/tada-team/tdproto"

// Emoji response
type Emoji struct {
	// Emoji list
	All []tdproto.Emoji `json:"all"`
}

package tdresp

import "github.com/tada-team/tdproto"

// Reactions to messages: frequently used and all available
type MyReactions struct {
	// My frequently used reactions
	Top []tdproto.Reaction `json:"top"`

	// All available reactions
	All []tdproto.Reaction `json:"all"`
}

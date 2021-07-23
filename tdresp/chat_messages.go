package tdresp

import "github.com/tada-team/tdproto"

// Chat messages
type ChatMessages struct {
	// Message list
	Messages []tdproto.Message `json:"messages"`
}

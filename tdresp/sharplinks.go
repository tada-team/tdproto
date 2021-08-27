package tdresp

import "github.com/tada-team/tdproto"

// #-links autocomplete response
type SharpLinks []SharpLink

// #-link autocomplete information
type SharpLink struct {
	// What should be inserted to the chat
	Key string `json:"key"`

	// What should be visible by user
	Title string `json:"title"`

	// Icon data, if any
	Icons *tdproto.IconData `json:"icons,omitempty"`

	// Internal details
	Meta SharpLinkMeta `json:"meta"`
}

// #-link autocomplete details
type SharpLinkMeta struct {
	// Chat id
	Jid tdproto.JID `json:"jid"`

	// Chat type
	ChatType tdproto.ChatType `json:"chat_type"`

	// Is task or group public for non-guests
	Public bool `json:"public,omitempty"`

	// Task status (for tasks)
	TaskStatus string `json:"task_status,omitempty"`

	// Task number (for tasks)
	Num uint `json:"num,omitempty"`

	// Deprecated: use `TaskStatus == "done"` comparsion
	Done bool `json:"done,omitempty"`
}

package tdproto

import (
	"fmt"
)

// Markup entity. Experimental
type MarkupEntity struct {
	// Text replacement.
	Repl string `json:"repl,omitempty"`

	// Time, for Time type
	Time string `json:"time,omitempty"`

	// Marker type
	Type MarkupType `json:"typ"`

	// Url, for Link type
	Url string `json:"url,omitempty"`

	// List of internal markup entities
	Childs []MarkupEntity `json:"childs,omitempty"`

	// Close marker length
	CloseLength int `json:"cllen,omitempty"`

	// Open marker offset
	Open int `json:"op"`

	// Open marker length
	OpenLength int `json:"oplen,omitempty"`

	// Close marker offset
	Close int `json:"cl"`
}

func (e MarkupEntity) String() string {
	return fmt.Sprintf("%d..+%d %d..+%d %s", e.Open, e.OpenLength, e.Close, e.CloseLength, e.Type)
}

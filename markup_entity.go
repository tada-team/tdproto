package tdproto

import (
	"fmt"
)

const (
	Bold       = "bold"
	Italic     = "italic"
	Underscore = "underscore"
	Strike     = "strike"
	Code       = "code"
	CodeBlock  = "codeblock"
	Quote      = "quote"
	Link       = "link"
	Time       = "time"
	Unsafe     = "unsafe"
)

// Markup entity. Experimental
type MarkupEntity struct {
	// Open marker offset
	Open int `json:"op"`

	// Open marker length
	OpenLength int `json:"oplen,omitempty"`

	// Close marker offset
	Close int `json:"cl"`

	// Close marker length
	CloseLength int `json:"cllen,omitempty"`

	// Marker type
	Type string `json:"typ"`

	// Url, for Link type
	Url string `json:"url,omitempty"`

	// Text replacement.
	Repl string `json:"repl,omitempty"`

	// Time, for Time type
	Time string `json:"time,omitempty"`

	// List of internal markup entities
	Childs []MarkupEntity `json:"childs,omitempty"`
}

func (e MarkupEntity) String() string {
	return fmt.Sprintf("%d..+%d %d..+%d %s", e.Open, e.OpenLength, e.Close, e.CloseLength, e.Type)
}

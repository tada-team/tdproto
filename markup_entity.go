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
	Unsafe     = "unsafe"
)

// Markup entity. Experimental
type MarkupEntity struct {
	// Open marker offset
	Open int `json:"open"`

	// Open marker length
	OpenLength int `json:"open_length,omitempty"`

	// Close marker offset
	Close int `json:"close"`

	// Close marker length
	CloseLength int `json:"close_length,omitempty"`

	// Marker type
	Type string `json:"type"`

	// Link, if any
	Url string `json:"url,omitempty"`

	// List of internal markup entities
	Entities []MarkupEntity `json:"entities,omitempty"`
}

func (e MarkupEntity) String() string {
	return fmt.Sprintf("%d..+%d %d..+%d %s", e.Open, e.OpenLength, e.Close, e.CloseLength, e.Type)
}

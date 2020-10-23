package tdproto

import (
	"fmt"
	"time"
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
	Open int `json:"open"`

	// Open marker length
	OpenLength int `json:"open_length,omitempty"`

	// Close marker offset
	Close int `json:"close"`

	// Close marker length
	CloseLength int `json:"close_length,omitempty"`

	// Marker type
	Type string `json:"type"`

	// Url, for Link type
	Url string `json:"url,omitempty"`

	// Time, for Time type
	Time *time.Time `json:"time,omitempty"`

	// List of internal markup entities
	Entities []MarkupEntity `json:"entities,omitempty"`
}

func (e MarkupEntity) String() string {
	return fmt.Sprintf("%d..+%d %d..+%d %s", e.Open, e.OpenLength, e.Close, e.CloseLength, e.Type)
}

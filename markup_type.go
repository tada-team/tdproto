package tdproto

// Markup type
type MarkupType string

const (
	// Bold text
	Bold MarkupType = "bold"

	// Italic text
	Italic MarkupType = "italic"

	// Underscore text
	Underscore MarkupType = "underscore"

	// Striked text
	Strike MarkupType = "strike"

	// Inlined code
	Code MarkupType = "code"

	// Code block
	CodeBlock MarkupType = "codeblock"

	// Quote
	Quote MarkupType = "quote"

	// Link
	Link MarkupType = "link"

	// Datetime
	Time MarkupType = "time"

	// Unsafe html element
	Unsafe MarkupType = "unsafe"
)

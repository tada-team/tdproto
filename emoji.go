package tdproto

// Emoji
type Emoji struct {
	// Unicode symbol
	Char string `json:"char"`

	// Text representation
	Key  string `json:"key"`
}

package tdforms

// Update message form
type MessageUpdate struct {
	// Important flag. Not required. Default: false
	Important bool `json:"important,omitempty"`

	// Disable links preview generation. Not required. Default: false
	Nopreview bool `json:"nopreview,omitempty"`

	// Draft message, send later
	SendAt string `json:"send_at,omitempty"`
}

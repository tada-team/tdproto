package tdproto

// Small or large icon
type SingleIcon struct {
	// absolute url to icon
	Url string `json:"url"`

	// Icon width, in pixels
	Width int `json:"width"`

	// Icon height, in pixels
	Height int `json:"height"`
}

// Icon data. Contains sm+lg (for uploaded image) OR stub+letters+color (for icon generated from display name)
type IconData struct {
	// Small icon
	Sm *SingleIcon `json:"sm,omitempty"`

	// Large image
	Lg *SingleIcon `json:"lg,omitempty"`

	// Generated image with 1-2 letters
	Stub string `json:"stub,omitempty"`

	// Letters from stub icon
	Letters string `json:"letters,omitempty"`

	// Stub icon background color
	Color string `json:"color,omitempty"`
}

func (d *IconData) SmUrlOrStub() string {
	if d.Sm != nil {
		return d.Sm.Url
	}
	return d.Stub
}

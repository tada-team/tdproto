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

// Icon data. For icon generated from display name contains Letters + Color fields.
type IconData struct {
	// Letters (only for stub icon)
	Letters string `json:"letters,omitempty"`

	// Icon background color (only for stub icon)
	Color string `json:"color,omitempty"`

	// Compact representation of a placeholder for an image (experimental)
	Blurhash string `json:"blurhash,omitempty"`

	// Deprecated
	Stub string `json:"stub,omitempty"`

	// Small icon
	Sm SingleIcon `json:"sm"`

	// Large image
	Lg SingleIcon `json:"lg"`
}

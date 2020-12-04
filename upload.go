package tdproto

// Uploaded media
type Upload struct {
	// Upload id
	Uid string `json:"uid"`

	// Uploaded at
	Created DateTime `json:"created"`

	// Upload size in bytes
	Size int `json:"size"`

	// Mediafile duration (for audio/video only)
	Duration uint `json:"duration,omitempty"`

	// Filename
	Name string `json:"name"`

	// Absolute url
	Url string `json:"url"`

	// Preview details
	Preview *UploadPreview `json:"preview,omitempty"`

	// Content type
	ContentType string `json:"content_type"`

	// Is animated (images only)
	Animated bool `json:"animated,omitempty"`

	// File still processing (video only)
	Processing bool `json:"processing,omitempty"`

	// PDF version of file. Experimental
	PdfVersion *PdfVersion `json:"pdf_version,omitempty"`
}

// Upload preview
type UploadPreview struct {
	// Absolute url to image
	Url string `json:"url"`

	// Absolute url to high resolution image (retina)
	Url2x string `json:"url_2x"`

	// Width in pixels
	Width int `json:"width"`

	// Height in pixels
	Height int `json:"height"`
}

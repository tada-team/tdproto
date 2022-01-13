package tdproto

// Uploaded media
type Upload struct {
	// Preview details
	Preview *UploadPreview `json:"preview,omitempty"`

	// PDF version of file. Experimental
	PdfVersion *PdfVersion `json:"pdf_version,omitempty"`

	// Uploaded at
	Created ISODateTimeString `json:"created"`

	// Compact representation of a placeholder for an image (images only)
	Blurhash string `json:"blurhash,omitempty"`

	// Filename
	Name string `json:"name"`

	// Absolute url
	Url string `json:"url"`

	// Content type
	ContentType string `json:"content_type"`

	// Upload id
	Uid string `json:"uid"`

	// ?type=file,image,audio,video
	MediaType UploadMediaType `json:"type"`

	// Mediafile duration (for audio/video only)
	Duration uint `json:"duration,omitempty"`

	// Upload size in bytes
	Size int `json:"size"`

	// Is animated (images only)
	Animated bool `json:"animated,omitempty"`

	// File still processing (video only)
	Processing bool `json:"processing,omitempty"`
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

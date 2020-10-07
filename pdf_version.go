package tdproto

// PDF preview of mediafile. Experimental
type PdfVersion struct {
	// Absolute url
	Url string `json:"url"`

	// First string of text content
	TextPreview string `json:"text_preview,omitempty"`
}

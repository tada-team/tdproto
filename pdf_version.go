package tdproto

type PdfVersion struct {
	Url         string `json:"url"`
	TextPreview string `json:"text_preview,omitempty"`
}

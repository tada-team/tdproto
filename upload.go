package tdproto

type Upload struct {
	Uid         string         `json:"uid"`
	Size        int            `json:"size"`
	Duration    uint           `json:"duration,omitempty"`
	Name        string         `json:"name"`
	Url         string         `json:"url"`
	Preview     *UploadPreview `json:"preview,omitempty"`
	ContentType string         `json:"content_type"`
	Animated    bool           `json:"animated,omitempty"`
	Processing  bool           `json:"processing,omitempty"`
	PdfVersion  *PdfVersion    `json:"pdf_version,omitempty"`
}

type UploadPreview struct {
	Url    string `json:"url"`
	Url2x  string `json:"url_2x"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

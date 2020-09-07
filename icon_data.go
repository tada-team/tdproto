package tdproto

type SingleIcon struct {
	Url    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height"`
}

type IconData struct {
	Sm      *SingleIcon `json:"sm,omitempty"`
	Lg      *SingleIcon `json:"lg,omitempty"`
	Stub    string      `json:"stub,omitempty"`
	Letters string      `json:"letters,omitempty"`
	Color   string      `json:"color,omitempty"`
}

func (d *IconData) SmUrlOrStub() string {
	if d.Sm != nil {
		return d.Sm.Url
	}
	return d.Stub
}

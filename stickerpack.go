package tdproto

// Collection of stickers
type Stickerpack struct {
	Uid        string    `json:"uid"`
	Name       string    `json:"name"`
	Title      string    `json:"title"`
	Author     string    `json:"author,omitempty"`
	AuthorLink string    `json:"author_link,omitempty"`
	Stickers   []Sticker `json:"stickers"`
}

// Single sticker
type Sticker struct {
	Uid            string         `json:"uid"`
	Icon64         string         `json:"icon64"`
	Icon100        string         `json:"icon100"`
	Icon128        string         `json:"icon128"`
	Icon200        string         `json:"icon200"`
	MessageContent MessageContent `json:"message_content"`
}

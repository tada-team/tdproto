package tdproto

// Chat wallpaper
type Wallpaper struct {
	// Localized description
	Name string `json:"name"`

	// Url to jpg or png
	Url  string `json:"url"`
}

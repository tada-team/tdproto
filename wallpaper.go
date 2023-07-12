package tdproto

// Chat wallpaper
type Wallpaper struct {
	// Unique identifier
	Key string `json:"key"`

	// Localized description
	Name string `json:"name"`

	// Url to jpg or png
	Url string `json:"url"`
	
	// test description
	Description string `json:"description"`
}

package tdproto

// Task color rules color
type TaskColor struct {
	// Regular color
	Regular RGBColor `json:"regular"`

	// Dark color
	Dark RGBColor `json:"dark"`

	// Light color
	Light RGBColor `json:"light"`
}

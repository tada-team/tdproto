package tdproto

// Task color rules color
type TaskColor struct {
	// Regular color (HEX)
	Regular string `json:"regular"`

	// Dark color (HEX)
	Dark    string `json:"dark"`

	// Light color (HEX)
	Light   string `json:"light"`
}

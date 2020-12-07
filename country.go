package tdproto

// Country for phone numbers selection on login screen
type Country struct {
	// Country code
	Code string `json:"code"`

	// Country name
	Name string `json:"name"`

	// Selected by default
	Default bool `json:"default,omitempty"`

	// Is popular, need to cache
	Popular bool `json:"popular,omitempty"`
}

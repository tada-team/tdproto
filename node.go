package tdproto

// Node (for external users)
type Node struct {
	// Node uid
	Uid string `json:"uid"`

	// Node title
	Title string `json:"title"`

	// Synchronization with node works
	Enabled bool `json:"enabled"`
}

package tdproto

// Custom task status
type TaskStatus struct {
	// Status id
	Uid string `json:"uid,omitempty"`

	// Status internal name
	Name string `json:"name"`

	// Status localized name
	Title string `json:"title"`

	// Status sort ordering
	SortOrdering uint `json:"sort_ordering"`

	// Status not used anymore
	IsArchive bool `json:"is_archive,omitempty"`
}

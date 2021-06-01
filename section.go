package tdproto

// Task project or contact section
type Section struct {
	// Section uid
	Uid string `json:"uid"`

	// Sort ordering
	SortOrdering uint `json:"sort_ordering"`

	// Name
	Name string `json:"name"`

	// Object version
	Gentime int64 `json:"gentime"`

	// Description, if any
	Description string `json:"description,omitempty"`

	// Is deleted
	IsArchive bool `json:"is_archive,omitempty"`
}

// Deleted task project or contact section
type DeletedSection struct {
	// Section uid
	Uid string `json:"uid"`

	// Object version
	Gentime int64 `json:"gentime"`
}

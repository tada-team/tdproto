package tdproto

// Task tag
type Tag struct {
	// Tag id
	Uid string `json:"uid"`

	// Tag name
	Name string `json:"name"`
}

// Delete tag message
type DeletedTag struct {
	// Tag id
	Uid string `json:"uid"`
}

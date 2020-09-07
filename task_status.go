package tdproto

type TaskStatus struct {
	Uid          string `json:"uid,omitempty"`
	SortOrdering uint   `json:"sort_ordering"`
	Name         string `json:"name"`
	Title        string `json:"title"`
	IsArchive    bool   `json:"is_archive,omitempty"`
}

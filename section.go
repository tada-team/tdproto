package tdproto

type Section struct {
	Uid          string `json:"uid"`
	SortOrdering uint   `json:"sort_ordering"`
	Name         string `json:"name"`
	Gentime      int64  `json:"gentime"`
	Description  string `json:"description,omitempty"`
	IsArchive    bool   `json:"is_archive,omitempty"`
}

type DeletedSection struct {
	Uid string `json:"uid"`
}

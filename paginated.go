package tdproto

type PaginatedChats struct {
	Contacts []Contact `json:"contacts,omitempty"`
	Objects  []Chat    `json:"objects"`
	Count    int       `json:"count"`
	Limit    int       `json:"limit"`
	Offset   int       `json:"offset"`
}

type PaginatedMessages struct {
	Objects []Message `json:"objects"`
	Count   int       `json:"count"`
	Limit   int       `json:"limit"`
	Offset  int       `json:"offset"`
}

type PaginatedContacts struct {
	Objects []Contact `json:"objects"`
	Count   int       `json:"count"`
	Limit   int       `json:"limit"`
	Offset  int       `json:"offset"`
}

type ChatMessages struct {
	Messages []Message `json:"messages"`
}

type PaginatedUploadShortMessages struct {
	Objects []UploadShortMessage `json:"objects"`
	Count   int                  `json:"count"`
	Limit   int                  `json:"limit"`
	Offset  int                  `json:"offset"`
}

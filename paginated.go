package tdproto

// Paginated chats
type PaginatedChats struct {
	Contacts []Contact `json:"contacts,omitempty"`
	Objects  []Chat    `json:"objects"`
	Count    int       `json:"count"`
	Limit    int       `json:"limit"`
	Offset   int       `json:"offset"`
}

// Paginated messages
type PaginatedMessages struct {
	Objects []Message `json:"objects"`
	Count   int       `json:"count"`
	Limit   int       `json:"limit"`
	Offset  int       `json:"offset"`
}

// Paginated contacts
type PaginatedContacts struct {
	Objects []Contact `json:"objects"`
	Count   int       `json:"count"`
	Limit   int       `json:"limit"`
	Offset  int       `json:"offset"`
}

// Chat messages
type ChatMessages struct {
	Messages []Message `json:"messages"`
}

// Paginated UploadShortMessage
type PaginatedUploadShortMessages struct {
	Objects []UploadShortMessage `json:"objects"`
	Count   int                  `json:"count"`
	Limit   int                  `json:"limit"`
	Offset  int                  `json:"offset"`
}

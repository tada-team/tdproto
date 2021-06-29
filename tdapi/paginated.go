package tdapi

import "github.com/tada-team/tdproto"

// Paginated chats
type PaginatedChats struct {
	Contacts []tdproto.Contact `json:"contacts,omitempty"`
	Objects  []tdproto.Chat    `json:"objects"`
	Count    int                `json:"count"`
	Limit    int                `json:"limit"`
	Offset   int                `json:"offset"`
}

// Paginated messages
type PaginatedMessages struct {
	Objects []tdproto.Message `json:"objects"`
	Count   int                `json:"count"`
	Limit   int                `json:"limit"`
	Offset  int                `json:"offset"`
}

// Paginated contacts
type PaginatedContacts struct {
	Objects []tdproto.Contact `json:"objects"`
	Count   int                `json:"count"`
	Limit   int                `json:"limit"`
	Offset  int                `json:"offset"`
}

// Chat messages
type ChatMessages struct {
	Messages []tdproto.Message `json:"messages"`
}

// Paginated UploadShortMessage
type PaginatedUploadShortMessages struct {
	Objects []tdproto.UploadShortMessage `json:"objects"`
	Count   int                           `json:"count"`
	Limit   int                           `json:"limit"`
	Offset  int                           `json:"offset"`
}

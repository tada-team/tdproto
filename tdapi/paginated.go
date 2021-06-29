package tdapi

import "github.com/tada-team/tdproto/tdmodels"

// Paginated chats
type PaginatedChats struct {
	Contacts []tdmodels.Contact `json:"contacts,omitempty"`
	Objects  []tdmodels.Chat    `json:"objects"`
	Count    int                `json:"count"`
	Limit    int                `json:"limit"`
	Offset   int                `json:"offset"`
}

// Paginated messages
type PaginatedMessages struct {
	Objects []tdmodels.Message `json:"objects"`
	Count   int                `json:"count"`
	Limit   int                `json:"limit"`
	Offset  int                `json:"offset"`
}

// Paginated contacts
type PaginatedContacts struct {
	Objects []tdmodels.Contact `json:"objects"`
	Count   int                `json:"count"`
	Limit   int                `json:"limit"`
	Offset  int                `json:"offset"`
}

// Chat messages
type ChatMessages struct {
	Messages []tdmodels.Message `json:"messages"`
}

// Paginated UploadShortMessage
type PaginatedUploadShortMessages struct {
	Objects []tdmodels.UploadShortMessage `json:"objects"`
	Count   int                           `json:"count"`
	Limit   int                           `json:"limit"`
	Offset  int                           `json:"offset"`
}

package tdapi

import "github.com/tada-team/tdproto"

// ParserUploadArchiveResponse response structure for method UploadArchive
type ParserUploadArchiveResponse struct {
	// Success result
	Success bool `json:"success"`

	// ProcessingAction action for background notifications about archive unpacking.
	ProcessingAction string `json:"processing_action"`

	// ActionType must be archive_unpacking
	ActionType tdproto.ActionType `json:"action_type"`

	// ArchiveName name of archive
	ArchiveName string `json:"archive_name"`
}

// ParserGetStateResponse response structure for method GetArchiveState
type ParserGetStateResponse struct {
	// State of import chats
	State tdproto.ParseState `json:"state"`

	// Progress of archive unpacking
	Progress uint16 `json:"progress,omitempty"`

	// Action name
	Action string `json:"action,omitempty"`

	// ActionType. Ex: [archive_unpacking || generate_chats]
	ActionType tdproto.ActionType `json:"action_type,omitempty"`

	// Localized Message
	Message string `json:"message,omitempty"`

	// Localized Body
	Body string `json:"body,omitempty"`

	// ArchiveName name of archive
	ArchiveName string `json:"archive_name"`

	// Has error
	HasError bool `json:"has_error"`
}

// ParserSendArchiveStatusRequest ...
type ParserSendArchiveStatusRequest struct {
	// Status archive parse status
	Status tdproto.ParseStatus `json:"status"`

	// Progress of archive unpacking
	Progress uint16 `json:"progress"`

	// ErrorCode if smth went wrong
	ErrorCode tdproto.ParseErrCode `json:"error_code,omitempty"`
}

// ParserGetMappedUsersResponse ...
type ParserGetMappedUsersResponse struct {
	// Users ...
	Users []tdproto.MappedUser `json:"users"`

	// ChatName ...
	ChatName string `json:"chat_name"`
}

// ParserMapUsersRequest ...
type ParserMapUsersRequest struct {
	// Users ...
	Users []tdproto.MappedUser `json:"users"`
}

// ParserMapUsersResponse ...
type ParserMapUsersResponse struct {
	// Success result
	Success bool `json:"success"`
}

// ParserGenerateChatsResponse ...
type ParserGenerateChatsResponse struct {
	// ProcessingAction action for background notifications about generation of chats and messages.
	ProcessingAction string `json:"processing_action"`

	// ActionType must be generate_chat
	ActionType tdproto.ActionType `json:"action_type"`

	// ArchiveName name of archive
	ArchiveName string `json:"archive_name"`
}

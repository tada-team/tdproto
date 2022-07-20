package tdapi

import "github.com/tada-team/tdproto"

// ParserUploadArchiveRequest request structure for method UploadArchive, which take and parse archive
type ParserUploadArchiveRequest struct {
}

// ParserUploadArchiveResponse response structure for method UploadArchive
type ParserUploadArchiveResponse struct {
	// Success result
	Success bool `json:"success"`

	// ProcessingAction action for background notifications about archive unpacking.
	ProcessingAction string `json:"processing_action"`

	// ActionType must be archive_unpacking
	ActionType tdproto.ActionType `json:"action_type"`
}

// ParserGetArchiveStatusRequest request structure for method GetArchiveStatus
type ParserGetArchiveStatusRequest struct {
}

// ParserGetArchiveStatusResponse response structure for method GetArchiveStatus
type ParserGetArchiveStatusResponse struct {
	// Status archive parse status
	Status tdproto.ParseStatus `json:"status"`
}

// ParserSendArchiveStatusRequest ...
type ParserSendArchiveStatusRequest struct {
	Status   tdproto.ParseStatus `json:"status"`
	Progress uint16              `json:"progress"`
}

// ParserSendArchiveStatusResponse ...
type ParserSendArchiveStatusResponse struct {
}

// ParserGetMappedUsersRequest ...
type ParserGetMappedUsersRequest struct {
}

// ParserGetMappedUsersResponse ...
type ParserGetMappedUsersResponse struct {
	// Users ...
	Users []*tdproto.MappedUser `json:"users"`
}

// ParserMapUsersRequest ...
type ParserMapUsersRequest struct {
	// Users ...
	Users []*tdproto.MappedUser `json:"users"`
}

// ParserMapUsersResponse ...
type ParserMapUsersResponse struct {
	// Success result
	Success bool `json:"success"`
}

// ParserGenerateChatsRequest
type ParserGenerateChatsRequest struct {
}

// ParserGenerateChatsResponse
type ParserGenerateChatsResponse struct {
	// ProcessingAction action for background notifications about generation of chats and messages.
	ProcessingAction string `json:"processing_action"`

	// ActionType must be generate_chat
	ActionType tdproto.ActionType `json:"action_type"`
}

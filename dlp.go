package tdproto

import "time"

// DLPBasicData structure for basic data (used for team, group, task, etc.)
type DLPBasicData struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Number      uint   `json:"number"`
}

// DLPUserData structure to store sender/receiver user information
type DLPUserData struct {
	UUID         string `json:"uuid"`
	ContactUUID  string `json:"contact_uuid"`
	Name         string `json:"name"`
	FirstName    string `json:"first_name"`
	MiddleName   string `json:"middle_name"`
	LastName     string `json:"last_name"`
	Role         string `json:"role"`
	ContactPhone string `json:"contact_phone"`
	ContactEmail string `json:"contact_email"`
	IsBot        bool   `json:"is_bot"`
}

// DLPMessageData structure to store information about message
type DLPMessageData struct {
	Text    string `json:"text"`
	Comment string `json:"comment"`
}

// DLPFileData structure to store information about file event
type DLPFileData struct {
	Link string `json:"link"` // TODO
	// Content []byte `json:"content"` // TODO
	Comment string `json:"comment"` // TODO
}

// DLPEvent structure to store all information about event
type DLPEvent struct {
	EventTime    time.Time      `json:"event_time"`
	GroupData    DLPBasicData   `json:"group_data"`
	TeamData     DLPBasicData   `json:"team_data"`
	TaskData     DLPBasicData   `json:"task_data"`
	FileData     DLPFileData    `json:"file_data"`
	MessageData  DLPMessageData `json:"message_data"`
	UUID         string         `json:"uuid"`
	EventChannel ChatType       `json:"event_channel"`
	MediaType    Mediatype      `json:"media_type"`
	SenderData   DLPUserData    `json:"sender_data"`
	DirectData   DLPUserData    `json:"direct_data"`
}

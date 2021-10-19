package tdproto

import "time"

// DLPBasicData structure for basic data (used for team, group, task, etc.)
type DLPBasicData struct {
	UUID        string
	Number      uint64
	Name        string
	Description string
}

// DLPUserData structure to store sender/receiver user information
type DLPUserData struct {
	UUID         string
	ContactUUID  string
	Name         string
	FirstName    string
	MiddleName   string
	LastName     string
	Role         string
	ContactPhone string
	ContactEmail string
	IsBot        bool
}

// DLPMessageData structure to store information about message
type DLPMessageData struct {
	Text    string
	Comment string
}

// DLPFileData structure to store information about file event
type DLPFileData struct {
	Link    string // TODO
	Content []byte // TODO
	Comment string // TODO
}

// DLPEvent structure to store all information about event
type DLPEvent struct {
	UUID         string
	EventChannel ChatType
	EventTime    time.Time
	MediaType    Mediatype
	TeamData     DLPBasicData
	SenderData   DLPUserData
	DirectData   DLPUserData
	GroupData    DLPBasicData
	TaskData     DLPBasicData
	FileData     DLPFileData
	MessageData  DLPMessageData
}

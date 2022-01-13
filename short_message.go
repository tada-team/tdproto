package tdproto

// Short message based on chat message
type ShortMessage struct {
	// Sender contact id
	From JID `json:"from" tdproto:"readonly"`

	// Recipient id (group, task or contact)
	To JID `json:"to"`

	// Message uid
	MessageId string `json:"message_id"`

	// Message creation datetime (set by server side) or sending datetime in future for draft messages
	Created ISODateTimeString `json:"created" tdproto:"readonly"`

	// Chat type
	ChatType ChatType `json:"chat_type" tdproto:"readonly"`

	// Chat id
	Chat JID `json:"chat" tdproto:"readonly"`

	// Object version
	Gentime int64 `json:"gentime" tdproto:"readonly"`

	// This message is archive
	IsArchive bool `json:"is_archive,omitempty" tdproto:"readonly"`
}

// Upload + ShortMessage
type UploadShortMessage struct {
	// Short message information
	Message ShortMessage `json:"message"`

	// Upload information
	Upload Upload `json:"upload"`
}

package tdproto

// MessengerType type of messenger. Ex. Telegram, WhatsApp etc.
type MessengerType string

const (
	MessengerTypeTelegram MessengerType = "telegram"
)

// ParseStatus parse status
type ParseStatus string

const (
	ParseStatusCreated        ParseStatus = "created"         // archive just start for uploading
	ParseStatusUploaded                   = "uploaded"        // archive fully uploaded
	ParseStatusInProgress                 = "in_progress"     // parse in progress
	ParseStatusUploadingMedia             = "uploading_media" // uploading media in progress
	ParseStatusCompleted                  = "completed"       // parse and uploading media completed
)

func (ct ParseStatus) String() string { return string(ct) }

// MappedUser struct for map internal user with external user
type MappedUser struct {
	// Contact short tada contact data
	ContactID ContactShort `json:"contact"`

	// ExternalUserID user id from messenger
	ExternalUserID string `json:"external_user_id"`

	// ExternalUserName user name from messenger
	ExternalUserName string `json:"external_user_name"`

	// IsDeleted flag of deleted user from messenger
	IsDeleted bool `json:"is_deleted"`

	// IsArchive flag of archive user
	IsArchive bool `json:"is_archive"`

	// IsAdmin group admin flag
	IsAdmin bool `json:"is_admin"`
}

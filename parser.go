package tdproto

// MessengerType type of messenger. Ex. Telegram, WhatsApp etc.
type MessengerType string

const (
	MessengerTypeTelegram MessengerType = "telegram"
)

// ParseState parse status
type ParseState string

const (
	ParseStateNotFound    ParseState = "not_found"
	ParseStateUnpacking   ParseState = "unpacking"
	ParseStateNeedMapping ParseState = "need_mapping"
	ParseStateGenerating  ParseState = "generating"
)

func (ps ParseState) String() string { return string(ps) }

type ParseErrCode uint64

// 1001: ZIP - не zip
// 1002: result.json - отсутствует
// 1003: result.json - неверный формат json
// 1004: result.json - неправильный формат(отсутствие, например, create_time)
// 1005: отсутсвует нужный upload
// 1006: сервер не сервер
const (
	ParseErrCodeInvalidZip            ParseErrCode = 1001
	ParseErrCodeResultJsonNotFound    ParseErrCode = 1002
	ParseErrCodeResultJsonIsCorrupted ParseErrCode = 1003
	ParseErrCodeInvalidJsonFormat     ParseErrCode = 1004
	ParseErrCodeUploadNotFound        ParseErrCode = 1005
	ParseErrCodeServerError           ParseErrCode = 1006
)

// MappedUser struct for map internal user with external user
type MappedUser struct {
	// Contact tada contact data
	Contact *Contact `json:"contact,omitempty"`

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

package tdproto

// Account data
type User struct {
	// Phone for login
	Phone string `json:"phone,omitempty"`

	// Email for login
	Email string `json:"email,omitempty"`

	// Family name
	FamilyName string `json:"family_name,omitempty"`

	// Given name
	GivenName string `json:"given_name,omitempty"`

	// Patronymic, if any
	Patronymic string `json:"patronymic,omitempty"`

	// Default language code
	DefaultLang string `json:"default_lang,omitempty"`

	// Use Ctrl/Cmd + Enter instead Enter
	AltSend bool `json:"alt_send"`

	// Use * as @ for mentions
	AsteriskMention bool `json:"asterisk_mention"`

	// Send pushes even user is online
	AlwaysSendPushes bool `json:"always_send_pushes"`

	// Show unread chats in chat list first
	UnreadFirst bool `json:"unread_first"`

	// Show unread chats in chat list first on mobiles
	MUnreadFirst bool `json:"munread_first"`

	// Timezone
	Timezone string `json:"timezone"`

	// Start silently time (no pushes, no sounds)
	QuietTimeStart *string `json:"quiet_time_start"`

	// Finish silently time (no pushes, no sounds)
	QuietTimeFinish *string `json:"quiet_time_finish"`
}

type UserWithMe struct {
	User
	Inviter JID          `json:"inviter,omitempty"`
	Teams   []Team       `json:"teams"`
	Devices []PushDevice `json:"devices"`
}

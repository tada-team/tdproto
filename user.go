package tdproto

// Account data
type User struct {
	// Phone for login
	Phone *string `json:"phone"`

	// Email for login
	Email *string `json:"email"`

	// Default language code
	DefaultLang *string `json:"default_lang"`

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
	Inviter *JID         `json:"inviter,omitempty"`
	Teams   []Team       `json:"teams"`
	Devices []PushDevice `json:"devices"`
}

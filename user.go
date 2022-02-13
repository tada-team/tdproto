package tdproto

// Account data
type User struct {
	// Finish silently time (no pushes, no sounds)
	QuietTimeFinish *string `json:"quiet_time_finish"`

	// Start silently time (no pushes, no sounds)
	QuietTimeStart *string `json:"quiet_time_start"`

	// Phone for login
	Phone string `json:"phone,omitempty"`

	// Given name
	GivenName string `json:"given_name,omitempty"`

	// Patronymic, if any
	Patronymic string `json:"patronymic,omitempty"`

	// Default language code
	DefaultLang string `json:"default_lang,omitempty"`

	// Timezone
	Timezone string `json:"timezone"`

	// Email for login
	Email string `json:"email,omitempty"`

	// Family name
	FamilyName string `json:"family_name,omitempty"`

	// Icon data
	Icons IconData `json:"icons"`

	// Show unread chats in chat list first
	UnreadFirst bool `json:"unread_first"`

	// Show unread chats in chat list first on mobiles
	MUnreadFirst bool `json:"munread_first"`

	// Hide pushes body
	HidePushesContent bool `json:"hide_pushes_content"`

	// Use * as @ for mentions
	AsteriskMention bool `json:"asterisk_mention"`

	// Use Ctrl/Cmd + Enter instead Enter
	AltSend bool `json:"alt_send"`

	// Send pushes even user is online
	AlwaysSendPushes bool `json:"always_send_pushes"`
}

// Account data with extra information
type UserWithMe struct {
	User

	// Inviter id, if any
	Inviter JID `json:"inviter,omitempty"`

	// Available teams
	Teams []Team `json:"teams"`

	// Registered push devices
	Devices []PushDevice `json:"devices"`

	// Personal account from billing
	Account *PersonalAccountBilling `json:"account,omitempty"`
}

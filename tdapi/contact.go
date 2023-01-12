package tdapi

import "github.com/tada-team/tdproto"

type ContactFilter struct {
	Paginator

	// Comma separated statuses in team
	Status string `schema:"status"`

	//* ?id=jid,jid
	ID string `schema:"id"`

	//* ?display_name=
	DisplayName string `schema:"display_name"`

	//* ?is_archived=
	IsArchived *bool `schema:"is_archived"`
}

type ContactsSectionFilter struct {
	Paginator

	// Team Section
	SectionUid string `schema:"section_uid"`

	//* ?is_archived=
	IsArchived *bool `schema:"is_archived"`
}

// Contact invite/edit form
type Contact struct {
	// User uid (for invitation from other team)
	UserUid string `json:"user_uid,omitempty"`

	// Phone number (for invitation by phone)
	Phone string `json:"phone"`

	// Contact phone number (part of contact information in team, visible by all)
	ContactPhone string `json:"contact_phone"`

	// Contact email (part of contact information in team, visible by all)
	ContactEmail string `json:"contact_email"`

	// Family name
	FamilyName string `json:"family_name"`

	// Given name
	GivenName string `json:"given_name"`

	// Patronymic
	Patronymic string `json:"patronymic"`

	// Role in team
	Role string `json:"role"`

	// Mood in team
	Mood string `json:"mood"`

	// Enable debug messages in UI
	DebugShowActivity bool `json:"debug_show_activity"`

	// Status in team
	Status tdproto.TeamStatus `json:"status"`

	// ContactSection uids, if any
	Sections []string `json:"sections"`

	// Use Ctrl/Cmd + Enter instead Enter
	AltSend bool `json:"alt_send"`

	// Focus mode enabled until
	FocusUntil tdproto.ISODateTimeString `json:"focus_until"`

	// Use * as @ for mentions
	AsteriskMention bool `json:"asterisk_mention"`

	// Send push notifications even contact is online
	AlwaysSendPushes bool `json:"always_send_pushes"`

	// Hide pushes body
	HidePushesContent bool `json:"hide_pushes_content"`

	// Show unread chats first (web/desktop)
	UnreadFirst bool `json:"unread_first"`

	// Show unread chats first (mobile app)
	MunreadFirst bool `json:"munread_first"`

	// Default language code
	DefaultLang string `json:"default_lang"`

	// Timezone, if any
	Timezone string `json:"timezone"`

	// Start silently time (no pushes, no sounds)
	QuietTimeStart string `json:"quiet_time_start"`

	// Quiet time finish
	QuietTimeFinish string `json:"quiet_time_finish"`

	// Push notifications for group chats
	GroupNotificationsEnabled bool `json:"group_notifications_enabled"`

	// Push notifications for task chats
	TaskNotificationsEnabled bool `json:"task_notifications_enabled"`

	// Push notifications for meeting chats
	MeetingNotificationsEnabled bool `json:"meeting_notifications_enabled"`

	// Short view in contact list
	ContactShortView bool `json:"contact_short_view"`

	// Short view in group list
	GroupShortView bool `json:"group_short_view"`

	// Short view in task list
	TaskShortView bool `json:"task_short_view"`

	// Short view in meeting list
	MeetingShortView bool `json:"meeting_short_view"`

	// Short view in contact list in mobile app
	ContactMshortView bool `json:"contact_mshort_view"`

	// Show archived contacts in contact list
	ContactShowArchived bool `json:"contact_show_archived"`

	// Short view in group list in mobile app
	GroupMshortView bool `json:"group_mshort_view"`

	// Short view in task list in mobile app
	TaskMshortView bool `json:"task_mshort_view"`

	// Short view in meeting list in mobile app
	MeetingMshortView bool `json:"meeting_mshort_view"`
}

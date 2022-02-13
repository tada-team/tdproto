package tdproto

// Contact
type Contact struct {
	// Short view in contact list in mobile app
	ContactMshortView *bool `json:"contact_mshort_view,omitempty"`

	// Short view in contact list
	ContactShortView *bool `json:"contact_short_view,omitempty"`

	// Push notifications for task chats
	TaskNotificationsEnabled *bool `json:"task_notifications_enabled,omitempty"`

	// Quiet time finish
	QuietTimeFinish *string `json:"quiet_time_finish,omitempty"`

	// Quiet time start
	QuietTimeStart *string `json:"quiet_time_start,omitempty"`

	// Short view in task list
	TaskShortView *bool `json:"task_short_view,omitempty"`

	// Timezone, if any
	Timezone *string `json:"timezone,omitempty"`

	// Hide pushes body
	HidePushesContent *bool `json:"hide_pushes_content,omitempty"`

	// Send push notifications even contact is online
	AlwaysSendPushes *bool `json:"always_send_pushes,omitempty"`

	// Push notifications for group chats
	GroupNotificationsEnabled *bool `json:"group_notifications_enabled,omitempty"`

	// Use Ctrl/Cmd + Enter instead Enter
	AltSend *bool `json:"alt_send,omitempty"`

	// Enable remove all messages experimental features
	DropallEnabled *bool `json:"dropall_enabled,omitempty"`

	// Extra contact fields
	CustomFields *ContactCustomFields `json:"custom_fields,omitempty"`

	// Enable debug messages in UI
	DebugShowActivity *bool `json:"debug_show_activity,omitempty"`

	// Short view in task list in mobile app
	TaskMshortView *bool `json:"task_mshort_view,omitempty"`

	// Show unread chats first in feed in mobile app
	MUnreadFirst *bool `json:"munread_first,omitempty"`

	// Default language code
	DefaultLang *string `json:"default_lang,omitempty"`

	// Show unread chats first in feed
	UnreadFirst *bool `json:"unread_first,omitempty"`

	// Show archived contacts in contact list
	ContactShowArchived *bool `json:"contact_show_archived,omitempty"`

	// Use * as @ for mentions
	AsteriskMention *bool `json:"asterisk_mention,omitempty"`

	// Short view in group list in mobile app
	GroupMshortView *bool `json:"group_mshort_view,omitempty"`

	// Short view in group list
	GroupShortView *bool `json:"group_short_view,omitempty"`

	// Full name in chats
	DisplayName string `json:"display_name"`

	// Contact phone in this team
	ContactPhone string `json:"contact_phone"`

	// Given name
	GivenName string `json:"given_name,omitempty"`

	// Patronymic, if any
	Patronymic string `json:"patronymic,omitempty"`

	// Why I can't send message to this chat (if can't)
	CantSendMessageReason string `json:"cant_send_message_reason,omitempty"`

	// Bot name. Empty for users
	Botname string `json:"botname,omitempty"`

	// Last activity in this team (iso datetime)
	LastActivity ISODateTimeString `json:"last_activity,omitempty"`

	// Status in this team
	TeamStatus TeamStatus `json:"status"`

	// Node uid for external users
	Node string `json:"node,omitempty"`

	// Role in this team
	Role string `json:"role"`

	// Contact Id
	Jid JID `json:"jid"`

	// Two-factor authentication status
	Auth2faStatus string `json:"auth_2fa_status,omitempty"`

	// Contact email in this team
	ContactEmail string `json:"contact_email"`

	// Short name in chats
	ShortName string `json:"short_name"`

	// Focus mode enabled until
	FocusUntil ISODateTimeString `json:"focus_until,omitempty"`

	// Family name
	FamilyName string `json:"family_name,omitempty"`

	// Mood in this team
	Mood string `json:"mood,omitempty"`

	// Changeable fields
	ChangeableFields []string `json:"changeable_fields,omitempty"`

	// Section ids
	Sections []string `json:"sections"`

	// Icons data
	Icons IconData `json:"icons"`

	// Object version
	Gentime int64 `json:"gentime"`

	// Can I add new members to this team
	CanAddToTeam bool `json:"can_add_to_team,omitempty"`

	// Two-factor authentication is configured and confirmed
	Auth2faEnabled bool `json:"auth_2fa_enabled,omitempty"`

	// Can I import tasks in this team
	CanImportTasks bool `json:"can_import_tasks,omitempty"`

	// Deprecated
	CanDeleteAnyMessage bool `json:"can_delete_any_message,omitempty"`

	// Can I create task for this contact
	CanCreateTask bool `json:"can_create_task,omitempty"`

	// Can I call to this contact
	CanCall bool `json:"can_call,omitempty"`

	// Can I send message to this contact
	CanSendMessage bool `json:"can_send_message,omitempty"`

	// Can I remove this contact from team
	CanDelete bool `json:"can_delete,omitempty"`

	// Can I manage contact sections in this team
	CanManageSections bool `json:"can_manage_sections,omitempty"`

	// Can I manage task projects in this team
	CanManageProjects bool `json:"can_manage_projects,omitempty"`

	// Can I add this contact to group chats
	CanAddToGroup bool `json:"can_add_to_group,omitempty"`

	// Can I manage integrations in this team
	CanManageIntegrations bool `json:"can_manage_integrations,omitempty"`

	// Can I manage color rules in this team
	CanManageColorRules bool `json:"can_manage_color_rules,omitempty"`

	// Can I create group chats in this team
	CanCreateGroup bool `json:"can_create_group,omitempty"`

	// Can I view/join public group in this team
	CanJoinPublicGroups bool `json:"can_join_public_groups,omitempty"`

	// Can I view/join public tasks in this team
	CanJoinPublicTasks bool `json:"can_join_public_tasks,omitempty"`

	// Contact deleted
	IsArchive bool `json:"is_archive,omitempty"`

	// Can I manage tags in this team
	CanManageTags bool `json:"can_manage_tags,omitempty"`
}

// Extra contact fields
type ContactCustomFields struct {
	// Company
	Company string `json:"company,omitempty"`

	// Department
	Department string `json:"department,omitempty"`

	// Title
	Title string `json:"title,omitempty"`

	// MobilePhone
	MobilePhone string `json:"mobile_phone,omitempty"`

	// Import source
	Source string `json:"source,omitempty"`

	// User UUID in Active Directory
	ADUid string `json:"ad_uid,omitempty"`
}

// Short contact representation
type ContactShort struct {
	// Contact Id
	Jid JID `json:"jid"`

	// Full name in chats
	DisplayName string `json:"display_name"`

	// Short name in chats
	ShortName string `json:"short_name"`

	// Icons data
	Icons IconData `json:"icons"`

	// Object version
	Gentime int64 `json:"gentime"`
}

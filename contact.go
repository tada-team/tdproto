package tdproto

// Contact
type Contact struct {
	// Contact Id
	Jid JID `json:"jid"`

	// Full name in chats
	DisplayName string `json:"display_name"`

	// Short name in chats
	ShortName string `json:"short_name"`

	// Contact email in this team
	ContactEmail string `json:"contact_email"`

	// Contact phone in this team
	ContactPhone string `json:"contact_phone"`

	// Icons data
	Icons *IconData `json:"icons"`

	// Role in this team
	Role string `json:"role"`

	// Mood in this team
	Mood string `json:"mood,omitempty"`

	// Status in this team
	TeamStatus TeamStatus `json:"status"`

	// Last activity in this team (iso datetime)
	LastActivity *string `json:"last_activity"`

	// Can contact add users to this team
	AddToTeamRights bool `json:"add_to_team_rights,omitempty"`

	// Contact deleted
	IsArchive bool `json:"is_archive,omitempty"`

	// Bot name. Empty for users
	Botname string `json:"botname,omitempty"`

	// Section ids
	Sections []string `json:"sections"`

	// Can I send message to this contact
	CanSendMessage *bool `json:"can_send_message,omitempty"`

	// Why I can't send message to this chat (if can't)
	CantSendMessageReason string `json:"cant_send_message_reason,omitempty"`

	// Can I call to this contact
	CanCall *bool `json:"can_call,omitempty"`

	// Can I call create task for this contact
	CanCreateTask *bool `json:"can_create_task,omitempty"`

	// Can I add this contact to group chats
	CanAddToGroup *bool `json:"can_add_to_group,omitempty"`

	// Can I remove this contact from team
	CanDelete *bool `json:"can_delete,omitempty"`

	// Changeable fields
	ChangeableFields *[]string `json:"changeable_fields,omitempty"`

	// Family name
	FamilyName *string `json:"family_name,omitempty"`

	// Given name
	GivenName *string `json:"given_name,omitempty"`

	// Patronymic, if any
	Patronymic *string `json:"patronymic,omitempty"`

	// Default language code
	DefaultLang *string `json:"default_lang,omitempty"`

	// Enable debug messages in UI
	DebugShowActivity *bool `json:"debug_show_activity,omitempty"`

	// Enable remove all messages experimental features
	DropallEnabled *bool `json:"dropall_enabled,omitempty"`

	// Use Ctrl/Cmd + Enter insted Enter
	AltSend *bool `json:"alt_send,omitempty"`

	// Send push notifications even contact is online
	AlwaysSendPushes *bool `json:"always_send_pushes,omitempty"`

	// Timezone, if any
	Timezone *string `json:"timezone,omitempty"`

	// Quiet time start
	QuietTimeStart *string `json:"quiet_time_start,omitempty"`

	// Quiet time finish
	QuietTimeFinish *string `json:"quiet_time_finish,omitempty"`

	// Push notifications for group chats
	GroupNotificationsEnabled *bool `json:"group_notifications_enabled,omitempty"`

	// Push notifications for task chats
	TaskNotificationsEnabled *bool `json:"task_notifications_enabled,omitempty"`

	// Short view in contact list
	ContactShortView *bool `json:"contact_short_view,omitempty"`

	// Short view in group list
	GroupShortView *bool `json:"group_short_view,omitempty"`

	// Short view in task list
	TaskShortView *bool `json:"task_short_view,omitempty"`

	// Short view in contact list in mobile app
	ContactMshortView *bool `json:"contact_mshort_view,omitempty"`

	// Short view in group list in mobile app
	GroupMshortView *bool `json:"group_mshort_view,omitempty"`

	// Short view in task list in mobile app
	TaskMshortView *bool `json:"task_mshort_view,omitempty"`

	// Show archived contacts in contact list
	ContactShowArchived *bool `json:"contact_show_archived,omitempty"`

	// Show inread chats first in feed
	UnreadFirst *bool `json:"unread_first,omitempty"`

	// Show inread chats first in feed in mobile app
	MUnreadFirst *bool `json:"munread_first,omitempty"`

	// Can I add new members to this team
	CanAddToTeam *bool `json:"can_add_to_team,omitempty"`

	// Can I manage sections in this team
	CanManageSections *bool `json:"can_manage_sections,omitempty"`

	// Can I manage tags in this team
	CanManageTags *bool `json:"can_manage_tags,omitempty"`

	// Can I manage integrations in this team
	CanManageIntegrations *bool `json:"can_manage_integrations,omitempty"`

	// Can I manage color rules in this team
	CanManageColorRules *bool `json:"can_manage_color_rules,omitempty"`

	// Can I create group chats in this team
	CanCreateGroup *bool `json:"can_create_group,omitempty"`

	// Can I view/join public group in this team
	CanJoinPublicGroups *bool `json:"can_join_public_groups,omitempty"`

	// Can I view/join public tasks in this team
	CanJoinPublicTasks *bool `json:"can_join_public_tasks,omitempty"`

	// Deprecated: use CanDeleteAnyMessage in chat object
	CanDeleteAnyMessage *bool `json:"can_delete_any_message,omitempty"`

	// Extra contact fields
	CustomFields *ContactCustomFields `json:"custom_fields,omitempty"`
}

// Extra contact fields
type ContactCustomFields struct {
	Company     string `json:"company,omitempty"`
	Department  string `json:"department,omitempty"`
	Title       string `json:"title,omitempty"`
	MobilePhone string `json:"mobile_phone,omitempty"`
}

// Short contact representaion
type ContactShort struct {
	// Contact Id
	Jid JID `json:"jid"`

	// Full name in chats
	DisplayName string `json:"display_name"`

	// Short name in chats
	ShortName string `json:"short_name"`

	// Icons data
	Icons *IconData `json:"icons"`
}

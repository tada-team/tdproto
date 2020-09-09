package tdproto

type Contact struct {
	DisplayName     string     `json:"display_name"`
	ShortName       string     `json:"short_name,omitempty"`
	ContactEmail    string     `json:"contact_email"`
	ContactPhone    string     `json:"contact_phone"`
	Icons           *IconData  `json:"icons"`
	Jid             JID        `json:"jid"`
	Role            string     `json:"role"`
	Mood            string     `json:"mood,omitempty"`
	TeamStatus      TeamStatus `json:"status"`
	LastActivity    *string    `json:"last_activity"`
	AddToTeamRights bool       `json:"add_to_team_rights,omitempty"`
	IsArchive       bool       `json:"is_archive,omitempty"`
	Botname         string     `json:"botname,omitempty"`
	Sections        []string   `json:"sections"`
	// actor:
	CanSendMessage        *bool     `json:"can_send_message,omitempty"`
	CantSendMessageReason string    `json:"cant_send_message_reason,omitempty"`
	CanCall               *bool     `json:"can_call,omitempty"`
	CanCreateTask         *bool     `json:"can_create_task,omitempty"`
	CanAddToGroup         *bool     `json:"can_add_to_group,omitempty"`
	CanDelete             *bool     `json:"can_delete,omitempty"`
	ChangeableFields      *[]string `json:"changeable_fields,omitempty"`
	FamilyName            *string   `json:"family_name,omitempty"`
	GivenName             *string   `json:"given_name,omitempty"`
	Patronymic            *string   `json:"patronymic,omitempty"`
	DefaultLang           *string   `json:"default_lang,omitempty"`
	// self:
	DebugShowActivity         *bool   `json:"debug_show_activity,omitempty"`
	DropallEnabled            *bool   `json:"dropall_enabled,omitempty"`
	AltSend                   *bool   `json:"alt_send,omitempty"`
	AlwaysSendPushes          *bool   `json:"always_send_pushes,omitempty"`
	Timezone                  *string `json:"timezone,omitempty"`
	QuietTimeStart            *string `json:"quiet_time_start,omitempty"`
	QuietTimeFinish           *string `json:"quiet_time_finish,omitempty"`
	GroupNotificationsEnabled *bool   `json:"group_notifications_enabled,omitempty"`
	TaskNotificationsEnabled  *bool   `json:"task_notifications_enabled,omitempty"`
	ContactShortView          *bool   `json:"contact_short_view,omitempty"`
	GroupShortView            *bool   `json:"group_short_view,omitempty"`
	TaskShortView             *bool   `json:"task_short_view,omitempty"`
	ContactMshortView         *bool   `json:"contact_mshort_view,omitempty"`
	ContactShowArchived       *bool   `json:"contact_show_archived,omitempty"`
	GroupMshortView           *bool   `json:"group_mshort_view,omitempty"`
	TaskMshortView            *bool   `json:"task_mshort_view,omitempty"`
	UnreadFirst               *bool   `json:"unread_first,omitempty"`
	MUnreadFirst              *bool   `json:"munread_first,omitempty"`
	// admin fields:
	CanAddToTeam          *bool `json:"can_add_to_team,omitempty"`
	CanManageSections     *bool `json:"can_manage_sections,omitempty"`
	CanManageTags         *bool `json:"can_manage_tags,omitempty"`
	CanManageIntegrations *bool `json:"can_manage_integrations,omitempty"`
	CanManageColorRules   *bool `json:"can_manage_color_rules,omitempty"`
	CanCreateGroup        *bool `json:"can_create_group,omitempty"`
	CanJoinPublicGroups   *bool `json:"can_join_public_groups,omitempty"`
	CanJoinPublicTasks    *bool `json:"can_join_public_tasks,omitempty"`
	CanDeleteAnyMessage   *bool `json:"can_delete_any_message,omitempty"` // deprecated
	// custom fields:
	CustomFields *ContactCustomFields `json:"custom_fields,omitempty"`
}

type ContactCustomFields struct {
	Company     string `json:"company,omitempty"`
	Department  string `json:"department,omitempty"`
	Title       string `json:"title,omitempty"`
	MobilePhone string `json:"mobile_phone,omitempty"`
}

package tdproto

type TeamStatus string

const (
	TeamOwner  = TeamStatus("owner")
	TeamAdmin  = TeamStatus("admin")
	TeamMember = TeamStatus("member")
	TeamGuest  = TeamStatus("guest")
)

type Team struct {
	Uid                    string       `json:"uid"`
	IsArchive              bool         `json:"is_archive,omitempty"`
	Gentime                int64        `json:"gentime"`
	Name                   string       `json:"name"`
	DefaultTaskDeadline    string       `json:"default_task_deadline,omitempty"`
	MaxMessageUpdateAge    int          `json:"max_message_update_age"`
	Icons                  IconData     `json:"icons"`
	LastActive             bool         `json:"last_active"`
	ChangeableStatuses     []TeamStatus `json:"changeable_statuses,omitempty"`
	BadProfile             bool         `json:"bad_profile,omitempty"`
	NeedConfirmation       bool         `json:"need_confirmation"`
	UsePatronymic          bool         `json:"use_patronymic,omitempty"`
	UseTaskImportance      bool         `json:"use_task_importance,omitempty"`
	TaskImportanceMin      int          `json:"task_importance_min,omitempty"`
	TaskImportanceMax      int          `json:"task_importance_max,omitempty"`
	TaskImportanceRev      bool         `json:"task_importance_rev,omitempty"`
	UseTaskUrgency         bool         `json:"use_task_urgency,omitempty"`
	UseTaskComplexity      bool         `json:"use_task_complexity,omitempty"`
	UseTaskSpentTime       bool         `json:"use_task_spent_time,omitempty"`
	DisplayFamilyNameFirst bool         `json:"display_family_name_first,omitempty"`
	UserFields             []string     `json:"user_fields"`
	UploadsSize            int64        `json:"uploads_size,omitempty"`
	UploadsSizeLimit       int64        `json:"uploads_size_limit,omitempty"`
	Unreads                *TeamUnread  `json:"unread"`
	Me                     Contact      `json:"me,omitempty"`
	Contacts               []Contact    `json:"contacts,omitempty"`
	SingleGroup            *JID         `json:"single_group,omitempty"`
	Theme                  *Theme       `json:"theme,omitempty"`
	HideArchivedUsers      bool         `json:"hide_archived_users,omitempty"`
}

// Short team representation. For invites, push notifications, etc
type TeamShort struct {
	// Team id
	Uid string `json:"uid"`

	// Team name
	Name string `json:"name"`

	// Team icon data
	Icons IconData `json:"icons"`
}

type DeletedTeam struct {
	Uid       string `json:"uid"`
	IsArchive bool   `json:"is_archive"`
	Gentime   int64  `json:"gentime"`
}

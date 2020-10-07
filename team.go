package tdproto

type TeamStatus string

const (
	TeamOwner  = TeamStatus("owner")
	TeamAdmin  = TeamStatus("admin")
	TeamMember = TeamStatus("member")
	TeamGuest  = TeamStatus("guest")
)

// Team
type Team struct {
	// Team id
	Uid string `json:"uid"`

	// Team deleted
	IsArchive bool `json:"is_archive,omitempty"`

	// Object version
	Gentime int64 `json:"gentime"`

	// Team name
	Name string `json:"name"`

	// Default task deadline
	DefaultTaskDeadline string `json:"default_task_deadline,omitempty"`

	// Max message update/deletion age, in seconds
	MaxMessageUpdateAge int `json:"max_message_update_age"`

	// Team icons
	Icons IconData `json:"icons"`

	// User last activity was in this team
	LastActive bool `json:"last_active"`

	// What status I can set to other team mebers
	ChangeableStatuses []TeamStatus `json:"changeable_statuses,omitempty"`

	// My profile in this team isn't full
	BadProfile bool `json:"bad_profile,omitempty"`

	// Neet confirmation after invite to this team
	NeedConfirmation bool `json:"need_confirmation"`

	// Patronymic in usernames for this team
	UsePatronymic bool `json:"use_patronymic,omitempty"`

	// Username fields ordering
	UserFields []string `json:"user_fields"`

	// Family name should be first in display name
	DisplayFamilyNameFirst bool `json:"display_family_name_first,omitempty"`

	// Use importance field in task
	UseTaskImportance bool `json:"use_task_importance,omitempty"`

	// Minimal value of task imporance. Default is 1
	TaskImportanceMin int `json:"task_importance_min,omitempty"`

	// Maximum value of task imporance. Default is 5
	TaskImportanceMax int `json:"task_importance_max,omitempty"`

	// Bigger number = bigger importance. Default: lower number = bigger importance
	TaskImportanceRev bool `json:"task_importance_rev,omitempty"`

	// Use urgency field in task
	UseTaskUrgency bool `json:"use_task_urgency,omitempty"`

	// Use complexity field in task
	UseTaskComplexity bool `json:"use_task_complexity,omitempty"`

	// Use spent time field in task
	UseTaskSpentTime bool `json:"use_task_spent_time,omitempty"`

	// Total uploads size, bytes
	UploadsSize int64 `json:"uploads_size,omitempty"`

	// Maximum uploads size, bytes, if any
	UploadsSizeLimit int64 `json:"uploads_size_limit,omitempty"`

	// Unread message counters
	Unreads *TeamUnread `json:"unread"`

	// My profile in this team
	Me Contact `json:"me"`

	// Team contacts. Used only for team creation
	Contacts []Contact `json:"contacts,omitempty"`

	// For single group teams, jid of chat
	SingleGroup *JID `json:"single_group,omitempty"`

	// Color theme, if any
	Theme *Theme `json:"theme,omitempty"`

	// Don't show archived users by default
	HideArchivedUsers bool `json:"hide_archived_users,omitempty"`
}

// Short team representation. For invites, push notifications, etc
type TeamShort struct {
	// Team id
	Uid string `json:"uid"`

	// Team name
	Name string `json:"name"`

	// Team icons
	Icons IconData `json:"icons"`
}

// Team deletion message
type DeletedTeam struct {
	// Team id
	Uid string `json:"uid"`

	// Team deleted
	IsArchive bool `json:"is_archive"`

	// Object version
	Gentime int64 `json:"gentime"`
}

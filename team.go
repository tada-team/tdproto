package tdproto

// Team
type Team struct {
	// Team id
	Uid string `json:"uid" tdproto:"readonly"`

	// Team deleted
	IsArchive bool `json:"is_archive,omitempty" tdproto:"readonly"`

	// Object version
	Gentime int64 `json:"gentime" tdproto:"readonly"`

	// Team name
	Name string `json:"name"`

	// Default task deadline
	DefaultTaskDeadline string `json:"default_task_deadline,omitempty"`

	// Max message update/deletion age, in seconds
	MaxMessageUpdateAge int `json:"max_message_update_age"`

	// Team icons
	Icons IconData `json:"icons" tdproto:"readonly"`

	// User last activity was in this team
	LastActive bool `json:"last_active" tdproto:"readonly"`

	// What status I can set to other team members
	ChangeableStatuses []TeamStatus `json:"changeable_statuses,omitempty" tdproto:"readonly"`

	// My profile in this team isn't full
	BadProfile bool `json:"bad_profile,omitempty" tdproto:"readonly"`

	// Need confirmation after invite to this team
	NeedConfirmation bool `json:"need_confirmation" tdproto:"readonly"`

	// Patronymic in usernames for this team
	UsePatronymic bool `json:"use_patronymic,omitempty"`

	// Username fields ordering. Possible values: "family_name", "given_name", "patronymic"
	UserFields []string `json:"user_fields" tdproto:"readonly"`

	// Family name should be first in display name
	DisplayFamilyNameFirst bool `json:"display_family_name_first,omitempty"`

	// Use importance field in task
	UseTaskImportance bool `json:"use_task_importance,omitempty"`

	// Minimal value of task importance. Default is 1
	TaskImportanceMin int `json:"task_importance_min,omitempty"`

	// Maximum value of task importance. Default is 5
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
	UploadsSize int64 `json:"uploads_size,omitempty" tdproto:"readonly"`

	// Maximum uploads size, bytes, if any
	UploadsSizeLimit int64 `json:"uploads_size_limit,omitempty" tdproto:"readonly"`

	// Unread message counters
	Unreads *TeamUnread `json:"unread"`

	// My profile in this team
	Me Contact `json:"me"`

	// Team contacts. Used only for team creation
	Contacts []Contact `json:"contacts,omitempty" tdproto:"readonly"`

	// For single group teams, jid of chat
	SingleGroup JID `json:"single_group,omitempty" tdproto:"readonly"`

	// Color theme, if any
	Theme *Theme `json:"theme,omitempty" tdproto:"readonly"`

	// Don't show archived users by default
	HideArchivedUsers bool `json:"hide_archived_users,omitempty"`

	// Team pinned
	Pinned bool `json:"pinned,omitempty"`

	// Sort ordering for pinned team
	PinnedSortOrdering int `json:"pinned_sort_ordering,omitempty"`

	// Team's available tariff by includig archive ones
	AvailableTariffs []string `json:"available_tariffs,omitempty"`

	// Сurrent team subscription
	Subscription Subscription `json:"subscription,omitempty"`

	// Сurrent team favorites bot, jid of chat
	Notebot JID `json:"notebot,omitempty" tdproto:"readonly"`

	// Personal account status for current team
	AccountStatus PersonalAccountStatus `json:"account_status" tdproto:"readonly"`

	// Team owner
	Owner Contact `json:"owner,omitempty"`
}

// Short team representation. For invites, push notifications, etc. Readonly.
type TeamShort struct {
	// Team id
	Uid string `json:"uid"`

	// Team name
	Name string `json:"name"`

	// Team icons
	Icons IconData `json:"icons"`
}

// Team deletion message. Readonly.
type DeletedTeam struct {
	// Team id
	Uid string `json:"uid"`

	// Team deleted
	IsArchive bool `json:"is_archive"`

	// Object version
	Gentime int64 `json:"gentime"`
}

type GetTeamsFilter struct {
	Status TeamStatus `json:"status,omitempty"`
}

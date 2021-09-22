package tdapi

import "github.com/tada-team/tdproto"

type Team struct {
	// Team name
	Name string `json:"name"`

	// Default task deadline
	DefaultTaskDeadline string `json:"default_task_deadline"`

	// Max message update/deletion age, in seconds
	MaxMessageUpdateAge interface{} `json:"max_message_update_age"`

	// Patronymic in usernames for this team
	UsePatronymic bool `json:"use_patronymic"`

	// Family name should be first in display name
	DisplayFamilyNameFirst bool `json:"display_family_name_first"`

	// Use importance field in task
	UseTaskImportance bool `json:"use_task_importance"`

	// Reverse importance
	TaskImportanceRev bool `json:"task_importance_rev"`

	// Use urgency field in task
	UseTaskUrgency bool `json:"use_task_urgency"`

	// Use complexity field in task
	UseTaskComplexity bool `json:"use_task_complexity"`

	// Use spent time field in task
	UseTaskSpentTime bool `json:"use_task_spent_time"`

	// Is singlegroup (for creation only)
	SingleGroup bool `json:"single_group"`

	// Team pinned
	Pinned bool `json:"pinned,omitempty"`

	// Contacts (for creation only)
	Contacts []Contact `json:"contacts"`
}

// Invitation information
type TeamPrejoinInfo struct {
	// Short team information
	Team tdproto.TeamShort `json:"team"`
}

// Team confirmation
type TeamConfirmation struct {
	// Need confirmation after invite to this team
	NeedConfirmation bool `json:"need_confirmation" tdproto:"readonly"`
}

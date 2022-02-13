package tdproto

// Set of rules to apply to tasks for coloring.
type ColorRule struct {
	// Tags filter enabled
	TagsEnabled *bool `json:"tags_enabled,omitempty"`

	// Task urgency filter enabled
	TaskUrgencyEnabled *bool `json:"task_urgency_enabled,omitempty"`

	// Task importance if task importance filter enabled
	TaskImportance *int `json:"task_importance,omitempty"`

	// Task importance filter enabled
	TaskImportanceEnabled *bool `json:"task_importance_enabled,omitempty"`

	// Project filter enabled
	ProjectEnabled *bool `json:"section_enabled,omitempty"` // TODO: rename to "project_enabled"

	// Task urgency if task urgency filter enabled
	TaskUrgency *int `json:"task_urgency,omitempty"`

	// Project id if project filter enabled
	Project string `json:"section,omitempty"` // TODO: rename to "project"

	// Task status
	TaskStatus string `json:"task_status,omitempty"`

	// Rule priority
	Description string `json:"description,omitempty"`

	// Rule id
	Uid string `json:"uid"`

	// Tag ids if tags filter enabled
	Tags []string `json:"tags,omitempty"`

	// Rule priority
	Priority int `json:"priority"`

	// Color index
	ColorIndex uint16 `json:"color_index"`
}

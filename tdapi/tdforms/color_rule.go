package tdforms

// Task color rule form
type ColorRule struct {
	// Rule priority
	Priority *int `json:"priority,omitempty"`

	// Color index
	ColorIndex *uint16 `json:"color_index,omitempty"`

	// Rule description
	Description string `json:"description"`

	// Task status
	TaskStatus *string `json:"task_status,omitempty"`

	// Project filter enabled
	ProjectEnabled bool `json:"section_enabled,omitempty"` // TODO: rename to "project_enabled"

	// Project id if project filter enabled
	Project *string `json:"section,omitempty"` // TODO: rename to "project"

	// Task importance filter enabled
	TaskImportanceEnabled bool `json:"task_importance_enabled,omitempty"`

	// Task importance if task importance filter enabled
	TaskImportance *int `json:"task_importance,omitempty"`

	// Task urgency filter enabled
	TaskUrgencyEnabled bool `json:"task_urgency_enabled,omitempty"`

	// Task urgency if task urgency filter enabled
	TaskUrgency *int `json:"task_urgency,omitempty"`

	// Tags filter enabled
	TagsEnabled bool `json:"tags_enabled,omitempty"`

	// Tag ids if tags filter enabled
	Tags *[]string `json:"tags,omitempty"`
}

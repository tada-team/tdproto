package tdproto

// Set of rules to apply to tasks for coloring.
type ColorRule struct {
	// Rule id
	Uid string `json:"uid"`

	// Rule priority
	Priority int `json:"priority"`

	// Rule description
	Description string `json:"description,omitempty"`

	// Color index
	ColorIndex uint16 `json:"color_index"`

	// Project filter enabled
	SectionEnabled *bool `json:"section_enabled,omitempty"`

	// Project id if project filter enabled
	Section string `json:"section,omitempty"`

	// Tags filter enabled
	TagsEnabled *bool `json:"tags_enabled,omitempty"`

	// Tag ids if tags filter enabled
	Tags []string `json:"tags,omitempty"`

	// Task status
	TaskStatus string `json:"task_status,omitempty"`

	// Task importance filter enabled
	TaskImportanceEnabled *bool `json:"task_importance_enabled,omitempty"`

	// Task importance if task importance filter enabled
	TaskImportance *int `json:"task_importance,omitempty"`

	// Task urgency filter enabled
	TaskUrgencyEnabled *bool `json:"task_urgency_enabled,omitempty"`

	// Task urgency if task urgency filter enabled
	TaskUrgency *int `json:"task_urgency,omitempty"`
}

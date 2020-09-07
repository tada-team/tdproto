package tdproto

type ColorRule struct {
	Uid            string   `json:"uid"`
	Priority       int      `json:"priority"`
	ColorIndex     uint16   `json:"color_index"`
	Section        string   `json:"section,omitempty"`
	Tags           []string `json:"tags,omitempty"`
	Description    string   `json:"description,omitempty"`
	TaskStatus     string   `json:"task_status,omitempty"`
	TaskImportance *int     `json:"task_importance,omitempty"`
	TaskUrgency    *int     `json:"task_urgency,omitempty"`
}

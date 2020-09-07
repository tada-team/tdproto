package tdproto

type Roster struct {
	Badge           *uint        `json:"badge"`            // TODO: ,omitempty
	ContactSections []Section    `json:"contact_sections"` // TODO: ,omitempty
	Contacts        []Contact    `json:"contacts"`         // TODO: ,omitempty
	Devices         []PushDevice `json:"devices"`          // TODO: ,omitempty
	Groups          []Chat       `json:"groups"`           // TODO: ,omitempty
	DirectChats     []Chat       `json:"direct_chats"`     // TODO: ,omitempty
	GroupSections   []Section    `json:"group_sections"`   // TODO: ,omitempty
	TaskChats       []Chat       `json:"task_chats"`       // TODO: ,omitempty
	TaskSections    []Section    `json:"task_sections"`    // TODO: ,omitempty
	TaskTabs        []TaskTab    `json:"task_tabs"`        // TODO: ,omitempty
}

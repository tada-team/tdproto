package tdproto

import (
	"log"
)

// Chat type
type ChatType string

const (
	// Direct chat
	DirectChatType ChatType = "direct"

	// Group chat
	GroupChatType ChatType = "group"

	// Task
	TaskChatType ChatType = "task"

	//Meeting
	MeetingChatType ChatType = "meeting"
)

func (ct ChatType) JidPrefix() string {
	switch ct {
	case DirectChatType:
		return ContactPrefix
	case GroupChatType:
		return GroupPrefix
	case TaskChatType:
		return TaskPrefix
	case MeetingChatType:
		return MeetingPrefix
	default:
		log.Panicf("JidPrefix(): incorrect chat type: %s", ct)
		return ""
	}
}

func (ct ChatType) String() string { return string(ct) }

func (ct ChatType) IsDirect() bool  { return ct == DirectChatType }
func (ct ChatType) IsGroup() bool   { return ct == GroupChatType }
func (ct ChatType) IsTask() bool    { return ct == TaskChatType }
func (ct ChatType) IsMeeting() bool { return ct == MeetingChatType }

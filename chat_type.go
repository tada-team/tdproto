package tdproto

import (
	"log"
)

type ChatType string

const (
	DirectChatType = ChatType("direct")
	GroupChatType  = ChatType("group")
	TaskChatType   = ChatType("task")
)

func (ct ChatType) JidPrefix() string {
	switch ct {
	case DirectChatType:
		return ContactPrefix
	case GroupChatType:
		return GroupPrefix
	case TaskChatType:
		return TaskPrefix
	default:
		log.Panicf("JidPrefix(): incorrect chat type: %s", ct)
		return ""
	}
}

func (ct ChatType) String() string { return string(ct) }

func (ct ChatType) IsDirect() bool { return ct == DirectChatType }
func (ct ChatType) IsGroup() bool  { return ct == GroupChatType }
func (ct ChatType) IsTask() bool   { return ct == TaskChatType }

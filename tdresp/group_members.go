package tdresp

import "github.com/tada-team/tdproto"

// Group members
type GroupMembers struct {
	// Group members lust
	Members []tdproto.GroupMembership `json:"members"`
}

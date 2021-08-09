package tdforms

import "github.com/tada-team/tdproto"

// Group form
type Group struct {
	// Group title
	DisplayName string `json:"display_name"`

	// Group description, optional
	Description string `json:"description"`

	// Readonly for non-admins group chat (Like Channels in Telegram bug switchable)
	ReadonlyForMembers bool `json:"readonly_for_members"`

	// Is task or group public for non-guests
	Public bool `json:"public"`

	// Any new team member will be added to this group chat
	DefaultForAll bool `json:"default_for_all"`

	// Delete messages in this chat in seconds. Experimental function
	AutocleanupAge *int `json:"autocleanup_age"`

	// Pinned message for this group
	PinnedMessage string `json:"pinned_message"`

	// Group members list
	Members []tdproto.GroupMembership `json:"members"`

	// Deprecated
	Section string `json:"section"`
}

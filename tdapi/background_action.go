package tdapi

import "github.com/tada-team/tdproto"

// Background action (for example: Excel import) information
type BackgroundAction struct {
	// Action uid
	Action string `json:"action"`

	// ActionType ...
	ActionType tdproto.ActionType `json:"action_type,omitempty"`
}

package tdapi

import "github.com/tada-team/tdproto"

// Draft form and response
type Draft struct {
	// Draft content
	Draft string `json:"draft"`

	// Deprecated: use Revision instead
	DraftGentime int64 `json:"draft_gentime,omitempty"  tdproto:"readonly"`

	// Unixtime(ms)
	Revision int64 `json:"revision,omitempty"`

	// TdMarkup found in content (links included)
	Markup []tdproto.MarkupEntity `json:"markup,omitempty"  tdproto:"readonly"`

	// Deprecated: use Markup instead
	Links tdproto.MessageLinks `json:"links,omitempty"  tdproto:"readonly"`
}

package tdapi

import "github.com/tada-team/tdproto"

// Draft form and response
type Draft struct {
	// Draft content
	Draft string `json:"draft"`

	// Draft version
	DraftGentime int64 `json:"draft_gentime,omitempty"  tdproto:"readonly"`

	// Links found in content
	Links tdproto.MessageLinks `json:"links,omitempty"  tdproto:"readonly"`

	// TdMorkup found in content (links included)
	Markup []tdproto.MarkupEntity `json:"markup,omitempty"  tdproto:"readonly"`
}

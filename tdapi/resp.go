package tdapi

import "github.com/tada-team/tdproto"

// Server responce
type Resp struct {
	DebugTime string            `json:"_time,omitempty"`
	Ok        bool              `json:"ok"`
	Result    interface{}       `json:"result,omitempty"`
	Error     Err               `json:"error,omitempty"`
	Details   map[string]string `json:"details,omitempty"`
	// Reason answers why not ok or has error
	Reason string `json:"reason,omitempty"`
	// Entities for reason. Experimental
	Markup []tdproto.MarkupEntity `json:"markup,omitempty"`
}

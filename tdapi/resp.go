package tdapi

import "github.com/tada-team/tdproto"

// Server response
type Resp struct {
	// Request status
	Ok bool `json:"ok"`

	// Result only if ok is true)
	Result interface{} `json:"result,omitempty"`

	// Error (only if ok is false)
	Error Err `json:"error,omitempty"`

	// Error (only if ok is false and Error is 'InvalidData')
	Details map[string]string `json:"details,omitempty"`

	// Reason (only if ok is false and Error is `AccessDenied`)
	Reason string `json:"reason,omitempty"`

	// Reason markup (only if ok is false and Error is `AccessDenied`)
	Markup []tdproto.MarkupEntity `json:"markup,omitempty"`

	// Server side work time
	DebugTime string `json:"_time,omitempty"`
}

package tdforms

import "github.com/tada-team/tdproto"

// Integration
type Integration struct {
	// Enabled
	Enabled bool `json:"enabled"`

	// Optional commend
	Comment string `json:"comment"`

	// Plugin name
	Kind string `json:"kind"`

	// Group chat jid
	GroupJid tdproto.JID `json:"group"`

	// WebhookUrl, if any
	WebhookUrl string `json:"webhook_url"`

	// Integration form
	IntegrationForm *tdproto.IntegrationForm `json:"form"`
}

package tdapi

import "github.com/tada-team/tdproto"

type Integration struct {
	// Integration form
	IntegrationForm *tdproto.IntegrationForm `json:"form"`

	// Optional commend
	Comment string `json:"comment"`

	// Plugin name
	Kind string `json:"kind"`

	// Group chat jid
	GroupJid tdproto.JID `json:"group"`

	// WebhookUrl, if any
	WebhookUrl string `json:"webhook_url"`

	// Enabled
	Enabled bool `json:"enabled"`
}

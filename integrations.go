package tdproto

// Integration form field
type IntegrationField struct {
	// Label
	Label string `json:"label"`

	// Is field readonly
	Readonly bool `json:"readonly"`

	// Current value
	Value string `json:"value"`
}

// Integration form
type IntegrationForm struct {
	// Api key field, if any
	ApiKey *IntegrationField `json:"api_key,omitempty"`

	// Webhook url, if any
	WebhookUrl *IntegrationField `json:"webhook_url,omitempty"`

	// Url, if any
	Url *IntegrationField `json:"url,omitempty"`
}

// Integration for concrete chat
type Integration struct {
	// Id
	Uid string `json:"uid,omitempty"`

	// Comment, if any
	Comment string `json:"comment"`

	// Creation datetime, iso
	Created ISODateTimeString `json:"created,omitempty"`

	// Integration enabled
	Enabled bool `json:"enabled"`

	// Integration form
	Form IntegrationForm `json:"form"`

	// Chat id
	Group JID `json:"group"`

	// Full description
	Help string `json:"help,omitempty"`

	// Unique integration name
	Kind string `json:"kind"`

	Title string `json:"-"`
}

// Integration kind
type IntegrationKind struct {
	// Integration unique name
	Kind string `json:"kind"`

	// Plugin title
	Title string `json:"title"`

	// Integration template
	Template Integration `json:"template"`

	// Path to icon
	Icon string `json:"icon"`

	// Plugin description
	Description string `json:"description"`
}

// Complete integrations data, as received from server
type Integrations struct {
	// Currently existing integrations
	Integrations []Integration `json:"integrations"`
	// Types of integrations available for setup
	Kinds []IntegrationKind `json:"kinds"`
}

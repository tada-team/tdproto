package tdproto

type IntegrationField struct {
	Label    string `json:"label"`
	Readonly bool   `json:"readonly"`
	Value    string `json:"value"`
}

type IntegrationForm struct {
	ApiKey     *IntegrationField `json:"api_key,omitempty"`
	WebhookUrl *IntegrationField `json:"webhook_url,omitempty"`
	Url        *IntegrationField `json:"url,omitempty"`
}

type Integration struct {
	Comment string          `json:"comment"`
	Created string          `json:"created,omitempty"`
	Enabled bool            `json:"enabled"`
	Form    IntegrationForm `json:"form"`
	Group   *JID            `json:"group,omitempty"`
	Title   string          `json:"-"`
	Help    string          `json:"help,omitempty"`
	Kind    string          `json:"kind"`
	Uid     string          `json:"uid,omitempty"`
}

type IntegrationList []Integration

type IntegrationKind struct {
	Kind     string      `json:"kind"`
	Template Integration `json:"template"`
	Title    string      `json:"title"`
}

type IntegrationKindList []IntegrationKind

type Integrations struct {
	Integrations IntegrationList     `json:"integrations"`
	Kinds        IntegrationKindList `json:"kinds"`
}

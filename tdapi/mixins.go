package tdapi

type Paginator struct {
	// Limit
	Limit int `json:"limit,omitempty"`

	// Offset
	Offset int `json:"offset,omitempty"`
}

type UserParams struct {
	// Language code
	Lang     string `schema:"lang,omitempty"`

	// Timezone
	Timezone string `schema:"timezone,omitempty"`
}

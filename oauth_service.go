package tdproto

// OAuth service
type OAuthService struct {
	// Integration title
	Name string `json:"name"`

	// Redirect url
	Url  string `json:"url"`
}

package tdproto

// Contact preview
type ContactPreview struct {
	Error      string `json:"_error,omitempty"`
	Phone      string `json:"phone"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Patronymic string `json:"patronymic,omitempty"`
	Role       string `json:"role"`
	Section    string `json:"section"`
}

package tdproto

type CreateBillingContactRequest struct {
	Phone      string     `json:"phone"`
	GivenName  string     `json:"given_name,omitempty"`
	FamilyName string     `json:"family_name,omitempty"`
	Patronymic string     `json:"patronymic,omitempty"`
	Status     TeamStatus `json:"status,omitempty"`
	Role       string     `json:"role,omitempty"`
	UserUid    string     `json:"user_uid,omitempty"`
	Sections   []string   `json:"sections,omitempty"`
}

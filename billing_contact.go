package tdproto

type CreateBillingContactRequest struct {
	Phone      string     `json:"phone"`
	GivenName  string     `json:"given_name"`
	FamilyName string     `json:"family_name"`
	Patronymic string     `json:"patronymic"`
	Status     TeamStatus `json:"status"`
	Role       string     `json:"role"`
	UserUid    string     `json:"user_uid"`
	Sections   []string   `json:"sections"`
}

package tdproto

type ResponsiblePerson struct {
	Id          string  `json:"id"`
	DisplayName string  `json:"display_name"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Patronymic  *string `json:"patronymic,omitempty"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email"`
	HeldPost    *string `json:"held_post,omitempty"`
}

type ResponsiblePersonCreateRequest struct {
	PersonalAccountId string  `json:"personal_account_id"`
	FirstName         string  `json:"first_name"`
	LastName          string  `json:"last_name"`
	Patronymic        *string `json:"patronymic,omitempty"`
	Phone             string  `json:"phone"`
	Email             string  `json:"email"`
	HeldPost          *string `json:"held_post,omitempty"`
}

type ResponsiblePersonCreateResponse struct {
	ResponsiblePerson
}

type ResponsiblePersonUpdateRequest struct {
	Id          string  `json:"id"`
	DisplayName string  `json:"display_name"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Patronymic  *string `json:"patronymic,omitempty"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email"`
	HeldPost    *string `json:"held_post,omitempty"`
}

type ResponsiblePersonUpdateResponse struct {
	ResponsiblePerson
}

type ResponsiblePersonGetRequest struct {
	ResponsiblePersonIds string `json:"responsible_person_ids"`
}

type ResponsiblePersonGetListRequest struct {
	Limit  *uint32 `json:"limit,omitempty"`
	Offset *uint64 `json:"offset,omitempty"`
}

type ResponsiblePersonGetResponse struct {
	ResponsiblePersonList []ResponsiblePerson `json:"responsible_person_list"`
}

type ResponsiblePersonDeleteRequest struct {
	Id string `json:"id"`
}

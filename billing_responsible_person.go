package tdproto

type ResponsiblePerson struct {
	Id             string `json:"id"`
	CounterpartyId string `json:"counterparty_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Patronymic     string `json:"patronymic,omitempty"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
}

type ResponsiblePersonCreateRequest struct {
	CounterpartyId string `json:"counterparty_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Patronymic     string `json:"patronymic,omitempty"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
}

type ResponsiblePersonCreateResponse struct {
	ResponsiblePerson
}

type ResponsiblePersonUpdateRequest struct {
	Id             string `json:"id"`
	CounterpartyId string `json:"counterparty_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Patronymic     string `json:"patronymic"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
}

type ResponsiblePersonUpdateResponse struct {
	ResponsiblePerson
}

type ResponsiblePersonGetRequest struct {
	ResponsiblePersonIds string `json:"responsible_person_ids"`
}

type ResponsiblePersonGetListRequest struct {
	CounterpartyId string `json:"counterparty_id,omitempty"`
	Limit          uint32 `json:"limit,omitempty"`
	Offset         uint32 `json:"offset,omitempty"`
}

type ResponsiblePersonGetResponse struct {
	ResponsiblePersonList []ResponsiblePerson `json:"responsible_person_list"`
}

type ResponsiblePersonDeleteRequest struct {
	Id string `json:"id"`
}

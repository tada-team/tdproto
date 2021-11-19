package tdproto

// AddWorkplacesOnPersonalAccountRequest request on add workplace on personal account
type AddWorkplacesOnPersonalAccountRequest struct {
	PersonalAccountId string `json:"personal_account_id,omitempty"` // TODO: must be int64
	CountWorkplaces   int32  `json:"count_workplaces,omitempty"`
}

// AddWorkplacesOnPersonalAccountResponse response on add workplace on personal account
type AddWorkplacesOnPersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}

// ActivateWorkplacesOnPersonalAccountRequest request on activate workplace on personal account
type ActivateWorkplacesOnPersonalAccountRequest struct {
	PersonalAccountId string `json:"personal_account_id,omitempty"` // TODO: must be int64
	CountWorkplaces   int32  `json:"count_workplaces,omitempty"`
}

// ActivateWorkplacesOnPersonalAccountResponse response on activate workplace on personal account
type ActivateWorkplacesOnPersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}

// DeleteWorkplacesOnPersonalAccountRequest request on delete workplace on personal account
type DeleteWorkplacesOnPersonalAccountRequest struct {
	PersonalAccountId string `json:"personal_account_id,omitempty"` // TODO: must be int64
	CountWorkplaces   int32  `json:"count_workplaces,omitempty"`
}

// DeleteWorkplacesOnPersonalAccountResponse response on delete workplace on personal account
type DeleteWorkplacesOnPersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}

// AddUserInWorkplaceRequest request on add user in workplace on personal account
type AddUserInWorkplaceRequest struct {
	PersonalAccountId string `json:"personal_account_id,omitempty"` // TODO: must be int64
	UserUuid          string `json:"user_uuid,omitempty"`
}

// AddUserInWorkplaceResponse response on add user in workplace on personal account
type AddUserInWorkplaceResponse struct {
	Success bool `json:"success,omitempty"`
}

// DeleteUserFromWorkplaceRequest request on delete user from workplace on personal account
type DeleteUserFromWorkplaceRequest struct {
	PersonalAccountId string `json:"personal_account_id,omitempty"` // TODO: must be int64
	UserUuid          string `json:"user_uuid,omitempty"`
}

// DeleteUserFromWorkplaceResponse response on delete user from workplace on personal account
type DeleteUserFromWorkplaceResponse struct {
	Success bool `json:"success,omitempty"`
}

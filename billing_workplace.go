package tdproto

// WorkplaceBilling struct of workplace on personal account
type WorkplaceBilling struct {
	PersonalAccountId int64  `json:"personal_account_id"`
	WorkplaceId       int64  `json:"workplace_id,omitempty"`
	UserId            int64  `json:"user_id,omitempty"`
	UserUuid          string `json:"user_uuid,omitempty"`
}

// WorkplaceOptions struct for pagination
type WorkplaceOptions struct {
	Limit  int32 `json:"limit,omitempty"`
	Offset int32 `json:"offset,omitempty"`
}

// GetWorkplacesByPersonalAccountRequest request on get workplaces by personal account
type GetWorkplacesByPersonalAccountRequest struct {
	PersonalAccountId int64             `json:"personal_account_id"`
	Options           *WorkplaceOptions `json:"options,omitempty"`
}

// GetWorkplacesByPersonalAccountResponse response on get workplaces by personal account
type GetWorkplacesByPersonalAccountResponse struct {
	Workplaces []WorkplaceBilling `json:"workplaces,omitempty"`
}

// GetUnpaidWorkplacesByPersonalAccountRequest request on get count unpaid workplaces by personal account
type GetUnpaidWorkplacesByPersonalAccountRequest struct {
	PersonalAccountId int64 `json:"personal_account_id"`
}

// GetUnpaidWorkplacesByPersonalAccountResponse response on get count unpaid workplaces by personal account
type GetUnpaidWorkplacesByPersonalAccountResponse struct {
	Count int32 `json:"count,omitempty"`
}

// GetWorkplaceByPersonalAccountResponse response on get workplace by personal account
type GetWorkplaceByPersonalAccountResponse struct {
	WorkplaceBilling
}

// AddWorkplacesOnPersonalAccountRequest request on add workplace on personal account
type AddWorkplacesOnPersonalAccountRequest struct {
	PersonalAccountId int64 `json:"personal_account_id,omitempty"`
	CountWorkplaces   int32 `json:"count_workplaces,omitempty"`
}

// AddWorkplacesOnPersonalAccountResponse response on add workplace on personal account
type AddWorkplacesOnPersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}

// ActivateWorkplacesOnPersonalAccountRequest request on activate workplace on personal account
type ActivateWorkplacesOnPersonalAccountRequest struct {
	PersonalAccountId int64 `json:"personal_account_id,omitempty"`
	CountWorkplaces   int32 `json:"count_workplaces,omitempty"`
}

// ActivateWorkplacesOnPersonalAccountResponse response on activate workplace on personal account
type ActivateWorkplacesOnPersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}

// DeleteWorkplacesOnPersonalAccountRequest request on delete workplace on personal account
type DeleteWorkplacesOnPersonalAccountRequest struct {
	PersonalAccountId int64 `json:"personal_account_id,omitempty"`
	CountWorkplaces   int32 `json:"count_workplaces,omitempty"`
}

// DeleteWorkplacesOnPersonalAccountResponse response on delete workplace on personal account
type DeleteWorkplacesOnPersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}

// AddUserInWorkplaceRequest request on add user in workplace on personal account
type AddUserInWorkplaceRequest struct {
	PersonalAccountId int64  `json:"personal_account_id,omitempty"`
	UserUuid          string `json:"user_uuid,omitempty"`
}

// AddUserInWorkplaceResponse response on add user in workplace on personal account
type AddUserInWorkplaceResponse struct {
	Success bool `json:"success,omitempty"`
}

// DeleteUserFromWorkplaceRequest request on delete user from workplace on personal account
type DeleteUserFromWorkplaceRequest struct {
	PersonalAccountId int64  `json:"personal_account_id,omitempty"`
	UserUuid          string `json:"user_uuid,omitempty"`
}

// DeleteUserFromWorkplaceResponse response on delete user from workplace on personal account
type DeleteUserFromWorkplaceResponse struct {
	Success bool `json:"success,omitempty"`
}

package tdproto

import (
	"time"
)

// WorkplaceBilling struct of workplace on personal account
type WorkplaceBilling struct {
	PersonalAccountId int64  `json:"personal_account_id"`
	WorkplaceId       int64  `json:"workplace_id,omitempty"`
	UserId            int64  `json:"user_id,omitempty"`
	UserUuid          string `json:"user_uuid,omitempty"`
}

// UserInfo user information
type UserInfo struct {
	Uuid         string    `json:"uuid"`
	Fio          string    `json:"fio,omitempty"`
	Phone        string    `json:"phone,omitempty"`
	Email        string    `json:"email,omitempty"`
	LastActivity time.Time `json:"last_activity,omitempty"`
}

// UsersInfo users information
type UsersInfo struct {
	UserInfo []UserInfo `json:"user_info,omitempty"`
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

// GetUsersInfoByUserUUIDArrayRequest request on get user information by array of UUID's users
type GetUsersInfoByUserUUIDArrayRequest struct {
	UserUuid []string `json:"user_uuid"`
	Limit    int32    `json:"limit,omitempty"`
	Offset   int32    `json:"offset,omitempty"`
}

// GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest request on get user information by array of UUID's users excluding team members in uuid team
type GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest struct {
	UserUuid []string `json:"user_uuid"`
	TeamUuid string   `json:"team_uuid"`
	Limit    int32    `json:"limit,omitempty"`
	Offset   int32    `json:"offset,omitempty"`
}

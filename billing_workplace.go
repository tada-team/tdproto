package tdproto

import (
	"time"
)

// WorkplaceBilling struct of workplace on personal account
type WorkplaceBilling struct {
	WorkplaceId string `json:"workplace_id,omitempty"`
	UserId      string `json:"user_id,omitempty"`
	UserUuid    string `json:"user_uuid,omitempty"`
}

// UserInfo user information
type UserInfo struct {
	Uuid         string     `json:"uuid"`
	FullName     string     `json:"full_name,omitempty"`
	Phone        string     `json:"phone,omitempty"`
	Email        string     `json:"email,omitempty"`
	LastActivity *time.Time `json:"last_activity,omitempty"`
}

// WorkplaceOptions struct for pagination
type WorkplaceOptions struct {
	Limit  int32 `json:"limit,omitempty"`
	Offset int32 `json:"offset,omitempty"`
}

// GetWorkplacesByPersonalAccountRequest request on get workplaces by personal account
type GetWorkplacesByPersonalAccountRequest struct {
	Options *WorkplaceOptions `json:"options,omitempty"`
}

// GetWorkplacesByPersonalAccountResponse response on get workplaces by personal account
type GetWorkplacesByPersonalAccountResponse struct {
	Workplaces []WorkplaceBilling `json:"workplaces,omitempty"`
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
	WorkplacesCount int32 `json:"workplaces_count"`
}

// AddWorkplacesOnPersonalAccountResponse response on add workplace on personal account
type AddWorkplacesOnPersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}

// ActivateWorkplacesOnPersonalAccountRequest request on activate workplace on personal account
type ActivateWorkplacesOnPersonalAccountRequest struct {
	CountWorkplaces int32 `json:"count_workplaces"`
}

// ActivateWorkplacesOnPersonalAccountResponse response on activate workplace on personal account
type ActivateWorkplacesOnPersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}

// DeleteWorkplacesOnPersonalAccountRequest request on delete workplace on personal account
type DeleteWorkplacesOnPersonalAccountRequest struct {
	CountWorkplaces int32 `json:"count_workplaces"`
}

// DeleteWorkplacesOnPersonalAccountResponse response on delete workplace on personal account
type DeleteWorkplacesOnPersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}

// AddUserInWorkplaceRequest request on add user in workplace on personal account
type AddUserInWorkplaceRequest struct {
	UserUuid string `json:"user_uuid"`
}

// AddUserInWorkplaceResponse response on add user in workplace on personal account
type AddUserInWorkplaceResponse struct {
	Success bool `json:"success,omitempty"`
}

// DeleteUserFromWorkplaceRequest request on delete user from workplace on personal account
type DeleteUserFromWorkplaceRequest struct {
	UserUuid string `json:"user_uuid"`
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

// GetUsersInfoByUserUUIDArrayResponse response on get user information by array of UUID's users
type GetUsersInfoByUserUUIDArrayResponse struct {
	PaginatedBillingUsersInfo
}

// GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest request on get user information by array of UUID's users excluding team members in uuid team
type GetUsersInfoByUserUUIDArrayExcludingTeamMembersRequest struct {
	UserUuid []string `json:"user_uuid"`
	TeamUuid string   `json:"team_uuid"`
	Limit    int32    `json:"limit,omitempty"`
	Offset   int32    `json:"offset,omitempty"`
}

// GetUsersInfoByUserUUIDArrayExcludingTeamMembersResponse response on get user information by array of UUID's users excluding team members in uuid team
type GetUsersInfoByUserUUIDArrayExcludingTeamMembersResponse struct {
	UserInfo []UserInfo `json:"user_info,omitempty"`
}

// AddUserInWorkplaceByJidRequest request on add user on workplace by contact JID
type AddUserInWorkplaceByJidRequest struct {
	ContactUuid string `json:"contact_uuid,omitempty"`
}

// AddUserInWorkplaceByJidResponse response on add user on workplace by contact JID
type AddUserInWorkplaceByJidResponse struct {
	UserInfo
}

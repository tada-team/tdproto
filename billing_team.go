package tdproto

import "time"

// TeamBilling struct of billing api
type TeamBilling struct {
	DeleteDate time.Time `json:"delete_date,omitempty"`
	TeamUuid   string    `json:"team_uuid,omitempty"`
}

// AddTeamOnPersonalAccountRequest request on add team on personal account
type AddTeamOnPersonalAccountRequest struct {
	TeamBilling
}

// AddTeamOnPersonalAccountResponse response on add team on personal account
type AddTeamOnPersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}

// DeleteTeamFromPersonalAccountRequest request on delete team from personal account
type DeleteTeamFromPersonalAccountRequest struct {
	TeamBilling
}

// DeleteTeamFromPersonalAccountResponse response on delete team from personal account
type DeleteTeamFromPersonalAccountResponse struct {
	Success bool `json:"success,omitempty"`
}

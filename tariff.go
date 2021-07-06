package tdproto

// Tariff for teams
type Tariff struct {
	// Cloud space reserved for storing team users uploads in megabytes
	CloudSpace int `json:"cloud_space"`

	// Maximum allowed number of members in a team
	MaxMembersInTeam int `json:"max_members_in_team"`

	// Maximum number of participants per call
	MaxParticipantsPerCall int `json:"max_participants_per_call"`
}

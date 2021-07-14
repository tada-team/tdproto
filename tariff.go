package tdproto

// Tariff for teams
type Tariff struct {
	// Key for updating tariff in team
	Key string `json:"key"`

	// Name of tariff on english
	NameEn string `json:"name_en"`

	// Name of tariff on russian
	NameRu string `json:"name_ru"`

	// Cloud space reserved for storing team users uploads in megabytes
	CloudSpace int64 `json:"cloud_space"`

	// Maximum allowed number of members in a team
	MaxMembersInTeam int `json:"max_members_in_team"`

	// Maximum number of participants per call
	MaxParticipantsPerCall int `json:"max_participants_per_call"`
}

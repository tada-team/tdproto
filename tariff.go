package tdproto

// Tariff for teams
type Tariff struct {
	// Tariff id
	Uid string `json:"uid"`

	// Title of tariff on enlish
	TitleEn string `json:"title_en"`

	// Title of tariff on russian
	TitleRu string `json:"title_ru"`

	// Cloud space reserved for storing team users uploads in megabytes
	CloudSpace int64 `json:"cloud_space,omitempty"`

	// Maximum allowed number of members in a team
	MaxMembersInTeam int `json:"max_members_in_team,omitempty"`

	// Maximum number of participants per call
	MaxParticipantsPerCall int `json:"max_participants_per_call,omitempty"`

	//maximum file size for uploading
	MaxUploadFilesize int64 `json:"max_upload_filesize,omitempty"`
}

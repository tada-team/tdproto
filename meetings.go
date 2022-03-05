package tdproto

type Meeting struct {
	Id           string            `json:"id"`
	TeamUuid     string            `json:"team_uuid"`
	ChatUuid     string            `json:"chat_uuid"`
	StartAt      ISODateTimeString `json:"start_at"`
	EndAt        ISODateTimeString `json:"end_at"`
	FreqDays     []int             `json:"freq_days,omitempty"`
	Freq         int               `json:"freq,omitempty"`
	IsArchive    bool              `json:"is_archive,omitempty"`
	IsPublic     bool              `json:"is_public,omitempty"`
	IsOutside    bool              `json:"is_outside,omitempty"`
	CanAddMember bool              `json:"can_add_member,omitempty"`
	CanDelete    bool              `json:"can_delete,omitempty"`
	CanEdit      bool              `json:"can_edit,omitempty"`
	CanJoin      bool              `json:"can_join,omitempty"`
}

type MeetingsRequestParams struct {
	Year           int      `json:"year"`
	Month          int      `json:"month"`
	Day            int      `json:"day,omitempty"`
	TeamUuid       string   `json:"team_uuid,omitempty"`
	Owners         []string `json:"owners,omitempty"`
	OwnersSections []string `json:"owners_sections,omitempty"`
	Members        []string `json:"members,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Offset         int      `json:"offset,omitempty"`
	IsArchive      bool     `json:"is_archive,omitempty"`
	IsFreq         bool     `json:"is_freq,omitempty"`
	IsPublic       bool     `json:"is_public,omitempty"`
	IsOutside      bool     `json:"is_outside,omitempty"`
}

type MeetingsResponse struct {
	Items  []Meeting `json:"items"`
	Limit  int       `json:"limit,omitempty"`
	Offset int       `json:"offset,omitempty"`
	Total  int       `json:"total,omitempty"`
}

type MeetingsCreateRequestMembers struct {
	Jid        JID                 `json:"jid"`
	Status     MeetingMemberStatus `json:"status,omitempty"`
	IsRequired bool                `json:"is_required,omitempty"`
}

type MeetingsCreateRequest struct {
	TeamUuid      string                         `json:"team_uuid,omitempty"`
	StartAt       ISODateTimeString              `json:"start_at"`
	EndAt         string                         `json:"end_at"`
	Freq          int                            `json:"freq,omitempty"`
	FreqDays      []int                          `json:"freq_days,omitempty"`
	Members       []MeetingsCreateRequestMembers `json:"members"`
	OwnerPresence MeetingMemberPresence          `json:"owner_presence"`
	IsPublic      bool                           `json:"is_public,omitempty"`
	IsOutside     bool                           `json:"is_outside,omitempty"`
}

type MeetingsUpdateRequest struct {
	ActiveFrom string `json:"active_from,omitempty"`
	StartAt    string `json:"start_at,omitempty"`
	EndAt      string `json:"end_at,omitempty"`
	Freq       int    `json:"freq,omitempty"`
	FreqDays   []int  `json:"freq_days,omitempty"`
	IsPublic   bool   `json:"is_public,omitempty"`
	IsOutside  bool   `json:"is_outside,omitempty"`
}

type MeetingsDeleteRequestParams struct {
	Date ISODateTimeString `json:"date,omitempty"`
}

type MeetingMember struct {
	Contact           Contact               `json:"contact"`
	Presence          MeetingMemberPresence `json:"presence"`
	Status            MeetingMemberStatus   `json:"status"`
	IsRequired        bool                  `json:"is_required,omitempty"`
	CanChangePresence bool                  `json:"can_change_presence,omitempty"`
	CanChangeStatus   bool                  `json:"can_change_status,omitempty"`
	CanRemove         bool                  `json:"can_remove,omitempty"`
}

type MeetingsMembersRequestParams struct {
	Sections   []string              `json:"sections,omitempty"`
	Presence   MeetingMemberPresence `json:"presence,omitempty"`
	Status     MeetingMemberStatus   `json:"status,omitempty"`
	Limit      int                   `json:"limit,omitempty"`
	Offset     int                   `json:"offset,omitempty"`
	IsRequired bool                  `json:"is_required,omitempty"`
}

type MeetingsMembersResponse struct {
	Items  []MeetingMember `json:"items"`
	Limit  int             `json:"limit,omitempty"`
	Offset int             `json:"offset,omitempty"`
	Total  int             `json:"total,omitempty"`
}

type MeetingsMembersCreateRequest struct {
	Jid        JID                 `json:"jid"`
	Status     MeetingMemberStatus `json:"status,omitempty"`
	IsRequired bool                `json:"is_required,omitempty"`
}

type MeetingsMembersUpdateRequest struct {
	Status     MeetingMemberStatus `json:"status,omitempty"`
	IsRequired bool                `json:"is_required,omitempty"`
}

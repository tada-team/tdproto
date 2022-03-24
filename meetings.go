package tdproto

type Meeting struct {
	Id                string            `json:"id"`
	TeamUuid          string            `json:"team_uuid"`
	ChatUuid          string            `json:"chat_uuid"`
	StartAt           ISODateTimeString `json:"start_at"`
	EndAt             ISODateTimeString `json:"end_at"`
	PersonalAccountId string            `json:"personal_account_id"`
	FreqDays          []int32           `json:"freq_days,omitempty"`
	Freq              int32             `json:"freq,omitempty"`
	IsArchive         bool              `json:"is_archive,omitempty"`
	IsPublic          bool              `json:"is_public,omitempty"`
	IsOutside         bool              `json:"is_outside,omitempty"`
	CanAddMember      bool              `json:"can_add_member,omitempty"`
	CanDelete         bool              `json:"can_delete,omitempty"`
	CanEdit           bool              `json:"can_edit,omitempty"`
	CanJoin           bool              `json:"can_join,omitempty"`
	IsFreq            bool              `json:"is_freq"`
}

type MeetingsRequestParams struct {
	PersonalAccountId string   `json:"personal_account_id"`
	Year              int32    `json:"year"`
	Month             int32    `json:"month"`
	Day               int32    `json:"day,omitempty"`
	TeamUuid          string   `json:"team_uuid,omitempty"`
	Members           []string `json:"members,omitempty"`
	Limit             int      `json:"limit,omitempty"`
	Offset            int      `json:"offset,omitempty"`
	IsArchive         bool     `json:"is_archive,omitempty"`
	IsFreq            bool     `json:"is_freq,omitempty"`
	IsPublic          bool     `json:"is_public,omitempty"`
	IsOutside         bool     `json:"is_outside,omitempty"`
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
	TeamUuid  string                         `json:"team_uuid,omitempty"`
	StartAt   ISODateTimeString              `json:"start_at"`
	EndAt     string                         `json:"end_at"`
	Freq      int32                          `json:"freq,omitempty"`
	FreqDays  []int32                        `json:"freq_days,omitempty"`
	Members   []MeetingsCreateRequestMembers `json:"members"`
	IsPublic  bool                           `json:"is_public,omitempty"`
	IsOutside bool                           `json:"is_outside,omitempty"`
	IsFreq    bool                           `json:"is_freq"`
}

type MeetingsUpdateRequest struct {
	MeetingId  string  `json:"meeting_id"`
	ActiveFrom string  `json:"active_from,omitempty"`
	StartAt    string  `json:"start_at,omitempty"`
	EndAt      string  `json:"end_at,omitempty"`
	TeamUuid   string  `json:"team_uuid,omitempty"`
	Freq       int32   `json:"freq,omitempty"`
	FreqDays   []int32 `json:"freq_days,omitempty"`
	IsPublic   bool    `json:"is_public,omitempty"`
	IsOutside  bool    `json:"is_outside,omitempty"`
	IsFreq     bool    `json:"is_freq"`
}

type MeetingsDeleteRequestParams struct {
	Date ISODateTimeString `json:"date,omitempty"`
}

type MeetingMember struct {
	MeetingId         string                `json:"meeting_id,omitempty"`
	Contact           Contact               `json:"contact"`
	Presence          MeetingPresenceStatus `json:"presence"`
	Status            MeetingMemberStatus   `json:"status"`
	IsRequired        bool                  `json:"is_required,omitempty"`
	CanChangePresence bool                  `json:"can_change_presence,omitempty"`
	CanChangeStatus   bool                  `json:"can_change_status,omitempty"`
	CanRemove         bool                  `json:"can_remove,omitempty"`
}

type MeetingsMembersRequestParams struct {
	UuidSections []string              `json:"uuid_sections,omitempty"`
	Presence     MeetingPresenceStatus `json:"presence,omitempty"`
	Status       MeetingMemberStatus   `json:"status,omitempty"`
	MeetingId    string                `json:"meeting_id"`
	Limit        int32                 `json:"limit,omitempty"`
	Offset       int32                 `json:"offset,omitempty"`
}

type MeetingsMembersResponse struct {
	Items  []MeetingMember `json:"items"`
	Limit  int32           `json:"limit,omitempty"`
	Offset int32           `json:"offset,omitempty"`
	Total  int32           `json:"total,omitempty"`
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

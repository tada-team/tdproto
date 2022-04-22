package tdproto

type Meeting struct {
	Id                string            `json:"id"`
	TeamUuid          string            `json:"team_uuid"`
	GroupUuid         string            `json:"group_uuid"`
	OwnerUuid         string            `json:"owner_uuid"`
	PersonalAccountId string            `json:"personal_account_id,omitempty"`
	Title             string            `json:"title,omitempty"`
	Description       string            `json:"description,omitempty"`
	StartAt           ISODateTimeString `json:"start_at"`
	Duration          int32             `json:"duration"`
	Freq              *Freq             `json:"freq,omitempty"`
	IsArchive         bool              `json:"is_archive,omitempty"`
	IsPublic          bool              `json:"is_public,omitempty"`
	IsOutside         bool              `json:"is_outside,omitempty"`
	IsRequired        bool              `json:"is_required,omitempty"`
	CanAddMember      bool              `json:"can_add_member,omitempty"`
	CanDelete         bool              `json:"can_delete,omitempty"`
	CanEdit           bool              `json:"can_edit,omitempty"`
	CanJoin           bool              `json:"can_join,omitempty"`
	Members           []MeetingMember   `json:"meeting_members,omitempty"`
}

type MeetingsRequestParams struct {
	TeamUuid   string   `json:"team_uuid"`
	Year       int32    `json:"year"`
	Month      int32    `json:"month"`
	Day        *int32   `json:"day,omitempty"`
	Members    []string `json:"members,omitempty"`
	Limit      *int32   `json:"limit,omitempty"`
	Offset     *int32   `json:"offset,omitempty"`
	IsArchive  *bool    `json:"is_archive,omitempty"`
	IsFreq     *bool    `json:"is_freq,omitempty"`
	IsPublic   *bool    `json:"is_public,omitempty"`
	IsOutside  *bool    `json:"is_outside,omitempty"`
	IsRequired *bool    `json:"is_required,omitempty"`
}

type MeetingsResponse struct {
	Items  []Meeting `json:"items"`
	Limit  *int32    `json:"limit,omitempty"`
	Offset *int32    `json:"offset,omitempty"`
	Total  *int32    `json:"total,omitempty"`
}

type MeetingsCreateRequest struct {
	OwnerUuid   string                        `json:"owner_uuid"`
	TeamUuid    string                        `json:"team_uuid"`
	Title       string                        `json:"title,omitempty"`
	Description string                        `json:"description,omitempty"`
	StartAt     ISODateTimeString             `json:"start_at"`
	Duration    int32                         `json:"duration"`
	Freq        *Freq                         `json:"freq,omitempty"`
	Members     []MeetingsMembersCreateParams `json:"members"`
	IsPublic    bool                          `json:"is_public,omitempty"`
	IsOutside   bool                          `json:"is_outside,omitempty"`
}

type Freq struct {
	Frequency         int32                    `json:"frequency"`
	FreqDays          []int32                  `json:"freq_days,omitempty"`
	RepeatabilityType MeetingRepeatabilityType `json:"repeatability_type"`
}

type MeetingsUpdateRequest struct {
	MeetingId  string  `json:"meeting_id"`
	TeamUuid   string  `json:"team_uuid,omitempty"`
	ActiveFrom *string `json:"active_from,omitempty"`
	StartAt    *string `json:"start_at,omitempty"`
	Duration   *int32  `json:"duration,omitempty"`
	Freq       *Freq   `json:"freq,omitempty"`
	IsPublic   *bool   `json:"is_public,omitempty"`
	IsOutside  *bool   `json:"is_outside,omitempty"`
}

type MeetingsDeleteRequestParams struct {
	TeamUuid string            `json:"team_uuid"`
	Date     ISODateTimeString `json:"date,omitempty"`
}

type MeetingMember struct {
	MeetingId         string                `json:"meeting_id"`
	ChatUuid          string                `json:"chat_uuid"`
	Contact           Contact               `json:"contact"`
	Presence          MeetingPresenceStatus `json:"presence"`
	Status            MeetingMemberStatus   `json:"status"`
	CanChangePresence bool                  `json:"can_change_presence,omitempty"`
	CanChangeStatus   bool                  `json:"can_change_status,omitempty"`
	CanRemove         bool                  `json:"can_remove,omitempty"`
}

type MeetingsMembersRequestParams struct {
	MeetingId    string                 `json:"meeting_id"`
	TeamUuid     string                 `json:"team_uuid"`
	UuidSections []string               `json:"uuid_sections,omitempty"`
	Presence     *MeetingPresenceStatus `json:"presence,omitempty"`
	Status       *MeetingMemberStatus   `json:"status,omitempty"`
	Limit        *int32                 `json:"limit,omitempty"`
	Offset       *int32                 `json:"offset,omitempty"`
}

type MeetingsMembersResponse struct {
	Items  []MeetingMember `json:"items"`
	Limit  *int32          `json:"limit,omitempty"`
	Offset *int32          `json:"offset,omitempty"`
	Total  *int32          `json:"total,omitempty"`
}

type MeetingsMembersCreateRequest struct {
	Members  []MeetingsMembersCreateParams `json:"members"`
	TeamUuid string                        `json:"team_uuid"`
}
type MeetingsMembersCreateParams struct {
	Jid    JID                 `json:"jid"`
	Status MeetingMemberStatus `json:"status,omitempty"`
}

type MeetingsMembersCreateResponse struct {
	Members []MeetingMember `json:"members,omitempty"`
	Errors  []string        `json:"errors,omitempty"`
}

type MeetingsMembersUpdateRequest struct {
	Status   MeetingMemberStatus `json:"status,omitempty"`
	TeamUuid string              `json:"team_uuid,omitempty"`
}

type MeetingsMembersDeleteRequestParams struct {
	TeamUuid string `json:"team_uuid"`
}

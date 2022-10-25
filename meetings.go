package tdproto

type Meeting struct {
	Id                     string                 `json:"id"`
	TeamUuid               string                 `json:"team_uuid,omitempty"`
	OwnerContactUuid       JID                    `json:"owner_contact_uuid,omitempty"`
	OwnerUserUuid          string                 `json:"owner_user_uuid,omitempty"`
	PersonalAccountId      string                 `json:"personal_account_id,omitempty"`
	StartAt                ISODateTimeString      `json:"start_at"`
	EndAt                  ISODateTimeString      `json:"end_at"`
	Duration               int32                  `json:"duration"`
	Freq                   *Freq                  `json:"freq,omitempty"`
	IsArchive              bool                   `json:"is_archive,omitempty"`
	IsOutside              *bool                  `json:"is_outside,omitempty"`
	IsRequired             bool                   `json:"is_required,omitempty"`
	CanEdit                bool                   `json:"can_edit,omitempty"`
	MeetingMembers         []MeetingMember        `json:"meeting_members,omitempty"`
	VCSEnabled             bool                   `json:"vcs_enabled,omitempty"`
	IVCSInfo               *IVCSInfo              `json:"ivcs_info,omitempty"`
	MeetingRecipientStatus MeetingRecipientStatus `json:"meeting_recipient_status"`
	Chat
}

type IVCSInfo struct {
	ConferenceID string `json:"conference_id,omitempty"`
	JoinToken    string `json:"join_token,omitempty"`
}

type MeetingsGetRequest struct {
	TeamUuid    string   `json:"team_uuid"`
	DateFrom    string   `json:"date_from"`
	DateTo      string   `json:"date_to"`
	Limit       *int32   `json:"limit,omitempty"`
	Offset      *int32   `json:"offset,omitempty"`
	IsArchive   *bool    `json:"is_archive,omitempty"`
	IsFreq      *bool    `json:"is_freq,omitempty"`
	IsPublic    *bool    `json:"is_public,omitempty"`
	IsOutside   *bool    `json:"is_outside,omitempty"`
	IsRequired  *bool    `json:"is_required,omitempty"`
	MembersJids []string `json:"members_jids,omitempty"`
}

type MeetingsResponse struct {
	PaginatedMeetings
}

type MeetingsDatesResponse struct {
	Dates []string `json:"dates"`
}

type MeetingsCountResponse struct {
	CountCells    int32 `json:"count_cells"`
	CountMeetings int32 `json:"count_meetings"`
}

type MeetingsCreateRequest struct {
	OwnerContactUuid JID                          `json:"owner_contact_uuid"`
	TeamUuid         string                       `json:"team_uuid"`
	Title            string                       `json:"title,omitempty"`
	Description      string                       `json:"description,omitempty"`
	StartAt          string                       `json:"start_at"`
	Duration         int32                        `json:"duration"`
	Freq             *Freq                        `json:"freq,omitempty"`
	Members          []MeetingsMemberCreateParams `json:"members"`
	IsPublic         bool                         `json:"is_public,omitempty"`
	IsOutside        *bool                        `json:"is_outside,omitempty"`
	VCSEnabled       bool                         `json:"vcs_enabled,omitempty"`
}

type Freq struct {
	Frequency                int32                    `json:"frequency"`
	FreqDays                 []int32                  `json:"freq_days,omitempty"`
	RepeatabilityType        MeetingRepeatabilityType `json:"repeatability_type"`
	RepeatabilityDescription string                   `json:"repeatability_description,omitempty"`
}

type MeetingsUpdateRequest struct {
	MeetingId            string                       `json:"meeting_id"`
	TeamUuid             string                       `json:"team_uuid"`
	StartAt              *string                      `json:"start_at,omitempty"`
	Duration             *int32                       `json:"duration,omitempty"`
	Freq                 *Freq                        `json:"freq,omitempty"`
	IsPublic             *bool                        `json:"is_public,omitempty"`
	IsOutside            *bool                        `json:"is_outside,omitempty"`
	Title                *string                      `json:"title,omitempty"`
	Description          *string                      `json:"description,omitempty"`
	AddMembers           []MeetingsMemberCreateParams `json:"add_members,omitempty"`
	RemoveMembers        []JID                        `json:"remove_members,omitempty"`
	NotificationsEnabled *bool                        `json:"notifications_enabled,omitempty"`
	CountersEnabled      *bool                        `json:"counters_enabled,omitempty"`
	VCSEnabled           *bool                        `json:"vcs_enabled,omitempty"`
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

type MeetingsMembersGetRequestParams struct {
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
	Members  []MeetingsMemberCreateParams `json:"members"`
	TeamUuid string                       `json:"team_uuid"`
}

type MeetingsMemberCreateParams struct {
	Jid    JID                 `json:"jid"`
	Status MeetingMemberStatus `json:"status,omitempty"`
}

type MeetingsMembersCreateResponse struct {
	Members []MeetingMember `json:"members,omitempty"`
	Errors  []string        `json:"errors,omitempty"`
}

type MeetingsMembersUpdateRequest struct {
	Status   MeetingMemberStatus `json:"status,omitempty"`
	TeamUuid string              `json:"team_uuid"`
}

type MeetingsMembersDeleteRequestParams struct {
	TeamUuid string `json:"team_uuid"`
}

type MeetingsMembersBatchDeleteRequestParams struct {
	TeamUuid    string `json:"team_uuid"`
	MembersJids []JID  `json:"members_jids"`
}

type MeetingsGetFrequencyDescriptionParams struct {
	Frequency         int32                    `json:"frequency"`
	FreqDays          *string                  `json:"freq_days,omitempty"`
	RepeatabilityType MeetingRepeatabilityType `json:"repeatability_type"`
}

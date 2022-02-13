package tdproto

import "time"

type SimpleDailyIntStats map[time.Time]int

// Active daily users statistics
type ActiveUserDailyStat struct {
	Day              time.Time `json:"day"`
	CallsCount       *int      `json:"calls_count,omitempty"`
	FamilyName       *string   `json:"family_name,omitempty"`
	GivenName        *string   `json:"given_name,omitempty"`
	Patronymic       *string   `json:"patronymic,omitempty"`
	Phone            *string   `json:"phone,omitempty"`
	MessagesCount    *int      `json:"messages_count,omitempty"`
	CallSecondsTotal *int      `json:"call_seconds_total,omitempty"`
	UserId           int       `json:"user_id"`
}

type ActiveUserDailyStatList []ActiveUserDailyStat

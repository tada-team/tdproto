package tdproto

import "time"

type SimpleDailyIntStats map[time.Time]int

type ActiveUserDailyStat struct {
	Day              time.Time `json:"day"`
	UserId           int       `json:"user_id"`
	FamilyName       *string   `json:"family_name,omitempty"`
	GivenName        *string   `json:"given_name,omitempty"`
	Patronymic       *string   `json:"patronymic,omitempty"`
	Phone            *string   `json:"phone,omitempty"`
	MessagesCount    *int      `json:"messages_count,omitempty"`
	CallsCount       *int      `json:"calls_count,omitempty"`
	CallSecondsTotal *int      `json:"call_seconds_total,omitempty"`
}

type ActiveUserDailyStatList []ActiveUserDailyStat

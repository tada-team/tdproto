package tdproto

type User struct {
	Phone            *string `json:"phone"`
	Email            *string `json:"email"`
	DefaultLang      *string `json:"default_lang"`
	AltSend          bool    `json:"alt_send"`
	AlwaysSendPushes bool    `json:"always_send_pushes"`
	UnreadFirst      bool    `json:"unread_first"`
	MUnreadFirst     bool    `json:"munread_first"`
	Timezone         string  `json:"timezone"`
	QuietTimeStart   *string `json:"quiet_time_start"`
	QuietTimeFinish  *string `json:"quiet_time_finish"`
}

type UserWithMe struct {
	User
	Inviter *JID         `json:"inviter,omitempty"`
	Teams   []Team       `json:"teams"`
	Devices []PushDevice `json:"devices"`
}

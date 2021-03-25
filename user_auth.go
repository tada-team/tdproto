package tdproto

import "time"

const (
	AuthBySecureCookie   = "cookie"
	AuthByInsecureCookie = "insecure cookie"
	AuthByToken          = "token"
	AuthByBotToken       = "bot token"
)

type UserAuth struct {
	Created    time.Time   `json:"created"`
	LastAccess time.Time   `json:"last_access,omitempty"`
	DebugAge   int         `json:"_age,omitempty"`
	Uid        string      `json:"uid"`
	Kind       string      `json:"kind"`
	Addr       string      `json:"addr"`
	Useragent  string      `json:"user_agent"`
	Device     *PushDevice `json:"device,omitempty"`
}

package tdproto

import "time"

const (
	AuthBySecureCookie   = "cookie"
	AuthByInsecureCookie = "insecure cookie"
	AuthByToken          = "token"
	AuthByBotToken       = "bot token"
)

// User authentication
type UserAuth struct {
	Created    time.Time `json:"created"`
	LastAccess time.Time `json:"last_access,omitempty"`

	// info about device (struct)
	Device *PushDevice `json:"device,omitempty"`

	// name of region
	Region string `json:"region,omitempty"`

	//type of auth
	Kind string `json:"kind"`

	// ip address
	Addr string `json:"addr,omitempty"`

	// info about useragent
	Useragent string `json:"user_agent,omitempty"`

	// name of country
	Country  string `json:"country,omitempty"`
	Uid      string `json:"uid"`
	DebugAge int    `json:"_age,omitempty"`
}

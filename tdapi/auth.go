package tdapi

import "github.com/tada-team/tdproto"

type Auth struct {
	Token string             `json:"token,omitempty"`
	Me    tdproto.UserWithMe `json:"me"`
}

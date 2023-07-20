package tdapi

import "github.com/tada-team/tdproto"

type PublicStatusResp struct {
	All []tdproto.PublicStatus `json:"all"`
}

package tdapi

import "github.com/tada-team/tdproto"

type ReleasesResp struct {
	Android []tdproto.Dist `json:"android,omitempty"`
	Linux   []tdproto.Dist `json:"linux,omitempty"`
	Mac     []tdproto.Dist `json:"mac,omitempty"`
	Win     []tdproto.Dist `json:"win,omitempty"`
}

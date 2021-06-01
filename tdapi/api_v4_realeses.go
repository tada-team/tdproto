package tdapi

import (
	"github.com/tada-team/tdproto"
	"github.com/tada-team/tdproto/tdapi/openapi"
)

func init() {
	register("/api/v4/releases", openapi.Path{
		Get: &openapi.Operation{
			Summary: "Links to releases",
			Responses: openapi.Responses{
				Status200: okResponse(openapi.SchemaFromType(ReleasesResp{})),
			},
		},
	})
}

type ReleasesResp struct {
	Android []tdproto.Dist `json:"android,omitempty"`
	Linux   []tdproto.Dist `json:"linux,omitempty"`
	Mac     []tdproto.Dist `json:"mac,omitempty"`
	Win     []tdproto.Dist `json:"win,omitempty"`
}

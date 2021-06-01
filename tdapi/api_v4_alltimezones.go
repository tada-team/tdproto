package tdapi

import (
	"github.com/tada-team/tdproto/tdapi/openapi"
)

func init() {
	register("/api/v4/alltimezones", openapi.Path{
		Get: &openapi.Operation{
			Summary: "Timezones list",
			Responses: openapi.Responses{
				Status200: okResponse(openapi.SchemaFromType(TimezonesResp{})),
			},
		},
	})
}

type TimezonesResp struct {
	Timezones []string `json:"timezones"`
}

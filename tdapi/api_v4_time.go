package tdapi

import (
	"github.com/tada-team/tdproto/tdapi/openapi"
)

func init() {
	register("/api/v4/time", openapi.Path{
		Get: &openapi.Operation{
			Summary: "Server time",
			Responses: openapi.Responses{
				Status200: okResponse(openapi.Schema{
					Type:    openapi.String,
					Format:  openapi.DateTime,
					Example: "2021-04-09T15:06:46.555215Z",
				}),
			},
		},
	})
}

package tdapi

import (
	"github.com/tada-team/tdproto/tdapi/openapi"
)

func init() {
	register("/api/v4/ping", openapi.Path{
		Get: &openapi.Operation{
			Summary: "Server ping",
			Responses: openapi.Responses{
				Status200: okResponse(openapi.Schema{
					Type:    openapi.String,
					Example: "pong",
				}),
			},
		},
	})
}

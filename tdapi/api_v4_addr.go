package tdapi

import (
	"github.com/tada-team/tdproto/tdapi/openapi"
)

func init() {
	register("/api/v4/addr", openapi.Path{
		Get: &openapi.Operation{
			Summary: "Client address, for debug",
			Responses: openapi.Responses{
				Status200: okResponse(openapi.Property{
					Type:    openapi.String,
					Example: "127.0.0.1",
				}),
			},
		},
	})
}

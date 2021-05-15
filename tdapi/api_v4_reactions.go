package tdapi

import (
	"github.com/tada-team/tdproto"
	"github.com/tada-team/tdproto/tdapi/openapi"
)

func init() {
	register("/api/v4/reactions", openapi.Path{
		Get: &openapi.Operation{
			Summary: "All available message reactions",
			Responses: openapi.Responses{
				Status200: okResponse(openapi.SchemaFromType([]tdproto.Reaction{})),
			},
		},
	})
}

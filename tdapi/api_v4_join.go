package tdapi

import (
	"github.com/tada-team/tdproto"
	"github.com/tada-team/tdproto/tdapi/openapi"
)

func init() {
	register("/api/v4/join/[token]", openapi.Path{
		Parameters: []openapi.Parameter{openapi.PathParameter("token", "Invitation token")},
		Post: &openapi.Operation{
			Summary: "Join to the team",
			Responses: openapi.Responses{
				Status200: okResponse(openapi.SchemaFromType(tdproto.Team{})),
			},
		},
	})

	register("/api/v4/join/[token]/preview", openapi.Path{
		Parameters: []openapi.Parameter{openapi.PathParameter("token", "Invitation token")},
		Get: &openapi.Operation{
			Summary: "Invitation information",
			Responses: openapi.Responses{
				Status200: okResponse(openapi.SchemaFromType(tdproto.TeamShort{})),
			},
		},
	})
}
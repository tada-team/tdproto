package tdapi

import (
	"github.com/tada-team/tdproto"
	"github.com/tada-team/tdproto/tdapi/openapi"
)

func init() {
	register("features.json", openapi.Path{
		Get: &openapi.Operation{
			Summary: "Server information",
			Responses: openapi.Responses{
				Status200: okResponse(openapi.SchemaFromType(tdproto.Features{})),
			},
		},
	})
}

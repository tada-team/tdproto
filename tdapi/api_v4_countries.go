package tdapi

import (
	"github.com/tada-team/tdproto"
	"github.com/tada-team/tdproto/tdapi/openapi"
)

func init() {
	register("/api/v4/alltimezones", openapi.Path{
		Get: &openapi.Operation{
			Summary: "Countries list with phone codes",
			Responses: openapi.Responses{
				Status200: okResponse(openapi.SchemaFromType([]tdproto.Country{})),
			},
		},
	})
}

package tdapi

import (
	"github.com/tada-team/tdproto"
	"github.com/tada-team/tdproto/tdapi/openapi"
)

func init() {
	register("/api/v4/join/[token]/preview", openapi.Path{
		Parameters: []openapi.Parameter{{
			Name:        "token",
			In:          "path",
			Required:    true,
			Description: "Invitation token",
			Schema: openapi.Schema{
				Type: openapi.String,
			},
		}},
		Get: &openapi.Operation{
			Summary: "Invitation information",
			Responses: openapi.Responses{
				Status200: okResponse(openapi.Property{
					Type: openapi.Object,
					Properties: map[string]openapi.Property{
						"uid": {
							Type:        openapi.String,
							Description: "Team id",
						},
						"name": {
							Type:        openapi.String,
							Description: "Team name",
						},
						"icons": {
							Type:        openapi.String,
							Description: "Team icons",
							Properties: map[string]openapi.Property{
								"stub": {
									Type:        openapi.String,
									Description: "Generated image with 1-2 letters",
								},
								"letters": {
									Type:        openapi.String,
									Description: "Letters from stub icon",
								},
								"color": {
									Type:        openapi.String,
									Description: "Stub icon background color",
								},
							},
						},
					},
					Example: tdproto.TeamShort{
						Uid:  "123e4567-e89b-12d3-a456-426614174000",
						Name: "Test Team",
						Icons: tdproto.IconData{
							Stub:    "https://web.tada.team/a/e36659:tt/256.png",
							Letters: "TT",
							Color:   "#E36659",
						},
					},
				}),
			},
		},
	})
}

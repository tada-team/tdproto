package tdapi

import (
	"github.com/tada-team/tdproto/tdapi/openapi"
)

func init() {
	register("/api/v4/join/[token]/preview", openapi.Path{
		Parameters: []openapi.Parameter{openapi.PathParameter("token", "Invitation token")},
		Get: &openapi.Operation{
			Summary: "Invitation information",
			Responses: openapi.Responses{
				Status200: okResponse(openapi.ObjectProperty(
					"Short team representation. For invites, push notifications, etc. Readonly.",
					map[string]openapi.Property{
						"uid":  openapi.StringProperty("Team id", "123e4567-e89b-12d3-a456-426614174000"),
						"name": openapi.StringProperty("Team name", "Test Team"),
						"icons": openapi.ObjectProperty("Team icons", map[string]openapi.Property{
							"stub":    openapi.StringProperty("Generated image with 1-2 letters", "https://web.tada.team/a/e36659:tt/256.png"),
							"letters": openapi.StringProperty("Letters from stub icon", "TT"),
							"color":   openapi.StringProperty("Stub icon background color", "#E36659"),
						}, nil),
					},
					nil,
				)),
			},
		},
	})
}

package tdapi

import (
	"github.com/tada-team/tdproto/tdapi/openapi"
)

func init() {
	register("/api/v4/time", openapi.Path{
		Get: &openapi.Operation{
			Summary: "Server time",
			Responses: openapi.Responses{
				Status200: &openapi.Response{
					Description: "OK",
					Content: openapi.Contents{
						ApplicationJSON: &openapi.Content{
							Schema: openapi.Schema{
								Type: openapi.Object,
								Properties: map[string]openapi.Property{
									"ok": {
										Type: openapi.Boolean,
									},
									"result": {
										Type:    openapi.String,
										Format:  openapi.DateTime,
										Example: "2021-04-09T15:06:46.555215Z",
									},
								},
							},
						},
					},
				},
			},
		},
	})
}

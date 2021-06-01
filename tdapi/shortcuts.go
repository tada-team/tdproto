package tdapi

import "github.com/tada-team/tdproto/tdapi/openapi"

func okResponse(result openapi.Schema) *openapi.Response {
	return &openapi.Response{
		Description: "OK",
		Content: openapi.Contents{
			ApplicationJSON: &openapi.Content{
				Schema: openapi.Schema{
					Type: openapi.Object,
					Properties: map[string]openapi.Schema{
						"ok": {
							Type:    openapi.Boolean,
							Example: true,
						},
						"result": result,
					},
				},
			},
		},
	}
}


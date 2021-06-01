package tdapi

import (
	"github.com/tada-team/tdproto"
	"github.com/tada-team/tdproto/tdapi/openapi"
)

func init() {
	register("/api/v4/emoji", openapi.Path{
		Get: &openapi.Operation{
			Summary: "Emoji list",
			Responses: openapi.Responses{
				Status200: okResponse(openapi.SchemaFromType(EmojiResp{})),
			},
		},
	})
}

type EmojiResp struct {
	All []tdproto.Emoji `json:"all"`
}

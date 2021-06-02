package tdapi

import "github.com/tada-team/tdproto"

type EmojiResp struct {
	All []tdproto.Emoji `json:"all"`
}

package api_paths

import "github.com/tada-team/tdproto"

var ChatPaths = []PathSpec{
	{
		Path: "/api/v4/teams/{team_id}/chats/{chat_id}",
		Get: &HttpSpec{
			Responce:    tdproto.Chat{},
			Description: "Get the chat information.",
		},
		Put: &HttpSpec{
			Request:  tdproto.Chat{},
			Responce: tdproto.Chat{},
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/chats/{chat_id}/messages",
		Post: &HttpSpec{
			Request: tdproto.Message{},
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/chats/{chat_id}/messages/{message_id}",
		Post: &HttpSpec{
			Request:  tdproto.Message{},
			Responce: tdproto.Message{},
		},
		Delete: &HttpSpec{
			Responce: tdproto.Message{},
		},
	},
}

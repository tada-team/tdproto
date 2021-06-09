package api_paths

import (
	"github.com/tada-team/tdproto"
	"github.com/tada-team/tdproto/tdapi"
)

// TODO: http:get:: /api/v4/teams/{team_uid}/chats/{chat_jid}/messages

var ChatPaths = []PathSpec{
	{
		Path: "/api/v4/teams/{team_id}/chats/{chat_id}",
		Get: &OperationSpec{
			Responce:    tdproto.Chat{},
			Description: "Get the chat information.",
		},
		Put: &OperationSpec{
			Request:     tdproto.Chat{},
			Responce:    tdproto.Chat{},
			Description: "Change chat settings.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/chats/{chat_id}/messages",
		Get: &OperationSpec{
			QueryStruct: tdapi.MessageFilter{},
			Responce:    []tdproto.Message{},
		},
		Post: &OperationSpec{
			Request:     tdproto.Message{},
			Description: "Send text message to chat.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/chats/{chat_id}/messages/{message_id}",
		Post: &OperationSpec{
			Request:     tdproto.Message{},
			Responce:    tdproto.Message{},
			Description: "Edit message.",
		},
		Delete: &OperationSpec{
			Responce:    tdproto.Message{},
			Description: "Delete message.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/chats/{contact_id}/messages",
		Post: &OperationSpec{
			Request:     tdproto.Message{},
			Description: "Send text message to direct chat.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/chats/{contact_id}/messages/{message_id}",
		Post: &OperationSpec{
			Request:     tdproto.Message{},
			Responce:    tdproto.Message{},
			Description: "Edit message in direct chat.",
		},
		Delete: &OperationSpec{
			Responce:    tdproto.Message{},
			Description: "Delete message in direct chat.",
		},
	},
}

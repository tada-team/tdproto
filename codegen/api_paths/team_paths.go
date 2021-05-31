package api_paths

import "github.com/tada-team/tdproto"

var TeamPaths = []PathSpec{
	{
		Path: "/api/v4/teams",
		Get: &HttpSpec{
			Responce:            []tdproto.Team{},
			Description:         "Get the list of teams on the server.",
			ResponceDescription: "List of :ref:`tdproto-Team` objects.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}",
		Get: &HttpSpec{
			Responce: tdproto.Team{},
		},
		Put: &HttpSpec{
			Request:  tdproto.Team{},
			Responce: tdproto.Team{},
		},
		Delete: &HttpSpec{
			Responce: tdproto.Team{},
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/chats",
		Get: &HttpSpec{
			Responce: []tdproto.Chat{},
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/contacts",
		Get: &HttpSpec{
			Responce: []tdproto.Contact{},
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/contacts/{contact_id}",
		Get: &HttpSpec{
			Responce: tdproto.Contact{},
		},
		Put: &HttpSpec{
			Request:  tdproto.Contact{},
			Responce: tdproto.Contact{},
		},
		Delete: &HttpSpec{
			Responce: tdproto.Contact{},
		},
	},
}

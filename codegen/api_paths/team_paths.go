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
			Responce:            tdproto.Team{},
			Description:         "Get team info.",
			ResponceDescription: ":ref:`tdproto-Team` object.",
		},
		Put: &HttpSpec{
			Request:             tdproto.Team{},
			Responce:            tdproto.Team{},
			RequestDescription:  ":ref:`tdproto-Team` object with updated fields.",
			ResponceDescription: "Updated :ref:`tdproto-Team` object of the team.",
			Description:         "Update team settings. Must have admin rights.",
		},
		Delete: &HttpSpec{
			Responce:            tdproto.Team{},
			ResponceDescription: ":ref:`tdproto-Team` object of deleted team.",
			Description:         "Delete the team. Must have admin rights.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/chats",
		Get: &HttpSpec{
			Responce:            []tdproto.Chat{},
			Description:         "Get the list of chats in the team.",
			ResponceDescription: "List of :ref:`tdproto-Chat` objects.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/contacts",
		Get: &HttpSpec{
			Responce:            []tdproto.Contact{},
			Description:         "Get the list of contacts of the team.",
			ResponceDescription: "List of :ref:`tdproto-Contact` objects.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/contacts/{contact_id}",
		Get: &HttpSpec{
			Responce:            tdproto.Contact{},
			Description:         "Get contact details.",
			ResponceDescription: "The :ref:`tdproto-Contact` object.",
		},
		Post: &HttpSpec{
			Request:             tdproto.Contact{},
			Responce:            tdproto.Contact{},
			Description:         "Update contact details.",
			ResponceDescription: "Updated :ref:`tdproto-Contact` object.",
		},
		Delete: &HttpSpec{
			Responce:            tdproto.Contact{},
			Description:         "Remove contact from the team.",
			ResponceDescription: "Removed :ref:`tdproto-Contact` object.",
		},
	},
}

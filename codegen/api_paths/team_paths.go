package api_paths

import "github.com/tada-team/tdproto"

var TeamPaths = []PathSpec{
	{
		Path: "/api/v4/teams",
		Get: &HttpSpec{
			Responce:    []tdproto.Team{},
			Description: "Get the list of teams on the server.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}",
		Get: &HttpSpec{
			Responce:    tdproto.Team{},
			Description: "Get team info.",
		},
		Put: &HttpSpec{
			Request:             tdproto.Team{},
			Responce:            tdproto.Team{},
			Description:         []string{"Update team settings.", "Must have admin rights."},
			RequestDescription:  ":ref:`tdproto-Team` object with updated fields.",
			ResponceDescription: "Updated :ref:`tdproto-Team` object of the team.",
		},
		Delete: &HttpSpec{
			Responce:            tdproto.Team{},
			ResponceDescription: ":ref:`tdproto-Team` object of deleted team.",
			Description:         []string{"Delete the team.", "Must have admin rights."},
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/chats",
		Get: &HttpSpec{
			Responce:    []tdproto.Chat{},
			Description: "Get the list of chats in the team.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/contacts",
		Get: &HttpSpec{
			Responce:    []tdproto.Contact{},
			Description: "Get the list of contacts of the team.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/contacts/{contact_id}",
		Get: &HttpSpec{
			Responce:    tdproto.Contact{},
			Description: "Get contact details.",
		},
		Post: &HttpSpec{
			Request:             tdproto.Contact{},
			Responce:            tdproto.Contact{},
			Description:         "Update contact details.",
			ResponceDescription: "Updated :ref:`tdproto-Contact` object.",
		},
		Delete: &HttpSpec{
			Responce:            tdproto.Contact{},
			Description:         []string{"Remove contact from the team.", "Must have admin rights."},
			ResponceDescription: "Removed :ref:`tdproto-Contact` object.",
		},
	},
}

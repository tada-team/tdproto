package api_paths

import "github.com/tada-team/tdproto"

var TeamPaths = []PathSpec{
	{
		Path: "/api/v4/teams",
		Get: &OperationSpec{
			Response:    []tdproto.Team{},
			Description: "Get the list of teams on the server.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}",
		Get: &OperationSpec{
			Response:    tdproto.Team{},
			Description: "Get team info.",
		},
		Put: &OperationSpec{
			Request:             tdproto.Team{},
			Response:            tdproto.Team{},
			Description:         []string{"Update team settings.", "Must have admin rights."},
			RequestDescription:  ":tdproto:ref:`Team` object with updated fields.",
			ResponseDescription: "Updated :tdproto:ref:`Team` object of the team.",
		},
		Delete: &OperationSpec{
			Response:            tdproto.Team{},
			ResponseDescription: ":tdproto:ref:`Team` object of deleted team.",
			Description:         []string{"Delete the team.", "Must have admin rights."},
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/chats",
		Get: &OperationSpec{
			Response:    []tdproto.Chat{},
			Description: "Get the list of chats in the team.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/contacts",
		Get: &OperationSpec{
			Response:    []tdproto.Contact{},
			Description: "Get the list of contacts of the team.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/contacts/{contact_id}",
		Get: &OperationSpec{
			Response:    tdproto.Contact{},
			Description: "Get contact details.",
		},
		Post: &OperationSpec{
			Request:             tdproto.Contact{},
			Response:            tdproto.Contact{},
			Description:         "Update contact details.",
			ResponseDescription: "Updated :tdproto:ref:`Contact` object.",
		},
		Delete: &OperationSpec{
			Response:            tdproto.Contact{},
			Description:         []string{"Remove contact from the team.", "Must have admin rights."},
			ResponseDescription: "Removed :tdproto:ref:`Contact` object.",
		},
	},
}

package api_paths

import "github.com/tada-team/tdproto"

var GroupPaths = []PathSpec{
	{
		Path: "/api/v4/teams/{team_id}/groups",
		Get: &HttpSpec{
			Responce:    []tdproto.Chat{},
			Description: "Get all groups in the team.",
		},
		Delete: &HttpSpec{
			Description: "Delete the group.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/groups/{group_id}/members",
		Get: &HttpSpec{
			Responce:    []tdproto.GroupMembership{},
			Description: "Get the list of group members.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/groups/{group_id}/members/{contact_id}",
		Delete: &HttpSpec{
			Responce:    tdproto.GroupMembership{},
			Description: "Remove member from the group.",
		},
	},
}

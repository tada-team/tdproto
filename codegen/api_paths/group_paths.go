package api_paths

import "github.com/tada-team/tdproto"

var GroupPaths = []PathSpec{
	{
		Path: "/api/v4/teams/{team_id}/groups",
		Get: &OperationSpec{
			Response:    []tdproto.Chat{},
			Description: "Get all groups in the team.",
		},
		Delete: &OperationSpec{
			Description: "Delete the group.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/groups/{group_id}/members",
		Get: &OperationSpec{
			Response:    []tdproto.GroupMembership{},
			Description: "Get the list of group members.",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/groups/{group_id}/members/{contact_id}",
		Delete: &OperationSpec{
			Response:    tdproto.GroupMembership{},
			Description: "Remove member from the group.",
		},
	},
}

package api_paths

import "github.com/tada-team/tdproto"

var GroupPaths = map[string]PathSpec{
	"/api/v4/teams/{team_id}/groups": {
		Get: &HttpSpec{
			Responce: []tdproto.Chat{},
		},
		Delete: &HttpSpec{},
	},
	"/api/v4/teams/{team_id}/groups/{group_id}/members": {
		Get: &HttpSpec{
			Responce: []tdproto.GroupMembership{},
		},
	},
	"/api/v4/teams/{team_id}/groups/{group_id}/members/{contact_id}": {
		Delete: &HttpSpec{
			Responce: tdproto.GroupMembership{},
		},
	},
}

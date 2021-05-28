package api_paths

import "github.com/tada-team/tdproto"

var TeamPaths = map[string]PathSpec{
	"/api/v4/teams": {
		Get: &HttpSpec{
			Responce: []tdproto.Team{},
		},
	},
	"/api/v4/teams/{team_id}": {
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
	"/api/v4/teams/{team_id}/chats": {
		Get: &HttpSpec{
			Responce: []tdproto.Chat{},
		},
	},
	"/api/v4/teams/{team_id}/contacts": {
		Get: &HttpSpec{
			Responce: []tdproto.Contact{},
		},
	},
	"/api/v4/teams/{team_id}/contacts/{contact_id}": {
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

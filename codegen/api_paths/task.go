package api_paths

import (
	"github.com/tada-team/tdproto/tdapi"
)

var TaskPaths = []PathSpec{
	{
		Path: "/api/v4/teams/{team_id}/tasks",
		Post: &OperationSpec{
			Request:     tdapi.Task{},
			Responce:    tdapi.Task{},
			Description: "Create new task",
		},
	},
}

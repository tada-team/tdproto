package api_paths

import (
	"github.com/tada-team/tdproto/tdapi"
)

var TaskPaths = []PathSpec{
	{
		Path: "/api/v4/teams/{team_id}/tasks",
		Post: &OperationSpec{
			Request:     tdapi.Task{},
			Response:    tdapi.Task{},
			Description: "Create new task",
		},
		Get: &OperationSpec{
			QueryStruct: tdapi.TaskFilter{},
			Response:    []tdapi.Task{},
			Description: "Get the list of tasks",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/tasks/{task_id}",
		Get: &OperationSpec{
			Response:    tdapi.Task{},
			Description: "Get task",
		},
		Post: &OperationSpec{
			Request:     tdapi.Task{},
			Response:    tdapi.Task{},
			Description: "Update task",
		},
		Delete: &OperationSpec{
			Response:    tdapi.Task{},
			Description: "Delete task",
		},
	},
}

package api_paths

import (
	"github.com/tada-team/tdproto/tdforms"
	"github.com/tada-team/tdproto/tdquery"
)

var TaskPaths = []PathSpec{
	{
		Path: "/api/v4/teams/{team_id}/tasks",
		Post: &OperationSpec{
			Request:     tdforms.Task{},
			Response:    tdforms.Task{},
			Description: "Create new task",
		},
		Get: &OperationSpec{
			QueryStruct: tdquery.Tasks{},
			Response:    []tdforms.Task{},
			Description: "Get the list of tasks",
		},
	},
	{
		Path: "/api/v4/teams/{team_id}/tasks/{task_id}",
		Get: &OperationSpec{
			Response:    tdforms.Task{},
			Description: "Get task",
		},
		Post: &OperationSpec{
			Request:     tdforms.Task{},
			Response:    tdforms.Task{},
			Description: "Update task",
		},
		Delete: &OperationSpec{
			Response:    tdforms.Task{},
			Description: "Delete task",
		},
	},
}

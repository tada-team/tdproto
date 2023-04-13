package api_paths

import "github.com/tada-team/tdproto/codegen/openapi/tags"

type OperationSpec struct {
	Request             interface{}
	Response            interface{}
	Description         interface{}
	ResponseDescription string
	RequestDescription  string
	SecurityIsOptional  bool
	QueryStruct         interface{}
}

type PathSpec struct {
	Path   string
	Get    *OperationSpec
	Put    *OperationSpec
	Delete *OperationSpec
	Post   *OperationSpec
	Tags   []string
}

var AllPaths = map[string][]PathSpec{
	tags.GroupsTag:  GroupPaths,
	tags.TeamsTag:   TeamPaths,
	tags.ChatsTag:   ChatPaths,
	tags.MiscTag:    MiscPaths,
	tags.TasksTag:   TaskPaths,
	tags.BillingTag: BillingPaths,
}

var PathTitles = map[string]string{
	"group":   "Group related paths",
	"team":    "Team related paths",
	"chat":    "Chat related paths",
	"misc":    "Miscellaneous paths",
	"task":    "Task related paths",
	"billing": "Billing related paths",
}

package api_paths

type OperationSpec struct {
	Request             interface{}
	Response            interface{}
	Description         interface{}
	ResponseDescription string
	RequestDescription  string
	SecurityIsOptional  bool
}

type PathSpec struct {
	Path   string
	Get    *OperationSpec
	Put    *OperationSpec
	Delete *OperationSpec
	Post   *OperationSpec
}

var AllPaths = map[string][]PathSpec{
	"group": GroupPaths,
	"team":  TeamPaths,
	"chat":  ChatPaths,
	"misc":  MiscPaths,
	"task":  TaskPaths,
}

var PathTitles = map[string]string{
	"group": "Group related paths",
	"team":  "Team related paths",
	"chat":  "Chat related paths",
	"misc":  "Miscellaneous paths",
	"task":  "Task related paths",
}

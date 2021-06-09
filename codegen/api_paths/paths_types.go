package api_paths

type OperationSpec struct {
	Request             interface{}
	Responce            interface{}
	Description         interface{}
	ResponceDescription string
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
	"misc":  "Miscelaneos paths",
	"task":  "Task related paths",
}

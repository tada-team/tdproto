package api_paths

type HttpSpec struct {
	Request             interface{}
	Responce            interface{}
	Description         string
	ResponceDescription string
}

type PathSpec struct {
	Path   string
	Get    *HttpSpec
	Put    *HttpSpec
	Delete *HttpSpec
	Post   *HttpSpec
}

var AllPaths = map[string][]PathSpec{
	"group": GroupPaths,
	"team":  TeamPaths,
	"chat":  ChatPaths,
}

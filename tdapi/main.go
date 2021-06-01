package tdapi

type PathSpec struct {
	Input  interface{}
	Output interface{}
}

var Paths = make(map[string]PathSpec)

func registerPath(path string, spec PathSpec) {

	if _, ok := Paths[path]; ok {
		panic("path already registered")
	}
	Paths[path] = spec
}

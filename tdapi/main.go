package tdapi

import (
	"github.com/tada-team/tdproto/tdapi/openapi"
)

var paths = make(map[string]openapi.Path)

func GetPaths() map[string]openapi.Path {
	return paths
}

func register(path string, spec openapi.Path) {
	if _, ok := paths[path]; ok {
		panic("path already registered")
	}
	paths[path] = spec
}

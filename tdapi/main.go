package tdapi

import (
	"github.com/tada-team/tdproto/tdapi/openapi"
	"strings"
)

var paths = make(map[string]openapi.Path)

func GetPaths() map[string]openapi.Path {
	return paths
}

func register(path string, spec openapi.Path) {
	path = strings.ReplaceAll(path, "[", "{")
	path = strings.ReplaceAll(path, "]", "}")
	if _, ok := paths[path]; ok {
		panic("path already registered")
	}
	paths[path] = spec
}

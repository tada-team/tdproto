package tdmarkup

import "testing"

func TestIsPath(t *testing.T) {
	for _, path := range []string{
		"/aa/bb/",
		"/aa/bb/",
		"/aa/bb/ccc/",
	} {
		if !isPath(path) {
			t.Errorf("`%s` must be path", path)
		}
	}

	for _, s := range []string{
		"/",
		"//",
		"/a/",
		"/aa/",
		"/aa/bb",
	} {
		if isPath(s) {
			t.Errorf("`%s` must be not path", s)
		}
	}
}

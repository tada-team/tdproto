package tdproto

import (
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
)

func SourceDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func Gentime() int64 {
	return time.Now().UnixNano()
}

func ValidUid(s string) bool {
	if s == "" {
		return false
	}
	_, err := uuid.Parse(s)
	return err == nil
}

func IsoDatetime(dt time.Time) string {
	s := dt.UTC().Format("2006-01-02T15:04:05.000000-0700")
	return strings.Replace(s, "+0000", "Z", 1)
}

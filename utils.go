package tdproto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

func Gentime() int64 {
	return time.Now().UnixNano()
}

func ValidUid(s string) bool {
	if s == "" {
		return false
	}
	_, err := uuid.FromString(s)
	return err == nil
}

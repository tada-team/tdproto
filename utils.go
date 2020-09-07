package tdproto

import uuid "github.com/satori/go.uuid"

func IsValidUid(s string) bool {
	if s == "" {
		return false
	}
	_, err := uuid.FromString(s)
	return err == nil
}

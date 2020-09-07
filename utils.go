package tdproto

import (
	"crypto/rand"
	"log"
	"strings"
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

func IsoDatetime(dt time.Time) string {
	s := dt.UTC().Format("2006-01-02T15:04:05.000000-0700")
	return strings.Replace(s, "+0000", "Z", 1)
}

func NullableIsoDatetime(dt *time.Time) *string {
	if dt == nil {
		return nil
	}
	v := IsoDatetime(*dt)
	return &v
}

func ConfirmId() string {
	return randomSymbols(12, "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
}

func randomSymbols(n int, letters string) string {
	bytes, err := randomBytes(n)
	if err != nil {
		log.Panicln("random bytes fail:", err)
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}

func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

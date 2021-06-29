package tdws

import (
	"crypto/rand"
	"log"
)

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

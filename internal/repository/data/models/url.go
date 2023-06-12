package models

import (
	"math/rand"
	"strings"
	"time"
)

type URL struct {
	ID      int           `json:"id"`
	LongVer string        `json:"longVer"`
	SrtVer  string        `json:"shortVer"`
	Exp     time.Duration `json:"expires"`
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func New(userID int, longURL string) (*URL, error) {

	return nil, nil
}

func generateRandomString() string {
	return intToBase62(rand.Uint64())
}

func intToBase62(number uint64) string {
	length := len(alphabet)
	var encoder strings.Builder
	encoder.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encoder.WriteByte(alphabet[number%uint64(length)])
	}

	return encoder.String()
}

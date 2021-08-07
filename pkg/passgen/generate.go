package passgen

import (
	"math/rand"
	"time"
)

const (
	DIGITS     = "0123456789"
	CHAR_LOWER = "abcdefghijklmnopqrstuvwxyz"
	CHAR_UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SYMBOLS    = "~`!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/"
)

func generatePassword(length int) string {
	characters := getCharString()
	password := ""

	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < length; i++ {
		index := rand.Intn(len(characters))
		password = password + string(characters[index])
	}

	return password
}

func getCharString() string {
	s := DIGITS + CHAR_LOWER + CHAR_UPPER + SYMBOLS

	return s
}

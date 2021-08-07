package passgen

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	DIGITS     = "0123456789"
	CHAR_LOWER = "abcdefghijklmnopqrstuvwxyz"
	CHAR_UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SYMBOLS    = "~`!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/"
)

func generatePassword() string {
	characters := getCharString()
	password := ""

	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 12; i++ {
		index := rand.Intn(len(characters))
		password = password + string(characters[index])
	}

	return password
}

func getCharString() string {
	args := getFlags()

	fmt.Println(args)

	s := DIGITS + CHAR_LOWER + CHAR_UPPER + SYMBOLS

	return s
}

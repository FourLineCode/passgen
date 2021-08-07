package passgen

import (
	"math/rand"
	"time"
)

const (
	DIGITS     = "01234567890123456789"
	CHAR_LOWER = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
	CHAR_UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SYMBOLS    = "~!@#$%&*()_-+={[}]|\\:;<>?/"
)

func generatePassword(args Flags) string {
	characters := getCharString(args)
	generatedPassword := ""

	rand.Seed(int64(time.Now().Nanosecond()))

	for i := 0; i < args.length; i++ {
		index := rand.Intn(len(characters))
		generatedPassword = generatedPassword + string(characters[index])
	}

	return generatedPassword
}

func getCharString(args Flags) string {
	s := CHAR_LOWER

	if args.digitsOnly {
		return DIGITS
	}

	if args.uppercase {
		s += CHAR_UPPER
	}

	if args.digits {
		s += DIGITS
	}

	if args.symbols {
		s += SYMBOLS
	}

	return s
}

package passgen

import "flag"

type Flags struct {
	length     int
	digits     bool
	symbols    bool
	uppercase  bool
	digitsOnly bool
	noCopy     bool
}

func getFlags() Flags {
	length := flag.Int("length", 12, "Length of the generated password")
	noDigits := flag.Bool("no-digits", false, "Doesn't allow digits in the generated password")
	noSymbols := flag.Bool("no-symbols", false, "Doesn't allow symbols in the generated password")
	noUppercase := flag.Bool("no-uppercase", false, "Doesn't allow uppercase letters in the generated password")
	digitsOnly := flag.Bool("digits-only", false, "Only allow digits in the generated password")
	noCopy := flag.Bool("no-copy", false, "Doesn't allow copying generated password to clipboard")

	flag.Parse()

	return Flags{
		length:     *length,
		digits:     !(*noDigits),
		symbols:    !(*noSymbols),
		uppercase:  !(*noUppercase),
		digitsOnly: *digitsOnly,
		noCopy:     *noCopy,
	}
}

package passgen

import "flag"

type flags struct {
	length     int
	digits     bool
	symbols    bool
	uppercase  bool
	lowercase  bool
	characters bool
}

func getFlags() flags {
	flag.Parse()

	return flags{
		length:     12,
		digits:     true,
		symbols:    true,
		uppercase:  true,
		lowercase:  true,
		characters: true,
	}
}

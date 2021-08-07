package passgen

import (
	"fmt"
)

func Run() {
	pass := generatePassword()

	fmt.Println(pass)
}

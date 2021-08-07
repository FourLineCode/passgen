package passgen

import (
	"fmt"
	"os/exec"
)

const (
	RESET = "\033[0m"

	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	BLUE   = "\033[34m"
	CYAN   = "\033[36m"
)

func Run() {
	args := getFlags()
	pass := generatePassword(args)

	fmt.Printf("\n%sGenrated Password: %s%s\n", string(CYAN), string(GREEN), pass)

	if args.noCopy {
		fmt.Println()
		return
	}

	copy := exec.Command("bash", "-c", fmt.Sprintf("echo %s | xclip -selection clipboard\n", pass))

	if err := copy.Run(); err != nil {
		fmt.Printf("%sError: Failed to copy password to clipboard\n\n", string(RED))
		return
	}

	fmt.Printf("%sSuccess: %sCopied generated password to clipboard\n\n", string(CYAN), string(YELLOW))
}

package main

import (
	"fmt"
	"os"

	"github.com/user/linux_coreutils/pkg/cliutil"
)

func run() int {
	name := cliutil.ProgramName()
	if name == "reset" {
		fmt.Print("\033c")
	} else {
		fmt.Print("\033[H\033[2J")
	}
	return 0
}

func main() {
	os.Exit(run())
}

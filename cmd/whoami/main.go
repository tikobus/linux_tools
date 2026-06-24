package main

import (
	"fmt"
	"os"
	"os/user"
)

func run() int {
	u, err := user.Current()
	if err != nil {
		fmt.Fprintln(os.Stderr, "whoami:", err)
		return 1
	}
	fmt.Println(u.Username)
	return 0
}

func main() {
	os.Exit(run())
}

package main

import (
	"fmt"
	"os"
)

func run() int {
	h, err := os.Hostname()
	if err != nil {
		fmt.Fprintln(os.Stderr, "hostname:", err)
		return 1
	}
	fmt.Println(h)
	return 0
}

func main() {
	os.Exit(run())
}

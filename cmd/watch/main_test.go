package main

import (
	"flag"
	"os"
	"testing"
)

func TestWatchArgs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"watch"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	code := run()
	if code == 0 {
		t.Fatal("expected non-zero exit for missing command")
	}
}

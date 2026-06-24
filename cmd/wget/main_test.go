package main

import (
	"flag"
	"os"
	"testing"
)

func TestWgetArgs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"wget"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	code := run()
	if code == 0 {
		t.Fatal("expected non-zero exit for missing URL")
	}
}

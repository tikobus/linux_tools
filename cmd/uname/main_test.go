package main

import (
	"flag"
	"os"
	"testing"
)

func TestUname(t *testing.T) {
	// Save and restore args
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"uname", "-s"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	code := run()
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
}

func TestUnameAll(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"uname", "-a"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	code := run()
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
}

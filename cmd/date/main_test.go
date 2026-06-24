package main

import (
	"flag"
	"os"
	"testing"
)

func TestDateDefault(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"date"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	code := run()
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
}

func TestDateFormat(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"date", "-format", "%Y-%m-%d"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	code := run()
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
}

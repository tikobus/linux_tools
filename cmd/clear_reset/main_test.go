package main

import (
	"os"
	"testing"
)

func TestClear(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"clear"}
	code := run()
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
}

func TestReset(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"reset"}
	code := run()
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
}

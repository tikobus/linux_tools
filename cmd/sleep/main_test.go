package main

import (
	"os"
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	start := time.Now()
	os.Args = []string{"sleep", "100ms"}
	code := run()
	elapsed := time.Since(start)
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
	if elapsed < 80*time.Millisecond {
		t.Fatalf("slept too little: %v", elapsed)
	}
}

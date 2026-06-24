package main

import (
	"testing"
)

func TestTimeout(t *testing.T) {
	code := runArgs([]string{"1s", "sleep", "5"})
	if code == 0 {
		t.Fatal("expected non-zero exit for timeout")
	}
}

func TestTimeoutSuccess(t *testing.T) {
	code := runArgs([]string{"5s", "sleep", "0.1"})
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
}

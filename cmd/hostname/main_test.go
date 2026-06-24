package main

import (
	"testing"
)

func TestHostname(t *testing.T) {
	code := run()
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
}

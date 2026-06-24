package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestBasename(t *testing.T) {
	var buf bytes.Buffer
	if err := run([]string{"/usr/local/bin/go"}, &buf, "basename"); err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(buf.String()) != "go" {
		t.Fatalf("expected 'go', got: %s", buf.String())
	}
}

func TestDirname(t *testing.T) {
	var buf bytes.Buffer
	if err := run([]string{"/usr/local/bin/go"}, &buf, "dirname"); err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(buf.String()) != "/usr/local/bin" {
		t.Fatalf("expected '/usr/local/bin', got: %s", buf.String())
	}
}

func TestBasenameMultiple(t *testing.T) {
	var buf bytes.Buffer
	if err := run([]string{"a/b", "c/d"}, &buf, "basename"); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 2 || lines[0] != "b" || lines[1] != "d" {
		t.Fatalf("unexpected output: %s", buf.String())
	}
}

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWhichPath(t *testing.T) {
	// Create a temp dir and add it to PATH
	dir := t.TempDir()
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", dir+string(filepath.ListSeparator)+oldPath)
	defer os.Setenv("PATH", oldPath)

	// Create a fake executable
	exe := filepath.Join(dir, "mytool")
	os.WriteFile(exe, []byte("#!/bin/sh\necho hello\n"), 0755)

	var buf bytes.Buffer
	if err := run([]string{"mytool"}, &buf); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "mytool") {
		t.Fatalf("expected mytool in output, got: %s", out)
	}
}

func TestWhichNotFound(t *testing.T) {
	var buf bytes.Buffer
	if err := run([]string{"nonexistenttool12345"}, &buf); err == nil {
		t.Fatal("expected error for missing command")
	}
}

func TestWhichMultiple(t *testing.T) {
	dir := t.TempDir()
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", dir+string(filepath.ListSeparator)+oldPath)
	defer os.Setenv("PATH", oldPath)

	os.WriteFile(filepath.Join(dir, "tool1"), []byte(""), 0755)
	os.WriteFile(filepath.Join(dir, "tool2"), []byte(""), 0755)

	var buf bytes.Buffer
	if err := run([]string{"tool1", "tool2"}, &buf); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "tool1") || !strings.Contains(out, "tool2") {
		t.Fatalf("expected tool1 and tool2 in output, got: %s", out)
	}
}

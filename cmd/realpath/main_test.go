package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/user/linux_coreutils/internal/testutil"
)

func TestRealpathBasic(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "file.txt", "hello")

	var buf bytes.Buffer
	if err := run([]string{f}, &buf); err != nil {
		t.Fatal(err)
	}
	out := strings.TrimSpace(buf.String())
	if !strings.Contains(out, "file.txt") {
		t.Fatalf("expected file.txt in output, got: %s", out)
	}
	if !filepath.IsAbs(out) {
		t.Fatalf("expected absolute path, got: %s", out)
	}
}

func TestRealpathRelative(t *testing.T) {
	var buf bytes.Buffer
	if err := run([]string{"."}, &buf); err != nil {
		t.Fatal(err)
	}
	out := strings.TrimSpace(buf.String())
	wd, _ := os.Getwd()
	if out != wd {
		t.Fatalf("expected %q, got %q", wd, out)
	}
}

func TestRealpathMultiple(t *testing.T) {
	dir := testutil.TempDir(t)
	a := testutil.WriteFile(t, dir, "a.txt", "")
	b := testutil.WriteFile(t, dir, "b.txt", "")

	var buf bytes.Buffer
	if err := run([]string{a, b}, &buf); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(lines))
	}
}

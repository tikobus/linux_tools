package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/user/linux_coreutils/internal/testutil"
)

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestUniqBasic(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "a\na\nb\nb\nb\nc\n")

	out := captureStdout(t, func() {
		if err := run([]string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "a\nb\nc\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestUniqCount(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "a\na\nb\nb\nb\nc\n")

	*countOpt = true
	defer func() { *countOpt = false }()

	out := captureStdout(t, func() {
		if err := run([]string{f}); err != nil {
			t.Fatal(err)
		}
	})
	lines := strings.Split(strings.TrimSuffix(out, "\n"), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
	if !strings.HasPrefix(lines[0], "      2 a") {
		t.Fatalf("unexpected first line: %q", lines[0])
	}
	if !strings.HasPrefix(lines[1], "      3 b") {
		t.Fatalf("unexpected second line: %q", lines[1])
	}
	if !strings.HasPrefix(lines[2], "      1 c") {
		t.Fatalf("unexpected third line: %q", lines[2])
	}
}

func TestUniqNoAdjacent(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "a\nb\na\n")

	out := captureStdout(t, func() {
		if err := run([]string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "a\nb\na\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

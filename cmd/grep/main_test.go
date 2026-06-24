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

func TestGrepBasic(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "hello world\nfoo bar\nhello again\n")

	out := captureStdout(t, func() {
		if err := run("hello", []string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "hello world\nhello again\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestGrepInvert(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "apple\nbanana\napple pie\n")

	*invert = true
	defer func() { *invert = false }()

	out := captureStdout(t, func() {
		if err := run("apple", []string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "banana\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestGrepCaseInsensitive(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "Hello\nHELLO\nworld\n")

	*ignoreCase = true
	defer func() { *ignoreCase = false }()

	out := captureStdout(t, func() {
		if err := run("hello", []string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "Hello\nHELLO\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestGrepLineNum(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "a\nb\nc\n")

	*lineNum = true
	defer func() { *lineNum = false }()

	out := captureStdout(t, func() {
		if err := run("b", []string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "2:b\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestGrepRecursive(t *testing.T) {
	dir := testutil.TempDir(t)
	testutil.WriteFile(t, dir, "a.txt", "match in a\n")
	testutil.WriteFile(t, dir, "subdir/b.txt", "match in b\n")

	*recursive = true
	defer func() { *recursive = false }()

	out := captureStdout(t, func() {
		if err := run("match", []string{dir}); err != nil {
			t.Fatal(err)
		}
	})
	lines := strings.Split(strings.TrimSuffix(out, "\n"), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %q", len(lines), out)
	}
	for _, line := range lines {
		if !strings.Contains(line, "match") {
			t.Fatalf("unexpected line: %q", line)
		}
	}
}

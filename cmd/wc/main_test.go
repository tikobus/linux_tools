package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
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

func TestWcDefault(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "hello world\nfoo bar baz\n")

	// Reset flags to default
	*linesOpt = false
	*wordsOpt = false
	*bytesOpt = false

	out := captureStdout(t, func() {
		if err := run([]string{f}); err != nil {
			t.Fatal(err)
		}
	})
	// Should print lines words bytes filename
	if !strings.Contains(out, "2") || !strings.Contains(out, "5") {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestWcLines(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "a\nb\nc\n")

	*linesOpt = true
	*wordsOpt = false
	*bytesOpt = false
	defer func() { *linesOpt = false }()

	out := captureStdout(t, func() {
		if err := run([]string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "       3 " + filepath.Base(f) + "\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestWcBytes(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "abc\n")

	*bytesOpt = true
	*linesOpt = false
	*wordsOpt = false
	defer func() { *bytesOpt = false }()

	out := captureStdout(t, func() {
		if err := run([]string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "       4 " + filepath.Base(f) + "\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

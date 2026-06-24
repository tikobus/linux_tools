package main

import (
	"bytes"
	"io"
	"os"
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

func TestTailLines(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "a\nb\nc\nd\ne\n")

	*linesOpt = 3
	*bytesOpt = -1
	defer func() { *linesOpt = 10; *bytesOpt = -1 }()

	out := captureStdout(t, func() {
		if err := run([]string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "c\nd\ne\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestTailBytes(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "abcdef\n")

	*bytesOpt = 4
	*linesOpt = 10
	defer func() { *bytesOpt = -1; *linesOpt = 10 }()

	out := captureStdout(t, func() {
		if err := run([]string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "def\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestTailNoTrailingNewline(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "a\nb\nc")

	*linesOpt = 2
	*bytesOpt = -1
	defer func() { *linesOpt = 10; *bytesOpt = -1 }()

	out := captureStdout(t, func() {
		if err := run([]string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "b\nc\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

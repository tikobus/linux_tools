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

func TestHeadLines(t *testing.T) {
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
	want := "a\nb\nc\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestHeadBytes(t *testing.T) {
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
	want := "abcd"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestHeadStdin(t *testing.T) {
	*linesOpt = 2
	*bytesOpt = -1
	defer func() { *linesOpt = 10; *bytesOpt = -1 }()

	// Save stdin
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.WriteString("line1\nline2\nline3\n")
	w.Close()
	defer func() { os.Stdin = oldStdin }()

	out := captureStdout(t, func() {
		if err := run([]string{}); err != nil {
			t.Fatal(err)
		}
	})
	want := "line1\nline2\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

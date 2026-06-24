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

func TestCutBytes(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "abcdef\n123456\n")

	*bytesOpt = "1-3,5"
	defer func() { *bytesOpt = "" }()

	out := captureStdout(t, func() {
		if err := run([]string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "abce\n1235\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestCutFields(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "a\tb\tc\n1\t2\t3\n")

	*fieldsOpt = "1,3"
	defer func() { *fieldsOpt = "" }()

	out := captureStdout(t, func() {
		if err := run([]string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "a\tc\n1\t3\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestCutChars(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "hello world\nfoo bar\n")

	*charsOpt = "1-5"
	defer func() { *charsOpt = "" }()

	out := captureStdout(t, func() {
		if err := run([]string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "hello\nfoo b\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

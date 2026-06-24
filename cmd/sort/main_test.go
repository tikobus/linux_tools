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

func TestSortBasic(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "cherry\napple\nbanana\n")

	out := captureStdout(t, func() {
		if err := run([]string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "apple\nbanana\ncherry\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestSortNumeric(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "10\n2\n1\n20\n")

	*numeric = true
	defer func() { *numeric = false }()

	out := captureStdout(t, func() {
		if err := run([]string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "1\n2\n10\n20\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestSortUnique(t *testing.T) {
	dir := testutil.TempDir(t)
	f := testutil.WriteFile(t, dir, "test.txt", "apple\nbanana\napple\ncherry\nbanana\n")

	*uniqueOpt = true
	defer func() { *uniqueOpt = false }()

	out := captureStdout(t, func() {
		if err := run([]string{f}); err != nil {
			t.Fatal(err)
		}
	})
	want := "apple\nbanana\ncherry\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

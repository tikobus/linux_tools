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

func TestPasteBasic(t *testing.T) {
	dir := testutil.TempDir(t)
	f1 := testutil.WriteFile(t, dir, "a.txt", "a\nb\nc\n")
	f2 := testutil.WriteFile(t, dir, "b.txt", "1\n2\n3\n")

	out := captureStdout(t, func() {
		if err := run([]string{f1, f2}); err != nil {
			t.Fatal(err)
		}
	})
	want := "a\t1\nb\t2\nc\t3\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestPasteCustomDelim(t *testing.T) {
	dir := testutil.TempDir(t)
	f1 := testutil.WriteFile(t, dir, "a.txt", "a\nb\n")
	f2 := testutil.WriteFile(t, dir, "b.txt", "1\n2\n")

	*delimOpt = ","
	defer func() { *delimOpt = "\t" }()

	out := captureStdout(t, func() {
		if err := run([]string{f1, f2}); err != nil {
			t.Fatal(err)
		}
	})
	want := "a,1\nb,2\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestPasteUneven(t *testing.T) {
	dir := testutil.TempDir(t)
	f1 := testutil.WriteFile(t, dir, "a.txt", "a\nb\nc\n")
	f2 := testutil.WriteFile(t, dir, "b.txt", "1\n")

	out := captureStdout(t, func() {
		if err := run([]string{f1, f2}); err != nil {
			t.Fatal(err)
		}
	})
	want := "a\t1\nb\t\nc\t\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

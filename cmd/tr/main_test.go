package main

import (
	"bytes"
	"io"
	"os"
	"testing"
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

func TestTrTranslate(t *testing.T) {
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.WriteString("hello")
	w.Close()
	defer func() { os.Stdin = oldStdin }()

	out := captureStdout(t, func() {
		if err := run([]string{"a-z", "A-Z"}); err != nil {
			t.Fatal(err)
		}
	})
	want := "HELLO"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestTrDelete(t *testing.T) {
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.WriteString("hello world")
	w.Close()
	defer func() { os.Stdin = oldStdin }()

	*deleteOpt = "a-z"
	defer func() { *deleteOpt = "" }()

	out := captureStdout(t, func() {
		if err := run([]string{}); err != nil {
			t.Fatal(err)
		}
	})
	want := " "
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestTrExpand(t *testing.T) {
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.WriteString("abc")
	w.Close()
	defer func() { os.Stdin = oldStdin }()

	out := captureStdout(t, func() {
		if err := run([]string{"a-c", "x-z"}); err != nil {
			t.Fatal(err)
		}
	})
	want := "xyz"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

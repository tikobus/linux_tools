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

func TestPrintfBasic(t *testing.T) {
	out := captureStdout(t, func() {
		if err := run([]string{"hello %s", "world"}); err != nil {
			t.Fatal(err)
		}
	})
	want := "hello world"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestPrintfNewline(t *testing.T) {
	out := captureStdout(t, func() {
		if err := run([]string{"hello\n"}); err != nil {
			t.Fatal(err)
		}
	})
	want := "hello\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestPrintfNoFormat(t *testing.T) {
	out := captureStdout(t, func() {
		if err := run([]string{"plain text"}); err != nil {
			t.Fatal(err)
		}
	})
	want := "plain text"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

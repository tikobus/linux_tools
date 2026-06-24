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

func TestEchoBasic(t *testing.T) {
	out := captureStdout(t, func() {
		if err := run([]string{"hello", "world"}); err != nil {
			t.Fatal(err)
		}
	})
	want := "hello world\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestEchoNoNewline(t *testing.T) {
	*noNewline = true
	defer func() { *noNewline = false }()

	out := captureStdout(t, func() {
		if err := run([]string{"hello"}); err != nil {
			t.Fatal(err)
		}
	})
	want := "hello"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

func TestEchoEmpty(t *testing.T) {
	out := captureStdout(t, func() {
		if err := run([]string{}); err != nil {
			t.Fatal(err)
		}
	})
	want := "\n"
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
}

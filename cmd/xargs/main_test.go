package main

import (
	"bytes"
	"io"
	"os"
	"strings"
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

func TestXargsBasic(t *testing.T) {
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.WriteString("hello\n")
	w.Close()
	defer func() { os.Stdin = oldStdin }()

	out := captureStdout(t, func() {
		if err := run([]string{"echo"}); err != nil {
			t.Fatal(err)
		}
	})
	if strings.TrimSpace(out) != "hello" {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestXargsWithArgs(t *testing.T) {
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.WriteString("world\n")
	w.Close()
	defer func() { os.Stdin = oldStdin }()

	out := captureStdout(t, func() {
		if err := run([]string{"echo", "hello"}); err != nil {
			t.Fatal(err)
		}
	})
	if strings.TrimSpace(out) != "hello world" {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestXargsMultipleLines(t *testing.T) {
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.WriteString("a\nb\n")
	w.Close()
	defer func() { os.Stdin = oldStdin }()

	out := captureStdout(t, func() {
		if err := run([]string{"echo"}); err != nil {
			t.Fatal(err)
		}
	})
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %q", len(lines), out)
	}
}

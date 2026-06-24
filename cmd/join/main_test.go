package main

import (
	"bytes"
	"io"
	"os"
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

func TestJoinBasic(t *testing.T) {
	dir := testutil.TempDir(t)
	f1 := testutil.WriteFile(t, dir, "a.txt", "1 apple red\n2 banana yellow\n3 cherry red\n")
	f2 := testutil.WriteFile(t, dir, "b.txt", "1 sweet\n2 long\n3 tart\n")

	out := captureStdout(t, func() {
		if err := run([]string{f1, f2}); err != nil {
			t.Fatal(err)
		}
	})
	lines := strings.Split(strings.TrimSuffix(out, "\n"), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d: %q", len(lines), out)
	}
	if !strings.Contains(lines[0], "1") && !strings.Contains(lines[0], "2") && !strings.Contains(lines[0], "3") {
		t.Fatalf("unexpected first line: %q", lines[0])
	}
}

func TestJoinField2(t *testing.T) {
	dir := testutil.TempDir(t)
	f1 := testutil.WriteFile(t, dir, "a.txt", "1 apple red\n2 banana yellow\n")
	f2 := testutil.WriteFile(t, dir, "b.txt", "sweet 1\nlong 2\n")

	*field2Opt = 2
	defer func() { *field2Opt = 1 }()

	out := captureStdout(t, func() {
		if err := run([]string{f1, f2}); err != nil {
			t.Fatal(err)
		}
	})
	lines := strings.Split(strings.TrimSuffix(out, "\n"), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %q", len(lines), out)
	}
}

func TestJoinNoMatch(t *testing.T) {
	dir := testutil.TempDir(t)
	f1 := testutil.WriteFile(t, dir, "a.txt", "1 apple\n")
	f2 := testutil.WriteFile(t, dir, "b.txt", "2 banana\n")

	out := captureStdout(t, func() {
		if err := run([]string{f1, f2}); err != nil {
			t.Fatal(err)
		}
	})
	if strings.TrimSpace(out) != "" {
		t.Fatalf("expected no output, got %q", out)
	}
}
